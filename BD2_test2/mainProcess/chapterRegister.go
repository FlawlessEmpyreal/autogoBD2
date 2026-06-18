package mainProcess

//注册第一章节

import (
	"app/info"
)

var ChapterHandlers = []func(*Controller){
	RegChapter1, // 第1章
	RegChapter2, // 第2章
}

func RegisterChapters(Ctrl *Controller) {
	for i, handler := range ChapterHandlers {
		if info.RegCh[i] {
			handler(Ctrl)
		}
	}
}

//func RegisterChapters(Ctrl *Controller) {
//	// -------- Chapter 1 --------
//	if info.RegCh1 {
//		RegChapter1(Ctrl)
//	}
//	//━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//	// -------- Chapter 2 --------
//	if info.RegCh2 {
//		RegChapter2(Ctrl)
//	}
//	//━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//}

//func Chapter1() {
//
//	FindChapter(info.Ch1.ChapterImg_path)
//
//	Chapter_Tp(
//		info.Ch1.MapConfig[0].MapFind.MapColorsCmp,
//		info.Ch1.MapConfig[0].TpPoint,
//	)
//
//	ChapterRun(
//		"chapter1_1",
//		info.Ch1.MapConfig[0].BigMapPath,
//		info.Ch1.MapConfig[0].Bin_mapPath,
//		info.Ch1.MapConfig[0].MonsterLocation,
//	)
//
//	Chapter_Tp(
//		info.Ch1.MapConfig[1].MapFind.MapColorsCmp,
//		info.Ch1.MapConfig[1].TpPoint,
//	)
//
//	ChapterRun(
//		"chapter1_2",
//		info.Ch1.MapConfig[1].BigMapPath,
//		info.Ch1.MapConfig[1].Bin_mapPath,
//		info.Ch1.MapConfig[1].MonsterLocation,
//	)
//
//}

//func chapterRun(
//	RunName , 			//标识，仅说明
//	bigMapPath, 	    //大地图路径
//	bin_mapPath string, //二值化地图路径
//	MonsterLocation []image.Point,  //怪物位置
//	) {
//
//	println("即将运行地图:",RunName)
//	//obstacle, _ := aStar.LoadObstacleMap(bin_mapPath)
//	var inBattle bool
//	for i := 0; i < 20; i++ {
//		inBattle = MyOpenCV.ColorCmp(info.If_Map, 0.85)
//		if inBattle {
//			break
//		}
//		if i == 19 {
//			println("未进入跑图界面,未做异常处理")
//			os.Exit(0)
//		}
//		time.Sleep(1 * time.Second)
//	}
//
//	yolo := yolo.New("v5", 4, info.YoloParamPath, info.YoloBinPath, info.Yolo_labels)
//	if yolo == nil {
//		fmt.Println("模型加载失败")
//		return
//	}
//	defer yolo.Close()
//	fmt.Println("模型加载成功")
//
//	//var mu sync.Mutex
//	//go func() {
//	//	// 持续发送坐标
//	//	for {
//	//		x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//	//		mu.Lock()
//	//		sender.SendCoord(x, y)
//	//		mu.Unlock()
//	//		time.Sleep(200 * time.Millisecond)
//	//	}
//	//}()
//	ctx, cancel := context.WithCancel(context.Background())
//
//	go func(ctx context.Context) {
//		for {
//			select {
//			case <-ctx.Done(): // 监听到上下文被取消，主动退出
//				fmt.Println("收到Context信号，技能协程退出")
//				return
//			default:
//				if accelerate {
//					motion.Click(1000, 650, 2, 0) //加速
//					time.Sleep(1 * time.Second)
//				}
//				if Subdue {
//					motion.Click(967, 561, 2, 0) //压制
//					time.Sleep(1 * time.Second)
//				}
//				if stealth {
//					motion.Click(1000, 480, 2, 0) //隐身
//					time.Sleep(1 * time.Second)
//				}
//
//			}
//		}
//	}(ctx)
//
//	time.Sleep(1 * time.Second)
//	//=========================找第一个怪===================================
//	aStar.NavigateTo(bigMapPath, bin_mapPath,
//		aStar.Point{X: 256, Y: 397},
//		func() aStar.Point {
//			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//			return aStar.Point{X: x, Y: y}
//		},
//	)
//
//	//一直点击检测到的第一个目标位置，直到被消灭
//	aStar.YoloFind(yolo, bigMapPath)
//	//=========================找第一个怪===================================
//	time.Sleep(500 * time.Millisecond)
//	//=========================找第二个怪===================================
//	aStar.NavigateTo(bigMapPath, bin_mapPath,
//		aStar.Point{X: 338, Y: 380},
//		func() aStar.Point {
//			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//			return aStar.Point{X: x, Y: y}
//		},
//	)
//
//	//一直点击检测到的第一个目标位置，直到被消灭
//	aStar.YoloFind(yolo, bigMapPath)
//	//=========================找第二个怪===================================
//	time.Sleep(500 * time.Millisecond)
//	//=========================找第三个怪===================================
//	aStar.NavigateTo(bigMapPath, bin_mapPath,
//		aStar.Point{X: 420, Y: 332},
//		func() aStar.Point {
//			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//			return aStar.Point{X: x, Y: y}
//		},
//	)
//
//	//一直点击检测到的第一个目标位置，直到被消灭
//	aStar.YoloFind(yolo, bigMapPath)
//	//=========================找第三个怪===================================
//	time.Sleep(500 * time.Millisecond)
//	//=========================找第四个怪===================================
//	aStar.NavigateTo(bigMapPath, bin_mapPath,
//		aStar.Point{X: 304, Y: 233},
//		func() aStar.Point {
//			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//			return aStar.Point{X: x, Y: y}
//		},
//	)
//
//	//一直点击检测到的第一个目标位置，直到被消灭
//	aStar.YoloFind(yolo, bigMapPath)
//	//=========================找第四个怪===================================
//	time.Sleep(500 * time.Millisecond)
//	//=========================找第五个怪===================================
//	aStar.NavigateTo(bigMapPath, bin_mapPath,
//		aStar.Point{X: 175, Y: 191},
//		func() aStar.Point {
//			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//			return aStar.Point{X: x, Y: y}
//		},
//	)
//
//	//一直点击检测到的第一个目标位置，直到被消灭
//	aStar.YoloFind(yolo, bigMapPath)
//	cancel()
//	time.Sleep(2 * time.Second)
//}

