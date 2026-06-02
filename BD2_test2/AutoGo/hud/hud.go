package hud

// TextItem 表示HUD中的文本项。
type TextItem struct {
	TextColor string // 文字颜色，十六进制格式，例如 "#FFFFFF"
	Text      string // 文本内容
}

type HUD struct {
}

// New 创建一个新的HUD对象。
// 返回:
//
//	*HUD: 新创建的HUD对象。
func New() *HUD {
	return nil
}

// SetPosition 设置HUD的显示位置。
// 参数:
//
//	x1: 左上角X坐标。
//	y1: 左上角Y坐标。
//	x2: 右下角X坐标。
//	y2: 右下角Y坐标。
//
// 返回:
//
//	*HUD: 当前HUD对象。
func (h *HUD) SetPosition(x1, y1, x2, y2 int) *HUD {
	return nil
}

// SetBackgroundColor 设置HUD的背景颜色。
// 参数:
//
//	color: 背景颜色，十六进制格式，例如 "#000000"。
//
// 返回:
//
//	*HUD: 当前HUD对象。
func (h *HUD) SetBackgroundColor(color string) *HUD {
	return nil
}

// SetTextSize 设置HUD中文本的字体大小。
// 参数:
//
//	size: 字体大小。
//
// 返回:
//
//	*HUD: 当前HUD对象。
func (h *HUD) SetTextSize(size int) *HUD {
	return nil
}

// SetText 设置HUD中要显示的文本内容。
// 参数:
//
//	items: 文本项数组，每个项包含文本内容和颜色。
//
// 返回:
//
//	*HUD: 当前HUD对象。
func (h *HUD) SetText(items []TextItem) *HUD {
	return nil
}

// Show 显示HUD。
func (h *HUD) Show() {

}

// Hide 隐藏HUD。
func (h *HUD) Hide() {

}

// IsVisible 返回HUD是否可见。
// 返回:
//
//	bool: HUD是否可见。
func (h *HUD) IsVisible() bool {
	return false
}

// Destroy 销毁HUD对象，释放资源。
func (h *HUD) Destroy() {

}
