package mainProcess

import (
	"app/MyOpenCV"
	"app/aStar"
	"app/info"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Dasongzi1366/AutoGo/motion"
	"github.com/Dasongzi1366/AutoGo/yolo"
)

func chapter1(bigMapPath1, bigMapPath2, bin_mapPath1, bin_mapPath2, yolo_parapath, yolo_binpath, yolo_labels, chapterImg_path string) {
	FindChapter(chapterImg_path)
	Chapter_Tp(info.Chapter1_1ColorCmp, info.Chapter1_1TpPoint) //第一个参数是地图找色,第二个参数是地图传送点
	chapter1_1(bigMapPath1, bin_mapPath1, yolo_parapath, yolo_binpath, yolo_labels)
	Chapter_Tp(info.Chapter1_2ColorCmp, info.Chapter1_2TpPoint)
	chapter1_2(bigMapPath2, bin_mapPath2, yolo_parapath, yolo_binpath, yolo_labels)

}

func chapter1_1(bigMapPath, bin_mapPath, yolo_parapath, yolo_binpath, yolo_labels string) {

	//obstacle, _ := aStar.LoadObstacleMap(bin_mapPath)
	var inBattle bool
	for i := 0; i < 20; i++ {
		inBattle = MyOpenCV.ColorCmp(info.If_Map, 0.85)
		if inBattle {
			break
		}
		if i == 19 {
			println("未进入跑图界面,未做异常处理")
			os.Exit(0)
		}
		time.Sleep(1 * time.Second)
	}

	yolo := yolo.New("v5", 4, yolo_parapath, yolo_binpath, yolo_labels)
	if yolo == nil {
		fmt.Println("模型加载失败")
		return
	}
	defer yolo.Close()
	fmt.Println("模型加载成功")

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
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // 监听到上下文被取消，主动退出
				fmt.Println("收到Context信号，技能协程退出")
				return
			default:
				if accelerate {
					motion.Click(1000, 650, 2, 0) //加速
					time.Sleep(1 * time.Second)
				}
				if Subdue {
					motion.Click(967, 561, 2, 0) //压制
					time.Sleep(1 * time.Second)
				}
				if stealth {
					motion.Click(1000, 480, 2, 0) //隐身
					time.Sleep(1 * time.Second)
				}

			}
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	//=========================找第一个怪===================================
	aStar.NavigateTo(bigMapPath, bin_mapPath,
		aStar.Point{X: 256, Y: 397},
		func() aStar.Point {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			return aStar.Point{X: x, Y: y}
		},
	)

	//一直点击检测到的第一个目标位置，直到被消灭
	aStar.YoloFind(yolo, bigMapPath)
	//=========================找第一个怪===================================
	time.Sleep(500 * time.Millisecond)
	//=========================找第二个怪===================================
	aStar.NavigateTo(bigMapPath, bin_mapPath,
		aStar.Point{X: 338, Y: 380},
		func() aStar.Point {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			return aStar.Point{X: x, Y: y}
		},
	)

	//一直点击检测到的第一个目标位置，直到被消灭
	aStar.YoloFind(yolo, bigMapPath)
	//=========================找第二个怪===================================
	time.Sleep(500 * time.Millisecond)
	//=========================找第三个怪===================================
	aStar.NavigateTo(bigMapPath, bin_mapPath,
		aStar.Point{X: 420, Y: 332},
		func() aStar.Point {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			return aStar.Point{X: x, Y: y}
		},
	)

	//一直点击检测到的第一个目标位置，直到被消灭
	aStar.YoloFind(yolo, bigMapPath)
	//=========================找第三个怪===================================
	time.Sleep(500 * time.Millisecond)
	//=========================找第四个怪===================================
	aStar.NavigateTo(bigMapPath, bin_mapPath,
		aStar.Point{X: 304, Y: 233},
		func() aStar.Point {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			return aStar.Point{X: x, Y: y}
		},
	)

	//一直点击检测到的第一个目标位置，直到被消灭
	aStar.YoloFind(yolo, bigMapPath)
	//=========================找第四个怪===================================
	time.Sleep(500 * time.Millisecond)
	//=========================找第五个怪===================================
	aStar.NavigateTo(bigMapPath, bin_mapPath,
		aStar.Point{X: 175, Y: 191},
		func() aStar.Point {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			return aStar.Point{X: x, Y: y}
		},
	)

	//一直点击检测到的第一个目标位置，直到被消灭
	aStar.YoloFind(yolo, bigMapPath)
	cancel()
	time.Sleep(2 * time.Second)
}

