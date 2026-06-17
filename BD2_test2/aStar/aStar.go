package aStar

import (
	"app/MyOpenCV"
	"app/info"
	"app/info/info_chapter"
	"app/myMotion"
	"container/heap"
	"context"
	"fmt"
	"image"
	"math"
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

	// 开运算去除孤立白点噪声
	cleaned := opencv.NewMat()
	defer cleaned.Close()
	kernel := opencv.GetStructuringElement(opencv.MorphRect, image.Point{X: 3, Y: 3})
	defer kernel.Close()
	opencv.MorphologyEx(img, &cleaned, opencv.MorphOpen, kernel)

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

// 先用BFS标记终点所在的连通区域
func (m *AStarMap) getConnectedRegion(x, y int) map[[2]int]bool {
	region := make(map[[2]int]bool)
	if m.obstacle[y][x] {
		return region
	}

	type pos struct{ x, y int }
	queue := []pos{{x, y}}
	region[[2]int{x, y}] = true

	dx4 := []int{1, -1, 0, 0}
	dy4 := []int{0, 0, 1, -1}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for d := 0; d < 4; d++ {
			nx, ny := cur.x+dx4[d], cur.y+dy4[d]
			if nx < 0 || ny < 0 || nx >= m.cols || ny >= m.rows {
				continue
			}
			if !m.obstacle[ny][nx] && !region[[2]int{nx, ny}] {
				region[[2]int{nx, ny}] = true
				queue = append(queue, pos{nx, ny})
			}
		}
	}
	return region
}

