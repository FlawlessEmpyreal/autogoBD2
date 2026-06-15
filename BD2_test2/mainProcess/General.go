package mainProcess

import (
	"app/MyOpenCV"
	"app/aStar"
	"app/info"
	"context"
	"errors"
	"fmt"
	"image"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/Dasongzi1366/AutoGo/motion"
	"github.com/Dasongzi1366/AutoGo/opencv"
	"github.com/Dasongzi1366/AutoGo/yolo"
)

// 进入传送地图
func Chapter_Tp(
	chapterMap_Point []info.StructColorCmp, //找地图时用的找色
	Chapter_TpPoint image.Point, //传送点坐标
) (info.RecoveryAction, error) {
	for i := 0; i < 60; i++ {
		if MyOpenCV.If_LoadingInterface(0.8) {
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
		if i == 60 {
			println("卡加载60秒,直接退出")
			os.Exit(0)
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		if MyOpenCV.ListColorsCmp(info.IF.IF_TpInterface[:], 0.7) { //查看是否进入传送地图
			break
		}
		motion.Click(1089, 444, 0, 0) //点击传送阵
		time.Sleep(2 * time.Second)
		if MyOpenCV.ListColorsCmp(info.IF.IF_TpInterface[:], 0.7) { //查看是否进入传送地图
			break
		}
		motion.Click(870, 567, 0, 0) //如果在传送阵周围
		time.Sleep(2 * time.Second)
		if MyOpenCV.ListColorsCmp(info.IF.IF_TpInterface[:], 0.7) { //查看是否进入传送地图
			break
		}

		//如果在传送阵周围但是没有互动键
		theta := r.Float64() * 2 * math.Pi

		// 计算第一次点击的坐标
		x := 640 + int(float64(175)*math.Cos(theta))
		y := 360 + int(float64(175)*math.Sin(theta))

		motion.Click(x, y, 0, 0)

		if i == 9 {
			return info.GoToRunMapInterface, fmt.Errorf("未进入传送地图,回跑图界面,重试")
		}
		time.Sleep(2 * time.Second)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		if MyOpenCV.ColorCmp(info.IF.If_GenerateTP.ColorsCmp, 0.8) { //如果需要生成魔法阵
			motion.Click(735, 443, 0, 0)
		} else {
			break
		}
	}

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ { //来回找三次
		state_ := false
		for i := 0; i < 8; i++ { //将地图往左点
			if !MyOpenCV.ColorCmp(info.IF.If_TpLeftButton.ColorsCmp, 0.8) { //没按钮也不用点了
				break
			}
			if MyOpenCV.ColorCmp(chapterMap_Point, 0.7) {
				state_ = true
				break
			}
			motion.Click(300, 345, 0, 0)
			time.Sleep(1000 * time.Millisecond)
		}
		if state_ {
			break
		}

		for i := 0; i < 8; i++ { //将地图往右点
			if !MyOpenCV.ColorCmp(info.IF.If_TpRightButton.ColorsCmp, 0.8) { //没按钮也不用点了
				break
			}
			if MyOpenCV.ColorCmp(chapterMap_Point, 0.7) {
				state_ = true
				break
			}
			motion.Click(978, 345, 0, 0)
			time.Sleep(1000 * time.Millisecond)
		}
		if state_ {
			break
		}
		if i == 3 {
			return info.RetryStage, fmt.Errorf("未找到要传送的目标地图,重试本阶段")
		}
	}

	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		if MyOpenCV.ColorCmp(chapterMap_Point, 0.8) {
			motion.Click(Chapter_TpPoint.X, Chapter_TpPoint.Y, 0, 0)
		}

		time.Sleep(1 * time.Second)
		for i := 0; i < 60; i++ {
			if MyOpenCV.ColorCmp(info.IF.If_GenerateTP.ColorsCmp, 0.7) {
				motion.Click(734, 442, 0, 0) //点击确定
			}
			if !MyOpenCV.ColorCmp(info.IF.If_GenerateTP.ColorsCmp, 0.7) {
				break //确定没卡顿
			}
			time.Sleep(1000 * time.Millisecond)
			if i == 60 {
				println("确定生成传送阵时卡了60秒,直接退出")
				os.Exit(0)
			}
		}

		if !MyOpenCV.ListColorsCmp(info.IF.IF_TpInterface[:], 0.7) { //查看是否退出传送界面
			break
		} else {
			motion.Click(105, 40, 0, 0)
		}
		if i == 3 {
			return info.RetryStage, fmt.Errorf("未找到要传送的目标地图,重试本阶段")
		}
	}
	MyOpenCV.WaitLoading(20)
	time.Sleep(2 * time.Second)
	println("进入地图")
	return info.StageDone, nil
}

func FindChapter(chapterImg_path, type_ string) (info.RecoveryAction, error) { //必须在章节界面
	imgCSI := MyOpenCV.ChapterSelectInterface()
	imgCSB := MyOpenCV.GetByte(info.ChapterSelectButtonImg_path)

	chapterImg := MyOpenCV.GetByte(chapterImg_path)
	if chapterImg == nil {
		println("章节选择图片读取错误,请检查路径权限或文件是否存在")
		os.Exit(0)
	}
	if imgCSB == nil {
		println("章节按钮图片读取错误,请检查路径权限或文件是否存在")
		os.Exit(0)
	}
	if imgCSI == nil {
		println("章节检测图片读取错误,请检查路径权限或文件是否存在")
		os.Exit(0)
	}

	motion.Click(640, 360, 0, 0) //刷新运动状态
	time.Sleep(500 * time.Millisecond)
	motion.Click(640, 360, 0, 0)

	for i := 0; i < 3; i++ { //进入章节选择
		x, y := opencv.FindImage(519, 614, 651, 691, imgCSB, false, true, 0.5, 0)
		if x == -1 || y == -1 {
			if i == 2 {
				return info.GoToRunMapInterface, errors.New("未找到章节选择按钮,退到跑图界面重试")
			}
			time.Sleep(1 * time.Second)
			continue
		}
		motion.Click(x, y, 0, 0)
		time.Sleep(1 * time.Second)
		x, y = opencv.FindImage(37, 670, 1202, 720, imgCSI, false, false, 0.5, 0)
		//println(x, y)
		if x != -1 && y != -1 {
			break
		} else {
			time.Sleep(1 * time.Second)
		}
		if i == 2 {
			return info.GoToRunMapInterface, errors.New("没进入章节选择界面,退到跑图界面重试")
		}
	}

	for i := 0; i < 3; i++ { //选择主线和支线
		if type_ == "main" {
			motion.Click(info.MiscPoint.MainChapterChoose.X, info.MiscPoint.MainChapterChoose.Y, 0, 0)
			time.Sleep(300 * time.Millisecond)
		} else if type_ == "branch" {
			motion.Click(info.MiscPoint.BranchChapterChoose.X, info.MiscPoint.BranchChapterChoose.Y, 0, 0)
			time.Sleep(300 * time.Millisecond)
		} else {
			println("type_的值只能是main或者branch")
			os.Exit(0)
		}
	}
	time.Sleep(2 * time.Second)

	x, y := opencv.FindImage(1, 585, 1276, 677, chapterImg, false, false, 0.8, 0) //先找本页面有没有目标章节,如果有直接点击然后退出
	if x != -1 && y != -1 {
		motion.Click(x+10, y+5, 0, 0)

		time.Sleep(1000 * time.Millisecond)

		for i := 0; i < 6; i++ { //确认退出选章节界面
			x, y := opencv.FindImage(1, 585, 1276, 677, imgCSI, false, false, 0.8, 0)
			if x == -1 && y == -1 {
				break
			}
		}

		for i := 0; i < 60; i++ { //持续检测是否在加载界面
			if !MyOpenCV.If_LoadingInterface(0.8) {
				break
			}
			if i == 19 {
				println("卡在加载界面一分钟,退出进程")
				os.Exit(0)
			}
			time.Sleep(1 * time.Second)
		}
		return info.StageDone, nil
	}

	for i := 0; i < 4; i++ { //拉到最左边
		motion.Swipe(190, 636, 1170, 636, 500, 0, 0)
		time.Sleep(700 * time.Millisecond)
	}

	for i := 0; i < 6; i++ { //找目标章节
		x, y := opencv.FindImage(1, 585, 1276, 677, chapterImg, false, false, 0.8, 0)
		if x != -1 && y != -1 {
			motion.Click(x+10, y+5, 0, 0)
			break
		}
		motion.Swipe(937, 636, 549, 636, 1000, 0, 0)
		time.Sleep(2000 * time.Millisecond)
		if i == 5 {
			return info.StageDone, fmt.Errorf("寻找章节失败未找到章节,重新执行本阶段")
		}
	}

	for i := 0; i < 6; i++ { //确认退出选章节界面
		x, y := opencv.FindImage(1, 585, 1276, 677, imgCSI, false, false, 0.8, 0)
		if x == -1 && y == -1 {
			break
		}
	}

	//for i := 0; i < 20; i++ { //持续检测是否进入加载界面
	//	if MyOpenCV.If_LoadingInterface(0.8) {
	//		break
	//	}
	//	if i == 19 {
	//		return GoBattleInterface, errors.New("未找到目标章节,可能是主线或支线章节未切换或网络延迟,退到跑图界面重试")
	//		os.Exit(0)
	//	}
	//	time.Sleep(500 * time.Millisecond)
	//}
	for i := 0; i < 60; i++ { //持续检测是否在加载界面
		if !MyOpenCV.If_LoadingInterface(0.8) {
			break
		}
		if i == 19 {
			println("卡在加载界面一分钟,退出进程")
			os.Exit(0)
		}
		time.Sleep(1 * time.Second)
	}
	time.Sleep(2000 * time.Millisecond)

	//检测是否进入章节有bug
	//for i := 0; i < 3; i++ {
	//	if MyOpenCV.ColorCmp(info.ChapterSelectDetect, 0.85) { // 如果还没进入章节
	//		time.Sleep(1 * time.Second)
	//		if i == 2 { // 已经是最后一次尝试了
	//			println("未进入章节,暂时未做异常处理,退出")
	//			os.Exit(0)
	//		}
	//		continue // 跳过本次循环剩余部分，继续下一次检测
	//	}
	//
	//	// ColorCmp 返回 false，代表已成功进入章节
	//	println("已进入章节")
	//	break
	//}
	return info.StageDone, nil
}

func ChapterRun(
	bigMapPath, //大地图路径
	bin_mapPath string, //二值化地图路径
	MonsterLocation []image.Point, //怪物位置
	YoloLabel string,
	type_ string, //使用哪种模型
) (info.RecoveryAction, error) {

	astarMap, err := aStar.LoadObstacleMap(bin_mapPath)
	if err != nil {
		return info.AbortAll, fmt.Errorf("地图加载二进制地图失败,退出: %w", err)
	}

	var ifRunMap bool //检查是否在跑图界面
	for i := 0; i < 20; i++ {
		ifRunMap = MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.85)
		if ifRunMap {
			break
		}
		if i == 19 {
			return info.GoToRunMapInterface, fmt.Errorf("未进入跑图界面,返回到跑图界面")
		}
		time.Sleep(1 * time.Second)
	}

	var yoloModel *yolo.Yolo
	if type_ == "main" {
		yoloModel = yolo.New("v5", 4, info.YoloParamPath_Main, info.YoloBinPath_Main, YoloLabel)
		if yoloModel == nil {
			fmt.Println("模型加载失败,直接退出")
			os.Exit(0)
		}
		defer yoloModel.Close()
		fmt.Println("模型加载成功,标签:", YoloLabel)
	} else if type_ == "branch" {
		yoloModel = yolo.New("v5", 4, info.YoloParamPath_Branch, info.YoloBinPath_Branch, YoloLabel)
		if yoloModel == nil {
			fmt.Println("模型加载失败,直接退出")
			os.Exit(0)
		}
		defer yoloModel.Close()
		fmt.Println("模型加载成功,标签:", YoloLabel)
	}

	//var mu sync.Mutex
	//go func() {
	//	// 持续发送坐标
	//	for {
	//		x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
	//		mu.Lock()
	//		sender.SendCoord(x, y)
	//		mu.Unlock()
	//		time.Sleep(200 * time.Millisecond)
	//	}
	//}()
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//go func(ctx context.Context) {
	//	for {
	//		select {
	//		case <-ctx.Done(): // 监听到上下文被取消，主动退出
	//			fmt.Println("收到Context信号，释放技能的协程退出")
	//			return
	//		default:
	//			if info.Accelerate {
	//				motion.Click(info.BP.Accelerate.X, info.BP.Accelerate.Y, 2, 0) //加速
	//				time.Sleep(1 * time.Second)
	//			}
	//			if info.Subdue {
	//				motion.Click(info.BP.Subdue.X, info.BP.Subdue.Y, 2, 0) //压制
	//				time.Sleep(1 * time.Second)
	//			}
	//			if info.Stealth {
	//				motion.Click(info.BP.Stealth.X, info.BP.Stealth.Y, 2, 0) //隐身
	//				time.Sleep(1 * time.Second)
	//			}
	//
	//		}
	//	}
	//}(ctx)

	var ctx2 context.Context
	var cancel2 context.CancelFunc
	var waitSkill func()
	var stopSkill func()
	var startSkill func()

	stopSkill = func() {
		if cancel2 != nil {
			cancel2()
			waitSkill()
			cancel2 = nil
			//fmt.Println("SkillLoop已停止")
		}
	}

	startSkill = func() {
		ctx2, cancel2 = context.WithCancel(context.Background())
		waitSkill = SkillLoop(ctx2)
		//fmt.Println("SkillLoop已重启")
	}

	defer stopSkill()

	time.Sleep(1 * time.Second)
	startSkill()

	for i := 0; i < len(MonsterLocation); i++ {
		if info.STATE_Done != aStar.NavigateTo(bigMapPath, astarMap,
			aStar.Point{X: MonsterLocation[i].X, Y: MonsterLocation[i].Y},
			func() aStar.Point {
				x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
				return aStar.Point{X: x, Y: y}
			},
			stopSkill, startSkill,
		) {
			return info.RetryStage, fmt.Errorf("第%d个怪物寻路失败,重试本阶段", i)
		}

		//一直点击检测到的第一个目标位置，直到被消灭
		if info.STATE_Done != aStar.YoloFind(yoloModel, bigMapPath) {
			return info.RetryStage, fmt.Errorf("yolo找怪时小地图丢失,重新开始本阶段")
		}

		time.Sleep(500 * time.Millisecond)
	}
	////=========================找第一个怪===================================
	//if info.STATE_Done != aStar.NavigateTo(bigMapPath, astarMap,
	//	aStar.Point{X: MonsterLocation[0].X, Y: MonsterLocation[0].Y},
	//	func() aStar.Point {
	//		x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
	//		return aStar.Point{X: x, Y: y}
	//	},
	//) {
	//	return info.RetryStage, fmt.Errorf("第一个怪物寻路失败,重试本阶段")
	//}
	//
	////一直点击检测到的第一个目标位置，直到被消灭
	//if info.STATE_Done != aStar.YoloFind(yolo, bigMapPath){
	//	return info.RetryStage,fmt.Errorf("yolo找怪时小地图丢失,重新开始本阶段")
	//}
	////=========================找第一个怪===================================
	//time.Sleep(500 * time.Millisecond)
	////=========================找第二个怪===================================
	//if info.STATE_Done != aStar.NavigateTo(bigMapPath, astarMap,
	//	aStar.Point{X: MonsterLocation[1].X, Y: MonsterLocation[1].Y},
	//	func() aStar.Point {
	//		x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
	//		return aStar.Point{X: x, Y: y}
	//	},
	//) {
	//	return info.RetryStage, fmt.Errorf("第二个怪物寻路失败,重试本阶段")
	//}
	//
	////一直点击检测到的第一个目标位置，直到被消灭
	//if info.STATE_Done != aStar.YoloFind(yolo, bigMapPath){
	//	return info.RetryStage,fmt.Errorf("yolo找怪时小地图丢失,重新开始本阶段")
	//}
	////=========================找第二个怪===================================
	//time.Sleep(500 * time.Millisecond)
	////=========================找第三个怪===================================
	//if info.STATE_Done != aStar.NavigateTo(bigMapPath, astarMap,
	//	aStar.Point{X: MonsterLocation[2].X, Y: MonsterLocation[2].Y},
	//	func() aStar.Point {
	//		x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
	//		return aStar.Point{X: x, Y: y}
	//	},
	//) {
	//	return info.RetryStage, fmt.Errorf("第三个怪物寻路失败,重试本阶段")
	//}
	//
	////一直点击检测到的第一个目标位置，直到被消灭
	//if info.STATE_Done != aStar.YoloFind(yolo, bigMapPath){
	//	return info.RetryStage,fmt.Errorf("yolo找怪时小地图丢失,重新开始本阶段")
	//}
	////=========================找第三个怪===================================
	//time.Sleep(500 * time.Millisecond)
	////=========================找第四个怪===================================
	//if info.STATE_Done != aStar.NavigateTo(bigMapPath, astarMap,
	//	aStar.Point{X: MonsterLocation[3].X, Y: MonsterLocation[3].Y},
	//	func() aStar.Point {
	//		x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
	//		return aStar.Point{X: x, Y: y}
	//	},
	//) {
	//	return info.RetryStage, fmt.Errorf("第四个怪物寻路失败,重试本阶段")
	//}
	//
	////一直点击检测到的第一个目标位置，直到被消灭
	//if info.STATE_Done != aStar.YoloFind(yolo, bigMapPath){
	//	return info.RetryStage,fmt.Errorf("yolo找怪时小地图丢失,重新开始本阶段")
	//}
	////=========================找第四个怪===================================
	//time.Sleep(500 * time.Millisecond)
	////=========================找第五个怪===================================
	//if info.STATE_Done != aStar.NavigateTo(bigMapPath, astarMap,
	//	aStar.Point{X: MonsterLocation[4].X, Y: MonsterLocation[4].Y},
	//	func() aStar.Point {
	//		x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
	//		return aStar.Point{X: x, Y: y}
	//	},
	//) {
	//	return info.RetryStage, fmt.Errorf("第五个怪物寻路失败,重试本阶段")
	//}
	//
	////一直点击检测到的第一个目标位置，直到被消灭
	//if info.STATE_Done != aStar.YoloFind(yolo, bigMapPath){
	//	return info.RetryStage,fmt.Errorf("yolo找怪时小地图丢失,重新开始本阶段")
	//}

	time.Sleep(1 * time.Second)
	return info.StageDone, nil
}