func chapter1_2(bigMapPath, bin_mapPath, yolo_parapath, yolo_binpath, yolo_labels string) {

	//obstacle, _ := aStar.LoadObstacleMap(bin_mapPath)

	var inBattle bool
	for i := 0; i < 20; i++ {
		inBattle = MyOpenCV.ColorCmp(info.If_Map, 0.85)
		if inBattle {
			break
		}
		if i == 19 {
			println("未进入跑图界面,未做异常处理")
			os.Exit(0)
		}
		time.Sleep(1 * time.Second)
	}

	yolo := yolo.New("v5", 4, yolo_parapath, yolo_binpath, yolo_labels)
	if yolo == nil {
		fmt.Println("模型加载失败")
		return
	}
	defer yolo.Close()
	fmt.Println("模型加载成功")

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

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // 监听到上下文被取消，主动退出
				fmt.Println("收到Context信号，技能协程退出")
				return
			default:
				if accelerate {
					motion.Click(1000, 650, 2, 0) //加速
					time.Sleep(1 * time.Second)
				}
				if Subdue {
					motion.Click(967, 561, 2, 0) //压制
					time.Sleep(1 * time.Second)
				}
				if stealth {
					motion.Click(1000, 480, 2, 0) //隐身
					time.Sleep(1 * time.Second)
				}

			}
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	//=========================找第一个怪===================================
	aStar.NavigateTo(bigMapPath, bin_mapPath,
		aStar.Point{X: 244, Y: 223},
		func() aStar.Point {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			return aStar.Point{X: x, Y: y}
		},
	)

	//一直点击检测到的第一个目标位置，直到被消灭
	aStar.YoloFind(yolo, bigMapPath)
	//=========================找第一个怪===================================
	time.Sleep(500 * time.Millisecond)
	//=========================找第二个怪===================================
	aStar.NavigateTo(bigMapPath, bin_mapPath,
		aStar.Point{X: 296, Y: 264},
		func() aStar.Point {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			return aStar.Point{X: x, Y: y}
		},
	)

	//一直点击检测到的第一个目标位置，直到被消灭
	aStar.YoloFind(yolo, bigMapPath)
	//=========================找第二个怪===================================
	time.Sleep(500 * time.Millisecond)
	//=========================找第三个怪===================================
	aStar.NavigateTo(bigMapPath, bin_mapPath,
		aStar.Point{X: 203, Y: 419},
		func() aStar.Point {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			return aStar.Point{X: x, Y: y}
		},
	)

	//一直点击检测到的第一个目标位置，直到被消灭
	aStar.YoloFind(yolo, bigMapPath)
	//=========================找第三个怪===================================
	time.Sleep(500 * time.Millisecond)
	//=========================找第四个怪===================================
	aStar.NavigateTo(bigMapPath, bin_mapPath,
		aStar.Point{X: 391, Y: 228},
		func() aStar.Point {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			return aStar.Point{X: x, Y: y}
		},
	)

	//一直点击检测到的第一个目标位置，直到被消灭
	aStar.YoloFind(yolo, bigMapPath)
	//=========================找第四个怪===================================
	time.Sleep(500 * time.Millisecond)
	//=========================找第五个怪===================================
	aStar.NavigateTo(bigMapPath, bin_mapPath,
		aStar.Point{X: 416, Y: 132},
		func() aStar.Point {
			x, y := MyOpenCV.MapMatch(bigMapPath, 143, 110, 285, 252, true, false, 0.6)
			return aStar.Point{X: x, Y: y}
		},
	)

	//一直点击检测到的第一个目标位置，直到被消灭
	aStar.YoloFind(yolo, bigMapPath)

	cancel()
	time.Sleep(2 * time.Second)
}
