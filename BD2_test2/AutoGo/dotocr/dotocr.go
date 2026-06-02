package dotocr

import (
	"image"
)

// SetDict 设置字库
//
// 参数说明：
//
//	name: 字库名称，为空字符串时使用 "default"
//	dict: 字库内容字符串，按行分割，每行一条模板记录
func SetDict(name, dict string) {

}

// Ocr 从屏幕指定区域进行 OCR 识别
//
// 参数说明：
//
//	x1, y1: 识别区域的左上角坐标
//	x2, y2: 识别区域的右下角坐标
//	threshold: 阈值字符串，例如 "ffffff-101010"
//	sim: 匹配相似度阈值（0.0 ~ 1.0），例如 0.8
//	asJSON: 是否以 json 格式返回
//	dictName: 使用的字库名称，为空字符串时使用 "default"
//	displayId: 显示器 ID
//
// 返回值：
//
//	识别结果字符串（纯文本或 JSON 格式）
func Ocr(x1, y1, x2, y2 int, threshold string, sim float32, asJSON bool, dictName string, displayId int) string {
	return ""
}

// OcrFromImage 从图像对象进行 OCR 识别
//
// 参数说明：
//
//	img: NRGBA 格式的图像对象
//	threshold: 阈值字符串
//	sim: 匹配相似度阈值
//	asJSON: 是否以 json 格式返回
//	dictName: 使用的字库名称
//
// 返回值：
//
//	识别结果字符串
func OcrFromImage(img *image.NRGBA, threshold string, sim float32, asJSON bool, dictName string) string {
	return ""
}

// OcrFromBase64 从 Base64 编码的图像字符串进行 OCR 识别
//
// 参数说明：
//
//	b64: Base64 编码的图像数据
//	threshold: 阈值字符串
//	sim: 匹配相似度阈值
//	asJSON: 是否以 json 格式返回
//	dictName: 使用的字库名称
//
// 返回值：
//
//	识别结果字符串
func OcrFromBase64(b64, threshold string, sim float32, asJSON bool, dictName string) string {
	return ""
}

// OcrFromPath 从图像文件路径进行 OCR 识别
//
// 参数说明：
//
//	path: 图像文件路径
//	threshold: 阈值字符串
//	sim: 匹配相似度阈值
//	asJSON: 是否以 json 格式返回
//	dictName: 使用的字库名称
//
// 返回值：
//
//	识别结果字符串
func OcrFromPath(path, threshold string, sim float32, asJSON bool, dictName string) string {
	return ""
}

// FindStr 在屏幕指定区域中查找指定字符串的位置
//
// 参数说明：
//
//	x1, y1: 查找区域的左上角坐标
//	x2, y2: 查找区域的右下角坐标
//	text: 要查找的字符串
//	threshold: 阈值字符串，例如 "ffffff-101010"
//	sim: 匹配相似度阈值
//	dictName: 使用的字库名称
//	displayId: 显示器 ID
//
// 返回值：
//
//	找到时返回字符串第一个字符的坐标 (x, y)，未找到返回 (-1, -1)
func FindStr(x1, y1, x2, y2 int, text, threshold string, sim float32, dictName string, displayId int) (int, int) {
	return -1, -1
}

// FindStrFromImage 在图像对象中查找指定字符串的位置
//
// 参数说明：
//
//	img: NRGBA 格式的图像对象
//	text: 要查找的字符串
//	threshold: 阈值字符串
//	sim: 匹配相似度阈值
//	dictName: 使用的字库名称
//
// 返回值：
//
//	找到时返回字符串第一个字符的坐标 (x, y)，未找到返回 (-1, -1)
func FindStrFromImage(img *image.NRGBA, text, threshold string, sim float32, dictName string) (int, int) {
	return -1, -1
}

// FindStrFromBase64 在 Base64 编码的图像中查找指定字符串的位置
//
// 参数说明：
//
//	b64: Base64 编码的图像数据
//	text: 要查找的字符串
//	threshold: 阈值字符串
//	sim: 匹配相似度阈值
//	dictName: 使用的字库名称
//
// 返回值：
//
//	找到时返回字符串第一个字符的坐标 (x, y)，未找到返回 (-1, -1)
func FindStrFromBase64(b64, text, threshold string, sim float32, dictName string) (int, int) {
	return -1, -1
}

// FindStrFromPath 在图像文件中查找指定字符串的位置
//
// 参数说明：
//
//	path: 图像文件路径
//	text: 要查找的字符串
//	threshold: 阈值字符串
//	sim: 匹配相似度阈值
//	dictName: 使用的字库名称
//
// 返回值：
//
//	找到时返回字符串第一个字符的坐标 (x, y)，未找到返回 (-1, -1)
func FindStrFromPath(path, text, threshold string, sim float32, dictName string) (int, int) {
	return -1, -1
}