// 找到起点到终点连通区域最近点并开凿通路
func (m *AStarMap) carvePathToRegion(sx, sy int, region map[[2]int]bool) bool {
	type pos struct{ x, y int }
	visited := make([][]bool, m.rows)
	for i := range visited {
		visited[i] = make([]bool, m.cols)
	}

	queue := []pos{{sx, sy}}
	visited[sy][sx] = true

	dx4 := []int{1, -1, 0, 0}
	dy4 := []int{0, 0, 1, -1}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if region[[2]int{cur.x, cur.y}] {
			// 用L形路径开凿：先水平再垂直
			// 第一段：水平从sx到cur.x
			x0 := sx
			stepX := 1
			if cur.x < sx {
				stepX = -1
			}
			for x0 != cur.x {
				m.obstacle[sy][x0] = false
				x0 += stepX
			}
			m.obstacle[sy][cur.x] = false

			// 第二段：垂直从sy到cur.y
			y0 := sy
			stepY := 1
			if cur.y < sy {
				stepY = -1
			}
			for y0 != cur.y {
				m.obstacle[y0][cur.x] = false
				y0 += stepY
			}
			m.obstacle[cur.y][cur.x] = false

			//fmt.Printf("carvePathToRegion L形: 从(%d,%d)开凿到(%d,%d)\n", sx, sy, cur.x, cur.y)
			return true
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
	return false
}

func AStar(m *AStarMap, start, end Point) []Point {

	// 边界裁剪
	if start.X < 0 {
		start.X = 0
	}
	if start.Y < 0 {
		start.Y = 0
	}
	if start.X >= m.cols {
		start.X = m.cols - 1
	}
	if start.Y >= m.rows {
		start.Y = m.rows - 1
	}
	if end.X < 0 {
		end.X = 0
	}
	if end.Y < 0 {
		end.Y = 0
	}
	if end.X >= m.cols {
		end.X = m.cols - 1
	}
	if end.Y >= m.rows {
		end.Y = m.rows - 1
	}

	// 先获取终点的连通区域
	endRegion := m.getConnectedRegion(end.X, end.Y)
	if len(endRegion) == 0 {
		fmt.Println("终点在障碍物上且无法修正")
		return nil
	}

	// 起点不在终点连通区域时开凿通路
	if !endRegion[[2]int{start.X, start.Y}] {
		fmt.Println("起点不在终点连通区域，开凿通路")
		if !m.carvePathToRegion(start.X, start.Y, endRegion) {
			fmt.Println("无法开凿通路")
			return nil
		}
	}

	//fmt.Printf("修正后 start=(%d,%d) end=(%d,%d) startObstacle=%v endObstacle=%v\n",
	//	start.X, start.Y, end.X, end.Y,
	//	m.obstacle[start.Y][start.X], m.obstacle[end.Y][end.X])

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
	//fmt.Printf("初始化pq, len=%d\n", pq.Len()) // 加这行
	heap.Init(pq)
	//fmt.Printf("Init后pq, len=%d\n", pq.Len()) // 加这行

	loopCount := 0
	for pq.Len() > 0 {
		loopCount++
		cur := heap.Pop(pq).(*Node)
		cur.InOpen = false
		cur.InClosed = true

		if cur.X == end.X && cur.Y == end.Y {

			//fmt.Printf("找到路径,循环次数=%d\n", loopCount)
			return buildPath(cur)
		}

		for _, d := range dirs {
			nx, ny := cur.X+d.X, cur.Y+d.Y
			if nx < 0 || ny < 0 || nx >= m.cols || ny >= m.rows {
				//fmt.Printf("邻居(%d,%d)越界\n", nx, ny)
				continue
			}
			if m.obstacle[ny][nx] {
				//fmt.Printf("邻居(%d,%d)是障碍\n", nx, ny)
				continue
			}
			// 斜角移动检查两侧
			if d.X != 0 && d.Y != 0 {
				if m.obstacle[cur.Y][nx] || m.obstacle[ny][cur.X] {
					//fmt.Printf("邻居(%d,%d)斜角被切角阻挡\n", nx, ny)
					continue
				}
			}
			//fmt.Printf("邻居(%d,%d)可通行,加入队列\n", nx, ny)
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
	fmt.Printf("未找到路径,循环次数=%d\n", loopCount)
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
	//fmt.Printf("path=%v, len=%d, nil=%v\n", raw, len(raw), raw == nil)
	if raw == nil {
		return nil
	}

	// Douglas-Peucker简化
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
	interpolated := interpolatePath(centered, info.MaxDist)
	// 验证插值结果
	//for i := 1; i < len(interpolated); i++ {
	//	dx := interpolated[i].X - interpolated[i-1].X
	//	dy := interpolated[i].Y - interpolated[i-1].Y
	//	dist := math.Sqrt(float64(dx*dx + dy*dy))
	//	if dist > 20 {
	//		fmt.Printf("插值后仍有超距: 点%d→点%d 距离=%.1f\n", i-1, i, dist)
	//	}
	//}

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

// 全局context，地图丢失时取消所有操作
var (
	globalCtx    context.Context
	globalCancel context.CancelFunc
)

// 启动地图检测线程
func StartMapMonitor(bigMapPath string, cancelFunc context.CancelFunc) {
	go func() {
		for {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.5)
			if x == -1 && y == -1 {
				//println("==========循环获取坐标:", x, y)
				//println("==========循环获取坐标出错")
				fmt.Println("地图丢失，终止所有操作")
				cancelFunc() // 取消所有操作
				return
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
}

func FollowPath(ctx context.Context, path []Point, getCurrentPos func() Point) info.WayFindState {
	if len(path) == 0 {
		println("1")
		return info.STATE_CDT_Useless // 假如路径为空当坐标异常处理，触发重新寻路
	}
	defer myMotion.StopMove()

	detachmentTime := 0

	for _, waypoint := range path {
		tryCount := 0

		for tryCount < info.PathfindingRetryCount {
			// 检查是否被取消
			select {
			case <-ctx.Done():
				fmt.Println("ctx被取消，原因:", ctx.Err())
				return info.STATE_Loss_Map
			default:
			}

			cur := getCurrentPos()

			dx := waypoint.X - cur.X
			dy := waypoint.Y - cur.Y
			dist := math.Sqrt(float64(dx*dx + dy*dy))

			//fmt.Printf("tryCount=%d cur=(%d,%d) waypoint=(%d,%d) dist=%.1f\n",
			//	tryCount, cur.X, cur.Y, waypoint.X, waypoint.Y, dist)

			if dist < info.SuccessDist {
				break
			}

			if dist > info.ErrDist {
				//println("2")
				holdX, holdY := myMotion.CalcHoldPoint(cur.X, cur.Y, waypoint.X, waypoint.Y)
				motion.Click(640, 360, 0, 0)
				time.Sleep(500 * time.Millisecond)
				motion.Click(holdX, holdY, 0, 0)
				return info.STATE_CDT_Useless
			}

			holdX, holdY := myMotion.CalcHoldPoint(cur.X, cur.Y, waypoint.X, waypoint.Y)
			//fmt.Printf("发送触摸: holdX=%d holdY=%d\n", holdX, holdY)
			myMotion.StartMoveXY(holdX, holdY)

			tryCount++
			time.Sleep(info.NavigationInterval)
		}

		if tryCount >= info.PathfindingRetryCount {
			detachmentTime++
		}

		if detachmentTime >= info.TimeoutCount {
			return info.STATE_Movement_timeout
		}
	}

	return info.STATE_Done
}

// 封装寻路+异常处理，支持重试
func NavigateTo(bigMapPath string, astarMap *AStarMap, end Point,
	getCurrentPos func() Point,
	stopSkill func(),
	startSkill func(),
) info.WayFindState {
	trycount := 0
	for {
		ctx, cancel := context.WithCancel(context.Background())
		// 启动地图监测
		StartMapMonitor(bigMapPath, cancel)

		// 获取当前坐标
		x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.5)
		//println("==========获取当前坐标:", x, y)
		if x == -1 && y == -1 {
			println("\n==========获取当前坐标出错=========\n")
			cancel()

			stopSkill()                        // 停止
			state := handleLossMap(bigMapPath) // 异常处理
			startSkill()                       // 重启

			trycount++
			if trycount > 5 {
				state = info.STATE_Fixed_Failed
			}
			if state != info.STATE_Fixed_succes {
				return state
			}

			continue // 处理完后重试
		}

		start := Point{X: x, Y: y}

		//path := AStar(astarMap, start, end)
		path := FindPath(astarMap, start, end)
		if path == nil || len(path) == 0 {
			fmt.Printf("路径为空，等待后重试,path长度:%d\n", len(path))
			time.Sleep(500 * time.Millisecond)
			trycount++
			if trycount > 5 {
				cancel()
				println("3")
				return info.STATE_CDT_Useless
			}
			cancel()
			continue // 重新循环重新寻路
		}

		state := FollowPath(ctx, path, getCurrentPos)
		cancel() // 停止地图监测

		switch state {
		case info.STATE_Done:
			return info.STATE_Done

		case info.STATE_Loss_Map:
			fmt.Println("地图丢失，进入异常处理")
			stopSkill()
			state = handleLossMap(bigMapPath)
			if state != info.STATE_Fixed_succes {
				return state
			}
			startSkill() // 重启
			// 异常处理完后继续循环重试

		case info.STATE_Movement_timeout:
			fmt.Println("寻路超时，重新规划路径")
			// 超时,一般时按键冲突，四周点两下然后重新寻路
			x1, y1, x2, y2 := myMotion.RandomPoint()
			for i := 0; i < 2; i++ {
				if i == 0 {
					motion.Click(x1, y1, 0, 0)
				} else {
					motion.Click(x2, y2, 0, 0)
				}
				time.Sleep(500 * time.Millisecond)
			}
		case info.STATE_CDT_Useless:
			//坐标异常,直接重新寻路
			fmt.Println("坐标异常，等待后重试")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 地图丢失的异常处理
func handleLossMap(bigMapPath string) info.WayFindState {
	fmt.Println("执行异常处理")

	time.Sleep(2 * time.Second)

	if MyOpenCV.ListColorsCmp(info.IF.IF_TpInterface[:], 0.8) { //看看是不是进到传送界面
		for i := 0; i < 3; i++ {
			motion.Click(105, 40, 0, 0)                           //退出传送界面
			if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //检测到跑图界面
				return info.STATE_Fixed_succes
			}
			if i == 2 {
				return info.STATE_Fixed_Failed
			} //没检测到
			time.Sleep(2000 * time.Millisecond)
		}
	}

	for i := 0; i < 3; i++ {
		inBattle := MyOpenCV.If_BattleInterface(0.85)
		if !inBattle { // 没检测出来战斗界面, 有可能是图片识别错误也可能是网络延迟,随便点两下
			// 生成一个随机角度
			fmt.Printf("未检测到战斗界面")
			if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //先判断是不是已经在跑图界面
				return info.STATE_Fixed_succes
			}

			if MyOpenCV.ListColorsCmp(info.IF.If_Backbutton[:], 0.7) { //如果有返回按钮
				motion.Click(info.MiscPoint.BackButton.X, info.MiscPoint.BackButton.Y, 0, 0)
				time.Sleep(1500 * time.Millisecond)
				if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
					return info.STATE_Fixed_succes
				}
			}

			if MyOpenCV.ColorCmp(info.IF.If_BattleFieldRole.ColorsCmp, 0.8) {
				motion.Click(644, 652, 0, 0) //点战场角色界面按钮,这里其实是退出战场角色界面
				time.Sleep(1500 * time.Millisecond)
				if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
					return info.STATE_Fixed_succes
				}
			}

			//如果还不在那可能是打开钻石金币弹窗了点一下战场界面按钮
			motion.Click(644, 652, 0, 0) //点战场角色界面按钮,关闭金币弹窗
			time.Sleep(1500 * time.Millisecond)
			if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
				return info.STATE_Fixed_succes
			}

			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6) //二次检测
			if !(x == -1 && y == -1) {
				fmt.Println("成功返回跑图界面")
				return info.STATE_Fixed_succes
			}
		} else {
			//进入战斗界面退出战斗界面
			for i := 0; i < 2; i++ {
				inBattle = MyOpenCV.If_BattleInterface(0.85)
				time.Sleep(2000 * time.Millisecond)
				if inBattle { //退出战斗
					println("检测到战斗界面,退出战斗")
					for i := 0; i < 2; i++ {
						MyOpenCV.If_LoadingInterface(10)
						if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
							return info.STATE_Fixed_succes
						}
						motion.Click(1178, 42, 0, 0)
						if MyOpenCV.ColorCmp(info.IF.If_Pause.ColorsCmp, 0.8) {
							time.Sleep(1 * time.Second)
							motion.Click(510, 660, 0, 0)
							time.Sleep(1 * time.Second)
							motion.Click(721, 433, 0, 0)
						}

						MyOpenCV.If_LoadingInterface(10)
						time.Sleep(1 * time.Second)

						if MyOpenCV.ColorCmp(info.IF.If_Pause.ColorsCmp, 0.8) { //如果在暂停页面
							motion.Click(510, 660, 0, 0)
							time.Sleep(1 * time.Second)
							motion.Click(721, 433, 0, 0)
							time.Sleep(1500 * time.Millisecond)
							MyOpenCV.If_LoadingInterface(10)
							if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
								return info.STATE_Fixed_succes
							}
						}
						if MyOpenCV.If_escape() { //如果在逃跑页面
							motion.Click(721, 433, 0, 0)
							time.Sleep(1500 * time.Millisecond)
							MyOpenCV.If_LoadingInterface(10)
							if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
								return info.STATE_Fixed_succes
							}
						}

						time.Sleep(1000 * time.Millisecond)
					}
				}
				if i == 1 { //没返回到跑图界面
					return info.STATE_Fixed_Failed
				}
				time.Sleep(2 * time.Second)
			}
		}
		time.Sleep(2 * time.Second)
	}

	if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //返回到跑图界面
		fmt.Println("成功返回跑图界面")
		return info.STATE_Fixed_succes
	}

	return info.STATE_Fixed_Failed
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

func YoloFind(yoloPtr *yolo.Yolo, bigMapPath string) info.WayFindState {
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
				fmt.Errorf("yolo找怪时进入战斗,并且未退出战斗,重新开始本阶段")
				return info.STATE_Fixed_Failed

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
	return info.STATE_Done
}

func Handle() { //测试自动纠错
	handleLossMap(info_chapter.Ch1.ChapterImg_path)
}
