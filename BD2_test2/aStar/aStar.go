package aStar

import (
	"app/MyOpenCV"
	"app/myMotion"
	"container/heap"
	"context"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/Dasongzi1366/AutoGo/motion"
	"github.com/Dasongzi1366/AutoGo/opencv"
	"github.com/Dasongzi1366/AutoGo/yolo"
)

// 二值化地图
// 参数1：地图路径    参数2：给二值化后地图的名称
func BuildObstacleMap(mapPath string, fileName string) {
	mapImg := opencv.IMRead(mapPath, opencv.IMReadColor)
	if mapImg.Empty() {
		fmt.Println("错误：无法读取原始地图图片，请检查路径是否正确！")
		return
	}
	gray := opencv.NewMat()
	binary := opencv.NewMat()

	opencv.CvtColor(mapImg, &gray, opencv.ColorBGRToGray)
	// 阈值根据你的地图颜色调整
	opencv.Threshold(gray, &binary, 100, 255, opencv.ThresholdBinary)

	dir := filepath.Dir(mapPath)
	//file := filepath.Base(mapPath)
	saveFileName := filepath.Join(dir, fileName)

	success := opencv.IMWrite(saveFileName, binary)
	if !success {
		fmt.Println("错误：保存图片失败！可能是路径权限问题，或者图片数据为空。")
		return
	}
	fmt.Println("二值化图片保存到：", saveFileName)
	return // 255=可通行 0=障碍
}

//// 通过二值化地图获得障碍物地图
//func LoadObstacleMap(imgPath string) ([][]bool, error) {
//	// 直接读取已经二值化的地图
//	img := opencv.IMRead(imgPath, opencv.IMReadGrayScale)
//	if img.Empty() {
//		return nil, fmt.Errorf("地图加载失败: %s", imgPath)
//	}
//	defer img.Close()
//
//	rows := img.Rows()
//	cols := img.Cols()
//	obstacle := make([][]bool, rows)
//
//	for y := 0; y < rows; y++ {
//		obstacle[y] = make([]bool, cols)
//		for x := 0; x < cols; x++ {
//			obstacle[y][x] = img.GetUCharAt(y, x) == 0 // 0=黑色=障碍
//		}
//	}
//
//	return obstacle, nil
//}

// ===================A*寻路==========================
const (
	STEP_VAL    = 10
	OBLIQUE_VAL = 14
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Node struct {
	Point
	G, H     int
	Parent   *Node
	Index    int
	InOpen   bool
	InClosed bool
}

func (n *Node) F() int { return n.G + n.H }

// 优先队列（小顶堆，F值小的优先）
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) } //返回集合（如切片、数组、字符串、Map或通道）中元素的个数（长度）

// 索引为 i 的元素，是否应该排在索引为 j 的元素前面,pq[i].F() < pq[j].F()成立则返回true,否则返回false
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].F() < pq[j].F() }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i] //交换切片中 i 和 j 位置的元素
	pq[i].Index = i             //因为位置换了，要更新 i 位置上元素的 Index 字段为 i
	pq[j].Index = j             //更新 j 位置上元素的 Index 字段为 j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := x.(*Node)
	n.Index = len(*pq)
	*pq = append(*pq, n)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := old[len(old)-1]
	*pq = old[:len(old)-1]
	return n
}

// AStarMap 封装地图数据和距墙距离缓存
type AStarMap struct {
	obstacle   [][]bool
	distToWall [][]int // 每个点到最近障碍的距离
	rows, cols int
}

// LoadObstacleMap 加载二值化地图
func LoadObstacleMap(imgPath string) (*AStarMap, error) {
	img := opencv.IMRead(imgPath, opencv.IMReadGrayScale)
	if img.Empty() {
		return nil, fmt.Errorf("地图加载失败: %s", imgPath)
	}
	defer img.Close()

	rows := img.Rows()
	cols := img.Cols()

	obstacle := make([][]bool, rows)
	for y := 0; y < rows; y++ {
		obstacle[y] = make([]bool, cols)
		for x := 0; x < cols; x++ {
			obstacle[y][x] = img.GetUCharAt(y, x) == 0
		}
	}

	m := &AStarMap{
		obstacle: obstacle,
		rows:     rows,
		cols:     cols,
	}
	m.buildDistToWall() // 预计算距墙距离
	return m, nil
}

