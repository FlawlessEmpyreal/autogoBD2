package MyOpenCV

import (
	"app/info"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Dasongzi1366/AutoGo/images"
	"github.com/Dasongzi1366/AutoGo/opencv"
)

var chapterSelectInterfaceImg_path = "/mnt/shared/Pictures/img/misc/chapterSelectDetect.png"

// 把图片变成灰度图片
func Img2Gery(path string) {
	fmt.Printf("正在读取图片: %s\n", path)
	//打开图片
	grayMap := opencv.IMRead(path, 0)
	if grayMap.Empty() {
		fmt.Println("错误：读取图片失败，请检查路径是否正确")
		return
	}

	//给图片名称前面加gery
	dir := filepath.Dir(path)
	file := filepath.Base(path)
	saveFileName := filepath.Join(dir, "grey_"+file)

	//打印绝对路径
	absPath, _ := filepath.Abs(saveFileName)
	fmt.Printf("准备保存到绝对路径: %s\n", absPath)

	//保存灰度图片
	success := opencv.IMWrite(saveFileName, grayMap)
	if !success {
		fmt.Println("错误：保存图片失败！可能是路径权限问题，或者图片数据为空。")
		return
	}
	fmt.Println("✅ 转换成功！")
}

// 截取模板
func ScreenShoot(x1, y1, x2, y2 int) *[]byte {
	imgNRGBA := images.CaptureScreen(143, 110, 285, 252, 0)
	if imgNRGBA == nil {
		fmt.Println("截图失败，img 为 nil！")
		return nil
	}
	// 走到这里说明截图成功
	//fmt.Println("成功截到图片！")
	//fmt.Printf("截图尺寸: %v\n", imgNRGBA.Bounds())

	templateBytes, err := NrgbaToTemplate(imgNRGBA)
	if err != nil {
		fmt.Println("nrgba转*[]byte失败")
		return nil
	}
	return templateBytes

	////保存截图用来验证
	//h := img.Bounds().Dy()
	//w := img.Bounds().Dx()
	//buf := img.Pix
	//mat, err := opencv.NewMatFromBytes(h, w, opencv.MatTypeCV8UC4, buf)
	//if err != nil {
	//	panic(err) // 转换失败通常是因为数据长度和宽高不匹配
	//}
	//success := opencv.IMWrite("/mnt/shared/Pictures/img/map/test_map.jpg", mat)
	////success := opencv.IMWrite("/storage/emulated/0/Android/img/map/test_map.jpg", mat)
	//if !success {
	//	fmt.Println("错误：保存图片失败！可能是路径权限问题，或者图片数据为空。")
	//	return
}

// 地图匹配坐标
// 大地图bigMapPath,屏幕截图位置x1, y1, x2, y2
func MapMatch(bigMapPath string, x1, y1, x2, y2 int, isGray bool, isTransparent bool, sim float32) (int, int) {
	//读取大图
	bigMap, err := GetNRGBA(bigMapPath)
	if err != nil {
		println(err.Error())
		return -1, -1
	}

	//截小地图
	templatePtr := ScreenShoot(x1, y1, x2, y2)

	//验证打开了大地图
	//fmt.Printf("大地图尺寸: %v\n", bigMap.Bounds())

	//匹配当前坐标
	x, y := opencv.FindImageFromImage(bigMap, templatePtr, isGray, isTransparent, sim)

	//fmt.Printf("匹配模板坐标返回:%d %d\n", x, y)
	if x != -1 && y != -1 {
		x3 := x + (x2-x1)/2 - 4
		y3 := y + (y2-y1)/2
		return x3, y3
	} else {
		return x, y
	}
}

