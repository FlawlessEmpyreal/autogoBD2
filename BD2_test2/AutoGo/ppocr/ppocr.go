package ppocr

import "image"

type Ppocr struct{}

// Result 表示文字检测和识别的结果，包括位置、标签和置信度。
type Result struct {
	X       int     `json:"X"`
	Y       int     `json:"Y"`
	Width   int     `json:"宽"`
	Height  int     `json:"高"`
	Label   string  `json:"标签"`
	Score   float64 `json:"精度"`
	CenterX int     `json:"-"` //中心坐标X
	CenterY int     `json:"-"` //中心坐标Y
}

// New 创建一个新的 PPOCR 实例。
// 参数:
//
//	version: 模型版本，目前仅支持"v2"和"v5"
//
// 返回:
//
//	*Ppocr: 新创建的 PPOCR 实例，如果失败则返回 nil。
func New(version string) *Ppocr {
	return nil
}

// Ocr 从屏幕指定区域进行识别。
// 参数:
//
//	x1, y1: 检测区域的左上角坐标。
//	x2, y2: 检测区域的右下角坐标。如果 x2 或 y2 为 0，则表示使用设备的最大宽度或高度。
//	colorStr: 指定文字的颜色(支持偏色),格式如 "CCCCCC-101010"
//	displayId: 屏幕ID。
//
// 返回:
//
//	[]Result: 检测结果列表。如果检测失败或没有检测到任何结果，则返回 nil。
func (p *Ppocr) Ocr(x1, y1, x2, y2 int, colorStr string, displayId int) []Result {
	return nil
}

// OcrFromImage 从内存中的图像进行识别
// 参数:
//
//	img - NRGBA格式的图像对象
//	colorStr: 指定文字的颜色(支持偏色),格式如 "CCCCCC-101010"
//
// 返回:
//
//	[]Result: 检测结果列表。如果检测失败或没有检测到任何结果，则返回 nil。
func (p *Ppocr) OcrFromImage(img *image.NRGBA, colorStr string) []Result {
	return nil
}

// OcrFromBase64 从Base64编码的图像进行识别
// 参数:
//
//	b64 - 图像的Base64编码字符串
//	colorStr: 指定文字的颜色(支持偏色),格式如 "CCCCCC-101010"
//
// 返回:
//
//	[]Result: 检测结果列表。如果检测失败或没有检测到任何结果，则返回 nil。
func (p *Ppocr) OcrFromBase64(b64 string, colorStr string) []Result {
	return nil
}

// OcrFromPath 从文件路径进行识别
// 参数:
//
//	path - 图像文件的路径
//	colorStr: 指定文字的颜色(支持偏色),格式如 "CCCCCC-101010"
//
// 返回:
//
//	[]Result: 检测结果列表。如果检测失败或没有检测到任何结果，则返回 nil。
func (p *Ppocr) OcrFromPath(path string, colorStr string) []Result {
	return nil
}

// Close 释放引擎资源
func (p *Ppocr) Close() {}
