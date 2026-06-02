package vdisplay

type Vdisplay struct {
}

// Create 创建一个虚拟显示设备，该方法要求安卓10及以上版本。
// 参数:
//
//	width: 显示设备宽度。
//	height: 显示设备高度。
//	dpi: 显示设备的DPI。
//
// 返回:
//
//	创建成功后的虚拟显示对象指针。如果创建失败，返回nil。
func Create(width, height, dpi int) *Vdisplay {
	return nil
}

// GetDisplayId 获取虚拟显示设备的DisplayId
func (v *Vdisplay) GetDisplayId() int {
	return -1
}

// LaunchApp 启动指定包名的应用到虚拟显示设备内
func (v *Vdisplay) LaunchApp(packageName string) bool {
	return false
}

// SetTitle 设置预览窗口标题
func (v *Vdisplay) SetTitle(title string) {

}

// SetTouchCallback 设置触控回调
// 参数:
//
// action参数说明，0=按下 1=抬起 2=移动
func (v *Vdisplay) SetTouchCallback(callback func(x, y, action, displayId int)) {

}

// ShowPreviewWindow 显示预览窗口
func (v *Vdisplay) ShowPreviewWindow(rotated bool) {

}

// HidePreviewWindow 隐藏预览窗口
func (v *Vdisplay) HidePreviewWindow() {

}

// SetPreviewWindowSize 设置预览窗口大小
func (v *Vdisplay) SetPreviewWindowSize(width, height int) {

}

// SetPreviewWindowPos 设置预览窗口位置
func (v *Vdisplay) SetPreviewWindowPos(x, y int) {

}

// Destroy 销毁指定的虚拟显示设备。
func (v *Vdisplay) Destroy() {

}