// buildDistToWall BFS多源最短距离，预计算每个点到障碍的距离
func (m *AStarMap) buildDistToWall() {
	m.distToWall = make([][]int, m.rows)
	for y := range m.distToWall {
		m.distToWall[y] = make([]int, m.cols)
		for x := range m.distToWall[y] {
			m.distToWall[y][x] = math.MaxInt32
		}
	}

	type pos struct{ x, y int }
	queue := []pos{}

	// 所有障碍点距离为0，加入队列
	for y := 0; y < m.rows; y++ {
		for x := 0; x < m.cols; x++ {
			if m.obstacle[y][x] {
				m.distToWall[y][x] = 0
				queue = append(queue, pos{x, y})
			}
		}
	}

	dx4 := []int{1, -1, 0, 0}
	dy4 := []int{0, 0, 1, -1}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curDist := m.distToWall[cur.y][cur.x]

		for d := 0; d < 4; d++ {
			nx, ny := cur.x+dx4[d], cur.y+dy4[d]
			if nx < 0 || ny < 0 || nx >= m.cols || ny >= m.rows {
				continue
			}
			if m.distToWall[ny][nx] > curDist+1 {
				m.distToWall[ny][nx] = curDist + 1
				queue = append(queue, pos{nx, ny})
			}
		}
	}
}

func (m *AStarMap) getDistToWall(x, y int) int {
	if x < 0 || y < 0 || x >= m.cols || y >= m.rows {
		return 0
	}
	d := m.distToWall[y][x]
	if d == math.MaxInt32 {
		return 999
	}
	return d
}

func (m *AStarMap) calcG(parent *Node, nx, ny int) int {
	dx := abs(nx - parent.X)
	dy := abs(ny - parent.Y)
	baseCost := STEP_VAL
	if dx+dy == 2 {
		baseCost = OBLIQUE_VAL
	}

	// 边缘惩罚：越靠近障碍代价越大
	dist := m.getDistToWall(nx, ny)
	penalty := 0
	if dist <= 0 {
		penalty = 300
	} else {
		penalty = 120 / dist
	}

	return parent.G + baseCost + penalty
}

func calcH(x, y, ex, ey int) int {
	return (abs(ex-x) + abs(ey-y)) * STEP_VAL
}

// findNearestPassable 起点在障碍物时BFS找最近可通行点
func (m *AStarMap) findNearestPassable(x, y int) (int, int, bool) {
	type pos struct{ x, y int }
	visited := make([][]bool, m.rows)
	for i := range visited {
		visited[i] = make([]bool, m.cols)
	}

	queue := []pos{{x, y}}
	visited[y][x] = true

	dx4 := []int{1, -1, 0, 0}
	dy4 := []int{0, 0, 1, -1}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if !m.obstacle[cur.y][cur.x] {
			return cur.x, cur.y, true
		}

		for d := 0; d < 4; d++ {
			nx, ny := cur.x+dx4[d], cur.y+dy4[d]
			if nx < 0 || ny < 0 || nx >= m.cols || ny >= m.rows {
				continue
			}
			if !visited[ny][nx] {
				visited[ny][nx] = true
				queue = append(queue, pos{nx, ny})
			}
		}
	}
	return x, y, false
}

// 八方向移动
var dirs = []Point{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1}, // 上下左右
	{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, // 斜角
}

//func heuristic(a, b Point) float64 { //计算当前坐标到终点的代价,用的是理论值,也就是直线
//	// 对角线距离，适合八方向移动
//	// 先走短边再走长边
//	dx := abs(a.X - b.X)
//	dy := abs(a.Y - b.Y)
//	if dx > dy {
//		return float64(dx) + 0.414*float64(dy)
//	}
//	return float64(dy) + 0.414*float64(dx)
//}

