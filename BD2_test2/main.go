package main

import "app/imGui"

func main() {

	//MyOpenCV.Img2Gery("/mnt/shared/Pictures/img/map/Extend_chapter3_1.jpg")
	//MyOpenCV.Img2Gery("/mnt/shared/Pictures/img/map/Extend_chapter3_2.jpg")
	//MyOpenCV.ReSizeImg("/mnt/shared/Pictures/img/map/grey_Extend_chapter3_1.jpg", 0.87)
	//MyOpenCV.ReSizeImg("/mnt/shared/Pictures/img/map/grey_Extend_chapter3_2.jpg", 0.87)
	//aStar.BuildObstacleMap("/mnt/shared/Pictures/img/map/scaled_grey_Extend_chapter3_1.jpg", "bin_map_chapter3_1.jpg")
	//aStar.BuildObstacleMap("/mnt/shared/Pictures/img/map/scaled_grey_Extend_chapter3_2.jpg", "bin_map_chapter3_2.jpg")

	//yoloModel := yolo2.New(
	//	"v5",
	//	4,
	//	"/mnt/shared/Pictures/img/misc/best.param",
	//	"/mnt/shared/Pictures/img/misc/best.bin",
	//	"chapter1_monster,chapter2_monster,chapter3_monster,chapter4_monster,chapter5_monster,chapter6_monster,chapter7_monster,chapter8_monster,chapter9_monster,chapter10_monster,chapter11_monster,chapter12_monster,chapter13_monster,chapter14_monster,chapter15_monster,chapter16_monster,chapter17_monster",
	//)
	////yoloModel := yolo2.New("v5", 4, "/mnt/shared/Pictures/img/misc/chapter1_best.param", "/mnt/shared/Pictures/img/misc/chapter1_best.bin", "chapter1_monster")
	//for i := 1; i <= 17; i++ {
	//	path := fmt.Sprintf("/mnt/shared/Pictures/Screenshots/%d.png", i)
	//	fmt.Printf("%s:", path)
	//	img := images.ReadFromPath(path)
	//	results := yoloModel.DetectFromImage(img)
	//	fmt.Println(results)
	//}

	//mainProcess.MainProcess()

	imGui.ImGuiRun()

	//x, y := MyOpenCV.MapMatch(info_chapter.Ch3.MapConfig[1].BigMapPath,
	//	143, 110, 285, 252, true, false, 0.5)
	//fmt.Printf("x: %d, y: %d\n", x, y)

	//start := aStar.Point{X: x, Y: y}
	//end := aStar.Point{X: 304, Y: 273}
	//astarMap, _ := aStar.LoadObstacleMap("/mnt/shared/Pictures/img/map/bin_map_chapter3_1.jpg")
	////path := aStar.FindPath(astarMap, start, end)
	//aStar.FindPath(astarMap, start, end)
	// 打印路径点，找到角色当前位置附近的那几个点
	//cur := aStar.getCurrentPos()
	//for i, p := range path {
	//	dx := p.X - start.X
	//	dy := p.Y - start.Y
	//	dist := math.Sqrt(float64(dx*dx + dy*dy))
	//	if dist < 50 { // 只打印附近的点
	//		fmt.Printf("路径点%d: (%d,%d) 距当前位置=%.1f\n", i, p.X, p.Y, dist)
	//	}
	//}

	//aStar.Handle()

	//// 获取当前工作目录
	//dir, err := os.Getwd()
	//if err != nil {
	//	fmt.Println("获取工作目录失败:", err)
	//	return
	//}
	//fmt.Println("当前工作目录:", dir)

}