// 用这个函数找到大地图缩放最佳比例，然后缩放获得最佳大小地图
func FindBestMatchByScalingMap(originalBigMap_path string, template *image.NRGBA) {
	originalBigMap_Mat := GetMat(originalBigMap_path)
	if originalBigMap_Mat.Empty() {
		fmt.Println("读取原图失败！")
		return
	}
	defer originalBigMap_Mat.Close()

	template_Mat, err := nrgbaToMat(*template)
	if err != nil || template_Mat.Empty() {
		fmt.Println("解码模板失败！")
		return
	}
	defer template_Mat.Close()
	defer runtime.KeepAlive(template)

	fmt.Printf("BigMap_Channl:%d, BigMap_Type:%d\n", originalBigMap_Mat.Channels(), originalBigMap_Mat.Type())
	fmt.Printf("template_Channl:%d, template_Type:%d\n", template_Mat.Channels(), template_Mat.Type())

	//	//=====================================================================
	//	//=======================在这里填入缩放范围===============================
	scales := []float64{0.8, 0.81, 0.82, 0.83, 0.84, 0.85, 0.86, 0.87, 0.88, 0.89, 0.9}
	//	//=====================================================================
	//	//=====================================================================

	scaledBigMap := opencv.NewMat()
	defer scaledBigMap.Close()

	result := opencv.NewMat()
	defer result.Close()

	mask := opencv.NewMat()
	defer mask.Close()

	for _, scale := range scales {

		opencv.Resize(originalBigMap_Mat, &scaledBigMap, image.Point{}, scale, scale, opencv.InterpolationLinear)
		opencv.MatchTemplate(scaledBigMap, template_Mat, &result, 5, mask)

		_, max_val, _, _ := opencv.MinMaxLoc(result)

		fmt.Printf("缩放比例:%f ,最高匹配相似度为: %f\n", scale, max_val)

	}
}

// 缩放图片工具
func ReSizeImg(path string, scaledProportion float64) {
	originalImg := GetMat(path)
	scaledImg := opencv.NewMat()
	opencv.Resize(originalImg, &scaledImg, image.Point{}, scaledProportion, scaledProportion, opencv.InterpolationLinear)

	dir := filepath.Dir(path)
	file := filepath.Base(path)
	saveFileName := filepath.Join(dir, "scaled_"+file)

	//打印绝对路径
	absPath, _ := filepath.Abs(saveFileName)
	fmt.Printf("准备保存到绝对路径: %s\n", absPath)

	//保存图片
	success := opencv.IMWrite(saveFileName, scaledImg)
	if !success {
		fmt.Println("错误：保存图片失败！可能是路径权限问题，或者图片数据为空。")
		return
	}
}

func If_BattleInterface(sim float32) bool {
	for i := 0; i < len(info.BattleInterface); i++ {
		matched := images.CmpColor(info.BattleInterface[i].X, info.BattleInterface[i].Y, info.BattleInterface[i].Color, sim, 0)
		if matched == false {
			return false
		}
	}
	return true
}

func If_LoadingInterface(sim float32) bool {
	for i := 0; i < len(info.LoadingInterface); i++ {
		matched := images.CmpColor(info.LoadingInterface[i].X, info.LoadingInterface[i].Y, info.LoadingInterface[i].Color, sim, 0)
		if matched == false {
			return false
		}
	}
	return true
}

func If_chapterSelectInterface() *[]byte {
	return GetByte(chapterSelectInterfaceImg_path)
}

func If_TpInterface(sim float32) bool {
	count1 := 0
	matched1, matched2 := false, false
	for i := 0; i < len(info.TpInterface1); i++ {
		if images.CmpColor(info.TpInterface1[i].X, info.TpInterface1[i].Y, info.TpInterface1[i].Color, sim, 0) {
			count1++
		}
		//matched1 = images.CmpColor(info.TpInterface1[i].X, info.TpInterface1[i].Y, info.TpInterface1[i].Color, sim, 0)
		if count1 >= 2 {
			matched1 = true
			break
		}
	}
	count2 := 0
	for i := 0; i < len(info.TpInterface2); i++ {
		if images.CmpColor(info.TpInterface2[i].X, info.TpInterface2[i].Y, info.TpInterface2[i].Color, sim, 0) {
			count2++
		}
		//matched1 = images.CmpColor(info.TpInterface1[i].X, info.TpInterface1[i].Y, info.TpInterface1[i].Color, sim, 0)
		if count2 >= 2 {
			matched2 = true
			break
		}
	}
	//for i := 0; i < len(info.TpInterface2); i++ {
	//	matched2 = images.CmpColor(info.TpInterface2[i].X, info.TpInterface2[i].Y, info.TpInterface2[i].Color, sim, 0)
	//	if matched2 == false {
	//		break
	//	}
	//}
	if matched1 || matched2 {
		return true
	} else {
		return false
	}
}

func ColorCmp(ColorCmp []info.StructColorCmp, sim float32) bool { //传送地图多点找色
	if len(ColorCmp) < 3 {
		println("找色点太少,至少需要3个")
		os.Exit(0)
	}
	count := 0
	for i := 0; i < len(ColorCmp); i++ {
		matched := images.CmpColor(ColorCmp[i].X, ColorCmp[i].Y, ColorCmp[i].Color, sim, 0)
		if matched {
			count++
		}
	}
	if count >= len(ColorCmp)-1 { //有一个找色点被遮蔽也能匹配
		return true
	} else {
		return false
	}
}
