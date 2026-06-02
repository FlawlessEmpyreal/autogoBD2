package mainProcess

import (
	"app/MyOpenCV"
	"app/info"
	"image"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/Dasongzi1366/AutoGo/motion"
	"github.com/Dasongzi1366/AutoGo/opencv"
)

// 使用技能
var accelerate bool = true
var stealth bool = false
var Subdue bool = true

func MainProcess() {
	bigMapPath1 := "/mnt/shared/Pictures/img/map/scaled_grey_Extend_chapter1_1.jpg"
	bigMapPath2 := "/mnt/shared/Pictures/img/map/scaled_grey_Extend_chapter1_2.jpg"
	bin_mapPath1 := "/mnt/shared/Pictures/img/map/bin_map_chapter1_1.jpg"
	bin_mapPath2 := "/mnt/shared/Pictures/img/map/bin_map_chapter1_2.jpg"
	yolo_parapath := "/mnt/shared/Pictures/img/misc/chapter1_best.param"
	yolo_binpath := "/mnt/shared/Pictures/img/misc/chapter1_best.bin"
	yolo_labels := "chapter1_monster"
	chapterImg_path := "/mnt/shared/Pictures/img/misc/chapter1Selcet.png"
	chapter1(bigMapPath1, bigMapPath2, bin_mapPath1, bin_mapPath2, yolo_parapath, yolo_binpath, yolo_labels, chapterImg_path)

}

// 进入传送地图
func Chapter_Tp(chapterMap_Point []info.StructColorCmp, //找地图时用的找色
	Chapter_TpPoint image.Point, //传送点坐标
) {
	for i := 0; i < 20; i++ {
		if MyOpenCV.If_LoadingInterface(0.8) {
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
		if i == 19 {
			println("卡加载,退出")
			os.Exit(0)
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		if MyOpenCV.If_TpInterface(0.7) { //查看是否进入传送地图
			break
		}
		motion.Click(1089, 444, 0, 0) //点击传送阵
		time.Sleep(2 * time.Second)
		if MyOpenCV.If_TpInterface(0.7) { //查看是否进入传送地图
			break
		}
		motion.Click(870, 567, 0, 0) //如果在传送阵周围
		time.Sleep(2 * time.Second)
		if MyOpenCV.If_TpInterface(0.7) { //查看是否进入传送地图
			break
		}

		//如果在传送阵周围但是没有互动键
		theta := r.Float64() * 2 * math.Pi

		// 计算第一次点击的坐标
		x := 640 + int(float64(175)*math.Cos(theta))
		y := 360 + int(float64(175)*math.Sin(theta))

		motion.Click(x, y, 0, 0)

		if i == 9 {
			print("未进入传送地图")
			os.Exit(0)
		}
		time.Sleep(2 * time.Second)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		if MyOpenCV.ColorCmp(info.If_GenerateTP, 0.8) { //如果需要生成魔法阵
			motion.Click(735, 443, 0, 0)
		} else {
			break
		}
	}

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ { //来回找三次
		state_ := false
		for i := 0; i < 8; i++ { //将地图往左点
			if !MyOpenCV.ColorCmp(info.TpLeftButten, 0.8) { //没按钮也不用点了
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
			if !MyOpenCV.ColorCmp(info.TpRightButten, 0.8) { //没按钮也不用点了
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
			println("未找到传送地图,暂时未做异常处理")
			os.Exit(0)
		}
	}

	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		if MyOpenCV.ColorCmp(chapterMap_Point, 0.8) {
			motion.Click(Chapter_TpPoint.X, Chapter_TpPoint.Y, 0, 0)
		}

		time.Sleep(1 * time.Second)
		for i := 0; i < 8; i++ {
			if MyOpenCV.ColorCmp(info.If_GenerateTP, 0.7) {
				motion.Click(734, 442, 0, 0) //点击确定
			}
			if !MyOpenCV.ColorCmp(info.If_GenerateTP, 0.7) {
				break //确定没卡顿
			}
			time.Sleep(1000 * time.Millisecond)
			if i == 7 {
				println("确定生成传送阵时卡顿了")
				os.Exit(0)
			}
		}

		if !MyOpenCV.If_TpInterface(0.80) { //查看是否退出传送界面
			break
		} else {
			motion.Click(105, 40, 0, 0)
		}
		if i == 3 {
			println("未退出传送地图")
		}
	}
	time.Sleep(5 * time.Second) //成功传送后等五秒要不容易直接跳步骤
	println("进入地图")
}

func FindChapter(chapterImg_path string) { //必须在章节界面
	img := MyOpenCV.If_chapterSelectInterface()
	chapterImg := MyOpenCV.GetByte(chapterImg_path)
	if chapterImg == nil {
		println("章节选择图片读取错误,退出")
		os.Exit(0)
	}

	for i := 0; i < 3; i++ { //进入章节选择
		motion.Click(553, 652, 0, 0)
		time.Sleep(1 * time.Second)
		x, y := opencv.FindImage(37, 587, 1202, 681, img, false, false, 0.5, 0)
		println(x, y)
		if x != -1 && y != -1 {
			break
		} else {
			time.Sleep(1 * time.Second)
		}
		if i == 2 {
			println("没进入章节选择界面,暂未做异常处理,退出")
			os.Exit(0)
		}
	}

	x, y := opencv.FindImage(1, 585, 1276, 677, chapterImg, false, false, 0.8, 0) //先找本页面有没有目标章节
	if x != -1 && y != -1 {
		motion.Click(x+10, y+5, 0, 0)
		for i := 0; i < 6; i++ { //确认退出选章节界面
			x, y := opencv.FindImage(1, 585, 1276, 677, chapterImg, false, false, 0.8, 0)
			if x == -1 && y == -1 {
				time.Sleep(3 * time.Second)
				return
			}
		}
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
	}

	for i := 0; i < 6; i++ { //确认退出选章节界面
		x, y := opencv.FindImage(1, 585, 1276, 677, chapterImg, false, false, 0.8, 0)
		if x == -1 && y == -1 {
			break
		}
	}

	for i := 0; i < 20; i++ { //持续检测是否进入加载界面
		if MyOpenCV.If_LoadingInterface(0.8) {
			break
		}
		if i == 19 {
			println("未进入加载界面")
			os.Exit(0)
		}
		time.Sleep(500 * time.Millisecond)
	}
	for i := 0; i < 20; i++ {
		if !MyOpenCV.If_LoadingInterface(0.8) {
			break
		}
		if i == 19 {
			println("卡在加载界面")
			os.Exit(0)
		}
		time.Sleep(1 * time.Second)
	}

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

}
