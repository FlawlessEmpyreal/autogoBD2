package opencv

import (
	"image"

	"github.com/Dasongzi1366/AutoGo/images"
)

// FindImage 在指定区域内查找匹配的图片模板，支持透明图像处理。
//
// 参数：
//   - x1, y1: 区域左上角的坐标。
//   - x2, y2: 区域右下角的坐标。当 x2 或 y2 为 0 时，表示使用图像的最大宽度或高度。
//   - template: 模板图片的字节数组指针，表示要在区域内查找的图片。
//   - isGray: 布尔值，指示是否将图像转换为灰度图进行匹配，提升匹配速度和鲁棒性。
//   - isTransparent: 是否按透明图处理。为 true 时，模板左上角第一个像素的 RGB 颜色会被当作透明色。
//   - sim: 相似度阈值，取值范围为 0.1 到 1.0，值越高表示匹配要求越精确。
//
// 返回值：
//   - (int, int): 返回找到的图片左上角坐标。如果未找到则返回 (-1, -1)。
//
// 透明图说明：
//   - isTransparent 为 false 时，按普通图片匹配，不生成遮罩。
//   - isTransparent 为 true 时，按透明图片匹配，忽略模板中与左上角像素同 RGB 的区域。
func FindImage(x1, y1, x2, y2 int, template *[]byte, isGray, isTransparent bool, sim float32, displayId int) (int, int) {
	return 0, 0
}

// FindImageFromImage 在给定图像中查找匹配的图片模板。
//
// 参数含义与 FindImage 相同，但 img 直接作为待匹配图像，不会进行屏幕截图。
func FindImageFromImage(img *image.NRGBA, template *[]byte, isGray, isTransparent bool, sim float32) (int, int) {
	return 0, 0
}

// FindImageAll 在指定区域内查找匹配的图片模板，返回所有符合条件的坐标。
//
// 参数：
//   - x1, y1: 区域左上角的坐标。
//   - x2, y2: 区域右下角的坐标。当 x2 或 y2 为 0 时，表示使用图像的最大宽度或高度。
//   - template: 模板图片的字节数组指针，表示要在区域内查找的图片。
//   - isGray: 布尔值，指示是否将图像转换为灰度图进行匹配，提升匹配速度和鲁棒性。
//   - isTransparent: 是否按透明图处理。为 true 时，模板左上角第一个像素的 RGB 颜色会被当作透明色。
//   - sim: 相似度阈值，取值范围为 0.1 到 1.0，值越高表示匹配要求越精确。
//
// 返回值：
//   - []images.Point: 返回所有符合条件的坐标。
func FindImageAll(x1, y1, x2, y2 int, template *[]byte, isGray, isTransparent bool, sim float32, displayId int) []images.Point {
	return nil
}