func AStar(m *AStarMap, start, end Point) []Point {
	// 起点在障碍物时找最近可通行点
	if m.obstacle[start.Y][start.X] {
		nx, ny, ok := m.findNearestPassable(start.X, start.Y)
		if !ok {
			return nil
		}
		start = Point{nx, ny}
	}

	if m.obstacle[end.Y][end.X] {
		return nil
	}

	nodes := make([][]*Node, m.rows)
	for i := range nodes {
		nodes[i] = make([]*Node, m.cols)
	}

	getNode := func(x, y int) *Node {
		if nodes[y][x] == nil {
			nodes[y][x] = &Node{Point: Point{x, y}}
		}
		return nodes[y][x]
	}

	startNode := getNode(start.X, start.Y)
	startNode.G = 0
	startNode.H = calcH(start.X, start.Y, end.X, end.Y)
	startNode.InOpen = true

	pq := &PriorityQueue{startNode}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Node)
		cur.InOpen = false
		cur.InClosed = true

		if cur.X == end.X && cur.Y == end.Y {
			return buildPath(cur)
		}

		for _, d := range dirs {
			nx, ny := cur.X+d.X, cur.Y+d.Y
			if nx < 0 || ny < 0 || nx >= m.cols || ny >= m.rows {
				continue
			}
			if m.obstacle[ny][nx] {
				continue
			}
			// 斜角移动检查两侧
			if d.X != 0 && d.Y != 0 {
				if m.obstacle[cur.Y][nx] || m.obstacle[ny][cur.X] {
					continue
				}
			}

			neighbor := getNode(nx, ny)
			if neighbor.InClosed {
				continue
			}

			ng := m.calcG(cur, nx, ny)

			if neighbor.InOpen {
				if ng < neighbor.G {
					neighbor.G = ng
					neighbor.Parent = cur
					heap.Fix(pq, neighbor.Index)
				}
			} else {
				neighbor.G = ng
				neighbor.H = calcH(nx, ny, end.X, end.Y)
				neighbor.Parent = cur
				neighbor.InOpen = true
				heap.Push(pq, neighbor)
			}
		}
	}
	return nil
}

//heap.Pop 和 heap.Pus会调用Less和swap方法自动保证pq中第一个的F是最小的

func buildPath(end *Node) []Point {
	var path []Point
	for n := end; n != nil; n = n.Parent {
		path = append(path, n.Point)
	}
	// 反转：从起点到终点
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		//i是数组开头,j是数组结尾,互换数组开头和结尾的的元素,然后两者向内部移动,直到i=j
		path[i], path[j] = path[j], path[i]
	}
	return path
}

// DouglasPeucker 路径简化
func DouglasPeucker(path []Point, epsilon float64) []Point {
	if len(path) <= 2 {
		return path
	}

	start := path[0]
	end := path[len(path)-1]

	dx := float64(end.X - start.X)
	dy := float64(end.Y - start.Y)
	lineLenSq := dx*dx + dy*dy

	maxDist := 0.0
	maxIdx := 0

	for i := 1; i < len(path)-1; i++ {
		var dist float64
		if lineLenSq == 0 {
			dist = math.Sqrt(float64(path[i].X-start.X)*float64(path[i].X-start.X) +
				float64(path[i].Y-start.Y)*float64(path[i].Y-start.Y))
		} else {
			t := math.Max(0, math.Min(1,
				(float64(path[i].X-start.X)*dx+float64(path[i].Y-start.Y)*dy)/lineLenSq))
			projX := float64(start.X) + t*dx
			projY := float64(start.Y) + t*dy
			dist = math.Sqrt((float64(path[i].X)-projX)*(float64(path[i].X)-projX) +
				(float64(path[i].Y)-projY)*(float64(path[i].Y)-projY))
		}
		if dist > maxDist {
			maxDist = dist
			maxIdx = i
		}
	}

	if maxDist > epsilon {
		left := DouglasPeucker(path[:maxIdx+1], epsilon)
		right := DouglasPeucker(path[maxIdx:], epsilon)
		return append(left, right[1:]...)
	}
	return []Point{start, end}
}

