package MyOpenCV

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/Dasongzi1366/AutoGo/opencv"
)

// 以*[]byte格式打开一个图片
func GetByte(path string) *[]byte {
	imageBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return nil
	}
	var imagePtr *[]byte = &imageBytes
	return imagePtr
}

// 以NRGBA格式打开一个图片
func GetNRGBA(filePath string) (*image.NRGBA, error) {
	// 1. 打开本地图片文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// 2. 解码文件，将其转为通用的 image.Image 接口
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// 3. 格式转换：将通用的 image.Image 转换为 *image.NRGBA
	// 先获取原图的边界
	bounds := img.Bounds()
	// 创建一个新的 NRGBA 画布
	nrgba := image.NewNRGBA(bounds)
	// 使用 draw.Draw 将原图的内容完美地绘制到新的 NRGBA 画布上
	draw.Draw(nrgba, bounds, img, bounds.Min, draw.Src)
	defer file.Close()
	return nrgba, nil
}

// nrgba转Template *[]byte
func NrgbaToTemplate(img *image.NRGBA) (*[]byte, error) {
	// 1. 创建一个字节缓冲区，用来接收编码后的图片数据
	var buf bytes.Buffer

	// 2. 将 NRGBA 图片编码为 PNG 格式的字节流写入缓冲区
	// （如果你的模板要求是 JPG，就换成 jpeg.Encode(&buf, img, nil)）
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}

	// 3. 将缓冲区的数据转为字节切片，并返回其指针
	templateBytes := buf.Bytes()
	return &templateBytes, nil
}

// 以Mat格式打开图片
func GetMat(path string) opencv.Mat {
	img := opencv.IMRead(path, opencv.IMReadColor)
	if img.Empty() {
		fmt.Println("无法读取图片，请检查路径是否正确！")
		return opencv.Mat{}
	}
	fmt.Printf("成功以 Mat 格式打开图片！尺寸: %v\n", img.Size())
	return img
}

// NRGBA转MAT
func nrgbaToMat(nrgbaImg image.NRGBA) (opencv.Mat, error) {
	// 直接使用 gocv.ImageToMat 进行转换
	mat, err := opencv.ImageToMatRGB(&nrgbaImg)
	if err != nil {
		return opencv.Mat{}, fmt.Errorf("转换失败: %w", err)
	}
	return mat, nil
}

// *[]byte转Mat
func byte2Mat(templateBytes *[]byte) opencv.Mat {
	// 1. 提前检查字节流是否为空
	if len(*templateBytes) == 0 {
		fmt.Println("传入的字节流为空！")
		return opencv.Mat{}
	}

	// 2. 直接使用 IMDecode 将字节流解码为图片 Mat
	// gocv.IMDecode 内部会自动完成从 []byte 到图片矩阵的转换
	img, _ := opencv.IMDecode(*templateBytes, opencv.IMReadColor)

	// 3. 检查解码是否成功
	if img.Empty() {
		fmt.Println("字节流解码为 Mat 失败，请检查是否为有效的图片数据！")
		return opencv.Mat{}
	}

	// 4. 返回真正的图片 Mat（注意：这里不要 Close，交给调用者去释放）
	return img
}
