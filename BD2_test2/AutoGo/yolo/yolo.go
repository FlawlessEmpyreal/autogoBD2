package yolo

import "image"

type Yolo struct {
}

// Result 表示对象检测的结果，包括位置、标签和置信度。
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

// New 创建一个新的 YOLO 实例，并加载模型和标签。
// 参数:
//
//	version: 模型版本，目前仅支持"v5"和"v8"
//	cpuThreadNum: 用于模型推理的 CPU 线程数。
//	paramPath: 模型参数文件路径。
//	binPath: 模型二进制文件路径。
//	labels: 标签文件路径。
//
// 返回:
//
//	*Yolo: 新创建的 YOLO 实例，如果加载失败则返回 nil。
func New(version string, cpuThreadNum int, paramPath, binPath, labels string) *Yolo {
	return nil
}

// SetImage 设置一个图片对象作为下次Detect方法的原始图像。
// 参数:
//
//	img: 图片对象。
func (y *Yolo) SetImage(img *image.NRGBA) {

}

// Detect 在指定的屏幕区域执行目标检测。
// 参数:
//
//	x1, y1: 检测区域的左上角坐标。
//	x2, y2: 检测区域的右下角坐标。如果 x2 或 y2 为 0，则表示使用设备的最大宽度或高度。
//	displayId: 屏幕ID。
//
// 返回:
//
//	[]Result: 检测结果列表。如果检测失败或没有检测到任何结果，则返回 nil。
func (y *Yolo) Detect(x1, y1, x2, y2, displayId int) []Result {
	return nil
}

// DetectFromImage 从内存中的图像进行识别
// 参数:
//
//	img - NRGBA格式的图像对象
//
// 返回:
//
//	[]Result: 检测结果列表。如果检测失败或没有检测到任何结果，则返回 nil。
func (y *Yolo) DetectFromImage(img *image.NRGBA) []Result {
	return nil
}

// DetectFromBase64 从Base64编码的图像进行识别
// 参数:
//
//	b64 - 图像的Base64编码字符串
//
// 返回:
//
//	[]Result: 检测结果列表。如果检测失败或没有检测到任何结果，则返回 nil。
func (y *Yolo) DetectFromBase64(b64 string) []Result {
	return nil
}

// DetectFromPath 从文件路径进行识别
// 参数:
//
//	path - 图像文件的路径
//
// 返回:
//
//	[]Result: 检测结果列表。如果检测失败或没有检测到任何结果，则返回 nil。
func (y *Yolo) DetectFromPath(path string) []Result {
	return nil
}

// Close 关闭 YOLO 模型实例，释放相关资源。
func (y *Yolo) Close() {

}