// PushPathToCenter 将路径点推向通道中心
func PushPathToCenter(m *AStarMap, path []Point, iterations, searchR int) []Point {
	result := make([]Point, len(path))
	copy(result, path)

	for iter := 0; iter < iterations; iter++ {
		for i := 1; i < len(result)-1; i++ {
			pt := result[i]
			bestDist := m.getDistToWall(pt.X, pt.Y)
			bestX, bestY := pt.X, pt.Y

			// 计算切线方向
			tdx := float64(result[i+1].X - result[i-1].X)
			tdy := float64(result[i+1].Y - result[i-1].Y)
			tlen := math.Sqrt(tdx*tdx + tdy*tdy)
			if tlen < 0.001 {
				continue
			}

			ux := tdx / tlen
			uy := tdy / tlen
			// 垂直方向
			px := -uy
			py := ux

			// 沿垂直方向搜索
			for _, side := range []float64{-1, 1} {
				for r := 1; r <= searchR; r++ {
					nx := int(float64(pt.X) + px*side*float64(r) + 0.5)
					ny := int(float64(pt.Y) + py*side*float64(r) + 0.5)
					if nx < 0 || ny < 0 || nx >= m.cols || ny >= m.rows {
						break
					}
					if m.obstacle[ny][nx] {
						break
					}
					d := m.getDistToWall(nx, ny)
					if d > bestDist {
						bestDist = d
						bestX, bestY = nx, ny
					}
				}
			}
			result[i] = Point{bestX, bestY}
		}
	}
	return result
}

// 插值：确保相邻路径点间距不超过maxDist
func interpolatePath(path []Point, maxDist float64) []Point {
	if len(path) <= 1 {
		return path
	}

	result := []Point{path[0]}

	for i := 1; i < len(path); i++ {
		prev := result[len(result)-1]
		curr := path[i]

		dx := float64(curr.X - prev.X)
		dy := float64(curr.Y - prev.Y)
		dist := math.Sqrt(dx*dx + dy*dy)

		// 超过maxDist就插入中间点
		if dist > maxDist {
			steps := int(dist/maxDist) + 1
			for s := 1; s < steps; s++ {
				t := float64(s) / float64(steps)
				mx := prev.X + int(t*dx)
				my := prev.Y + int(t*dy)
				result = append(result, Point{mx, my})
			}
		}

		result = append(result, curr)
	}

	return result
}