//func chapter1_2(bigMapPath, bin_mapPath, yolo_parapath, yolo_binpath, yolo_labels string) {
//
//	//obstacle, _ := aStar.LoadObstacleMap(bin_mapPath)
//
//	var inBattle bool
//	for i := 0; i < 20; i++ {
//		inBattle = MyOpenCV.ColorCmp(info.If_Map, 0.85)
//		if inBattle {
//			break
//		}
//		if i == 19 {
//			println("未进入跑图界面,未做异常处理")
//			os.Exit(0)
//		}
//		time.Sleep(1 * time.Second)
//	}
//
//	yolo := yolo.New("v5", 4, info.YoloParamPath, info.YoloBinPath, info.Yolo_labels)
//	if yolo == nil {
//		fmt.Println("模型加载失败")
//		return
//	}
//	defer yolo.Close()
//	fmt.Println("模型加载成功")
//
//	//var mu sync.Mutex
//	//go func() {
//	//	// 持续发送坐标
//	//	for {
//	//		x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//	//		mu.Lock()
//	//		sender.SendCoord(x, y)
//	//		mu.Unlock()
//	//		time.Sleep(200 * time.Millisecond)
//	//	}
//	//}()
//
//	ctx, cancel := context.WithCancel(context.Background())
//
//	go func(ctx context.Context) {
//		for {
//			select {
//			case <-ctx.Done(): // 监听到上下文被取消，主动退出
//				fmt.Println("收到Context信号，技能协程退出")
//				return
//			default:
//				if accelerate {
//					motion.Click(1000, 650, 2, 0) //加速
//					time.Sleep(1 * time.Second)
//				}
//				if Subdue {
//					motion.Click(967, 561, 2, 0) //压制
//					time.Sleep(1 * time.Second)
//				}
//				if stealth {
//					motion.Click(1000, 480, 2, 0) //隐身
//					time.Sleep(1 * time.Second)
//				}
//
//			}
//		}
//	}(ctx)
//
//	time.Sleep(1 * time.Second)
//	//=========================找第一个怪===================================
//	aStar.NavigateTo(bigMapPath, bin_mapPath,
//		aStar.Point{X: 244, Y: 223},
//		func() aStar.Point {
//			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//			return aStar.Point{X: x, Y: y}
//		},
//	)
//
//	//一直点击检测到的第一个目标位置，直到被消灭
//	aStar.YoloFind(yolo, bigMapPath)
//	//=========================找第一个怪===================================
//	time.Sleep(500 * time.Millisecond)
//	//=========================找第二个怪===================================
//	aStar.NavigateTo(bigMapPath, bin_mapPath,
//		aStar.Point{X: 296, Y: 264},
//		func() aStar.Point {
//			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//			return aStar.Point{X: x, Y: y}
//		},
//	)
//
//	//一直点击检测到的第一个目标位置，直到被消灭
//	aStar.YoloFind(yolo, bigMapPath)
//	//=========================找第二个怪===================================
//	time.Sleep(500 * time.Millisecond)
//	//=========================找第三个怪===================================
//	aStar.NavigateTo(bigMapPath, bin_mapPath,
//		aStar.Point{X: 203, Y: 419},
//		func() aStar.Point {
//			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//			return aStar.Point{X: x, Y: y}
//		},
//	)
//
//	//一直点击检测到的第一个目标位置，直到被消灭
//	aStar.YoloFind(yolo, bigMapPath)
//	//=========================找第三个怪===================================
//	time.Sleep(500 * time.Millisecond)
//	//=========================找第四个怪===================================
//	aStar.NavigateTo(bigMapPath, bin_mapPath,
//		aStar.Point{X: 391, Y: 228},
//		func() aStar.Point {
//			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//			return aStar.Point{X: x, Y: y}
//		},
//	)
//
//	//一直点击检测到的第一个目标位置，直到被消灭
//	aStar.YoloFind(yolo, bigMapPath)
//	//=========================找第四个怪===================================
//	time.Sleep(500 * time.Millisecond)
//	//=========================找第五个怪===================================
//	aStar.NavigateTo(bigMapPath, bin_mapPath,
//		aStar.Point{X: 416, Y: 132},
//		func() aStar.Point {
//			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
//			return aStar.Point{X: x, Y: y}
//		},
//	)
//
//	//一直点击检测到的第一个目标位置，直到被消灭
//	aStar.YoloFind(yolo, bigMapPath)
//
//	cancel()
//	time.Sleep(2 * time.Second)
//}