func interruptibleSleep(ctx context.Context, d time.Duration) {
	t := time.NewTimer(d)
	defer t.Stop() // 确保定时器资源被回收
	select {
	case <-t.C:
		// 正常等待
	case <-ctx.Done():
		return
		// 收到取消信号，立即返回
	}
}

// safeExecute 在执行实际操作前检查上下文是否已被取消
func safeExecute(ctx context.Context, action func()) bool {
	select {
	case <-ctx.Done():

		return false // 已取消，不执行
	default:
		action()
		return true // 未取消，执行动作
	}
}

// 启动技能循环线程
func SkillLoop(ctx context.Context) func() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer func() {
			//fmt.Println("SkillLoop goroutine 真正退出")
			wg.Done()
		}()
		for {
			select {
			case <-ctx.Done():
				//fmt.Println("SkillLoop 收到取消信号")
				//fmt.Println("收到Context信号，释放技能的协程退出")
				return
			default:
				if info.Accelerate {
					interruptibleSleep(ctx, 100*time.Millisecond)
					if !safeExecute(ctx, func() {
						motion.Click(info.BP.Accelerate.X, info.BP.Accelerate.Y, 2, 0)
					}) {
						return
					}
					interruptibleSleep(ctx, 1*time.Second)
				}
				if info.Subdue {
					if !safeExecute(ctx, func() {
						motion.Click(info.BP.Subdue.X, info.BP.Subdue.Y, 2, 0)
					}) {
						return
					}
					interruptibleSleep(ctx, 1*time.Second)
				}
				if info.Stealth {
					if !safeExecute(ctx, func() {
						motion.Click(info.BP.Stealth.X, info.BP.Stealth.Y, 2, 0)
					}) {
						return
					}
					interruptibleSleep(ctx, 1*time.Second)
				}
			}
		}
	}(ctx)
	// 返回等待函数
	return func() {
		//fmt.Println("开始等待SkillLoop退出")
		wg.Wait()
		//fmt.Println("SkillLoop已退出")
	}
}