// 完整寻路接口
func FindPath(m *AStarMap, start, end Point) []Point {
	raw := AStar(m, start, end)
	if raw == nil {
		return nil
	}

	// Douglas-Peucker简化
	// 0.5太激进，改成2.0保留更多转折点
	//simplified := DouglasPeucker(raw, 0.5)
	simplified := DouglasPeucker(raw, 0.5)

	// 推向通道中心
	centered := PushPathToCenter(m, simplified, 5, 15)

	//for i := 1; i < len(centered); i++ { //打印路径点距
	//	dx := centered[i].X - centered[i-1].X
	//	dy := centered[i].Y - centered[i-1].Y
	//	dist := math.Sqrt(float64(dx*dx + dy*dy))
	//	fmt.Printf("点%d→点%d 距离=%.1f\n", i-1, i, dist)
	//}

	// 确保相邻点间距不超过20像素
	interpolated := interpolatePath(centered, 20.0)
	// 验证插值结果
	for i := 1; i < len(interpolated); i++ {
		dx := interpolated[i].X - interpolated[i-1].X
		dy := interpolated[i].Y - interpolated[i-1].Y
		dist := math.Sqrt(float64(dx*dx + dy*dy))
		if dist > 20 {
			fmt.Printf("插值后仍有超距: 点%d→点%d 距离=%.1f\n", i-1, i, dist)
		}
	}

	return interpolated
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//===================A*寻路==========================

// ===============路径点转按键控制=================

// 跑完整条路径
//func FollowPath(path []Point, getCurrentPos func() Point) {
//	const arriveThreshold = 15
//	defer myMotion.StopMove()
//
//	for _, waypoint := range path {
//		for {
//			cur := getCurrentPos()
//
//			dx := waypoint.X - cur.X
//			dy := waypoint.Y - cur.Y
//			dist := math.Sqrt(float64(dx*dx + dy*dy))
//
//			if dist < arriveThreshold {
//				break
//			}
//
//			holdX, holdY := myMotion.CalcHoldPoint(cur.X, cur.Y, waypoint.X, waypoint.Y) //归一化向量
//			myMotion.StartMoveXY(holdX, holdY)
//
//			time.Sleep(16 * time.Millisecond)
//		}
//	}
//}

type WayFindState int

const (
	STATE_Done             WayFindState = iota // 正常到达终点
	STATE_Movement_timeout                     // 寻路超时
	STATE_CDT_Useless                          // 坐标异常
	STATE_Loss_Map                             //小地图丢失
)

// 全局context，地图丢失时取消所有操作
var (
	globalCtx    context.Context
	globalCancel context.CancelFunc
)

// 启动地图检测线程
func StartMapMonitor(bigMapPath string, cancelFunc context.CancelFunc) {
	go func() {
		for {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			if x == -1 && y == -1 {
				fmt.Println("地图丢失，终止所有操作")
				cancelFunc() // 取消所有操作
				return
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
}

func FollowPath(ctx context.Context, path []Point, getCurrentPos func() Point) WayFindState {
	if len(path) == 0 {
		return STATE_CDT_Useless // 假如路径为空当坐标异常处理，触发重新寻路
	}
	defer myMotion.StopMove()

	detachmentTime := 0

	for _, waypoint := range path {
		tryCount := 0

		for tryCount < 15 {
			// 检查是否被取消
			select {
			case <-ctx.Done():
				return STATE_Loss_Map
			default:
			}

			cur := getCurrentPos()

			dx := waypoint.X - cur.X
			dy := waypoint.Y - cur.Y
			dist := math.Sqrt(float64(dx*dx + dy*dy))

			fmt.Printf("tryCount=%d cur=(%d,%d) waypoint=(%d,%d) dist=%.1f\n",
				tryCount, cur.X, cur.Y, waypoint.X, waypoint.Y, dist)

			if dist < 15 {
				break
			}

			if dist > 40 {
				return STATE_CDT_Useless
			}

			holdX, holdY := myMotion.CalcHoldPoint(cur.X, cur.Y, waypoint.X, waypoint.Y)
			//fmt.Printf("发送触摸: holdX=%d holdY=%d\n", holdX, holdY)
			myMotion.StartMoveXY(holdX, holdY)

			tryCount++
			time.Sleep(16 * time.Millisecond)
		}

		if tryCount >= 15 {
			detachmentTime++
		}

		if detachmentTime >= 2 {
			return STATE_Movement_timeout
		}
	}

	return STATE_Done
}

// 封装寻路+异常处理，支持重试
func NavigateTo(bigMapPath, bin_mapPath string, end Point,
	getCurrentPos func() Point) WayFindState {

	for {
		// 每次寻路前创建新的context
		ctx, cancel := context.WithCancel(context.Background())

		// 启动地图监测
		StartMapMonitor(bigMapPath, cancel)

		// 获取当前坐标
		x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
		if x == -1 && y == -1 {
			cancel()
			handleLossMap(bigMapPath) // 异常处理
			continue                  // 处理完后重试
		}

		start := Point{X: x, Y: y}
		astarMap, err := LoadObstacleMap(bin_mapPath)
		if err != nil {
			panic(err)
		}
		//path := AStar(astarMap, start, end)
		path := FindPath(astarMap, start, end)
		if path == nil || len(path) == 0 {
			fmt.Println("路径为空，等待后重试")
			time.Sleep(500 * time.Millisecond)
			continue // 重新循环重新寻路
		}

		state := FollowPath(ctx, path, getCurrentPos)
		cancel() // 停止地图监测

		switch state {
		case STATE_Done:
			return STATE_Done

		case STATE_Loss_Map:
			fmt.Println("地图丢失，进入异常处理")
			handleLossMap(bigMapPath)
			// 异常处理完后继续循环重试

		case STATE_Movement_timeout:
			fmt.Println("寻路超时，重新规划路径")
			// 超时,一般时按键冲突，四周点两下然后重新寻路
			motion.Click(640, 360, 0, 0)
			time.Sleep(700 * time.Millisecond)
			motion.Click(650, 340, 0, 0)
		case STATE_CDT_Useless:
			//坐标异常,直接重新寻路
			fmt.Println("坐标异常，等待后重试")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 地图丢失的异常处理
func handleLossMap(bigMapPath string) {
	fmt.Println("执行异常处理")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(3 * time.Second)
	for i := 0; i < 4; i++ {
		if MyOpenCV.If_TpInterface(0.85) { //看看是不是进到传送界面
			for i := 0; i < 3; i++ {
				motion.Click(105, 40, 0, 0)
				if !MyOpenCV.If_TpInterface(0.85) {
					return
				}
				if i == 2 {
					println("进入传送界面未退出")
					os.Exit(0)
				}
				time.Sleep(1000 * time.Millisecond)
			}
		}

		inBattle := MyOpenCV.If_BattleInterface(0.85)
		if !inBattle { // 没检测出来战斗界面, 有可能是图片识别错误也可能是网络延迟,在地图上随便点两下
			// 生成一个随机角度
			theta := r.Float64() * 2 * math.Pi

			// 计算第一次点击的坐标
			x1 := 640 + int(float64(150)*math.Cos(theta))
			y1 := 360 + int(float64(150)*math.Sin(theta))

			// 计算相反方向的坐标 (角度 + Pi)
			x2 := 640 + int(float64(150)*math.Cos(theta+math.Pi))
			y2 := 360 + int(float64(150)*math.Sin(theta+math.Pi))

			for i := 0; i < 2; i++ {
				if i == 0 {
					motion.Click(x1, y1, 0, 0)
				} else {
					motion.Click(x2, y2, 0, 0)
				}
				time.Sleep(1000 * time.Millisecond)

				x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
				if !(x == -1 && y == -1) {
					return
				}
			}
		}

		//看看是不是进入战斗界面了
		for i := 0; i < 4; i++ {
			inBattle = MyOpenCV.If_BattleInterface(0.85)
			if inBattle { //退出战斗
				println("战斗界面,退出战斗")
				motion.Click(1178, 42, 0, 0)
				time.Sleep(1 * time.Second)
				motion.Click(510, 660, 0, 0)
				time.Sleep(1 * time.Second)
				motion.Click(721, 433, 0, 0)
			} else {
				return
			}
			time.Sleep(3 * time.Second)
		}

		time.Sleep(2 * time.Second)

		if i == 3 {
			println("地图消失问题未处理")
			os.Exit(0)
		}
		time.Sleep(1500 * time.Millisecond)
	}

}

//===============路径点转按键控制=================

//// ===============路径平滑（减少路径点）=================
//// 射线检测：两点之间是否全部可通行
//func lineOfSight(obstacle [][]bool, a, b Point) bool {
//	dx := abs(b.X - a.X)
//	dy := abs(b.Y - a.Y)
//	sx, sy := 1, 1
//	if a.X > b.X {
//		sx = -1
//	}
//	if a.Y > b.Y {
//		sy = -1
//	}
//	err := dx - dy
//
//	x, y := a.X, a.Y
//	for {
//		if obstacle[y][x] {
//			return false
//		}
//		if x == b.X && y == b.Y {
//			return true
//		}
//		e2 := 2 * err
//		if e2 > -dy {
//			err -= dy
//			x += sx
//		}
//		if e2 < dx {
//			err += dx
//			y += sy
//		}
//	}
//}
//
//// 路径平滑：跳过中间可以直达的点
//func smoothPath(obstacle [][]bool, path []Point) []Point {
//	if len(path) <= 2 {
//		return path
//	}
//
//	smooth := []Point{path[0]}
//	cur := 0
//
//	for cur < len(path)-1 {
//		// 尽量找最远的可直达点
//		next := cur + 1
//		for next+1 < len(path) && lineOfSight(obstacle, path[cur], path[next+1]) {
//			next++
//		}
//		smooth = append(smooth, path[next])
//		cur = next
//	}
//	return smooth
//}
//
////===============路径平滑（减少路径点）=================

func YoloFind(yoloPtr *yolo.Yolo, bigMapPath string) {
	println("yolo找怪")
	for {
		for i := 0; i < 3; i++ {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			if x == -1 && y == -1 { //检查是否进入战斗
				handleLossMap(bigMapPath) // 异常处理
			} else {
				break
			}
			if i == 2 {
				println("yolo找怪时进入战斗,并且未退出战斗,退出进程")
				os.Exit(0)
			}
			time.Sleep(1 * time.Second)
		}
		results := yoloPtr.Detect(0, 0, 0, 0, 0)
		if len(results) == 0 {
			break
		}
		motion.Click(results[0].CenterX, results[0].CenterY, 0, 0)
		time.Sleep(1 * time.Second)
	}
}
