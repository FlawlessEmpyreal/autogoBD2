package motion

// TouchDown 模拟触摸屏按下操作。
// 参数:
//
//	x: 触摸点的 X 坐标。
//	y: 触摸点的 Y 坐标。
//	fingerID: 触摸点的指针 ID（0-9）。
//	displayId: 屏幕ID。
func TouchDown(x, y, fingerID, displayId int) {

}

// TouchMove 模拟触摸屏移动操作。
// 参数:
//
//	x: 移动到的 X 坐标。
//	y: 移动到的 Y 坐标。
//	fingerID: 触摸点的指针 ID（0-9）。
//	displayId: 屏幕ID。
func TouchMove(x, y, fingerID, displayId int) {

}

// TouchUp 模拟触摸屏抬起操作。
// 参数:
//
//	x: 抬起点的 X 坐标。
//	y: 抬起点的 Y 坐标。
//	fingerID: 触摸点的指针 ID（0-9）。
//	displayId: 屏幕ID。
func TouchUp(x, y, fingerID, displayId int) {

}

// Click 模拟单击操作。
// 参数:
//
//	x: 单击点的 X 坐标。
//	y: 单击点的 Y 坐标。
//	fingerID: 触摸点的指针 ID（0-9）。
//	displayId: 屏幕ID。
func Click(x, y, fingerID, displayId int) {

}

// LongClick 模拟长按操作。
// 参数:
//
//	x: 长按点的 X 坐标。
//	y: 长按点的 Y 坐标。
//	duration: 长按持续时间（毫秒）。
//	fingerID: 触摸点的指针 ID（0-9）。
//	displayId: 屏幕ID。
func LongClick(x, y, duration, fingerID, displayId int) {

}

// Swipe 模拟滑动操作。
// 参数:
//
//	x1: 起始点的 X 坐标。
//	y1: 起始点的 Y 坐标。
//	x2: 结束点的 X 坐标。
//	y2: 结束点的 Y 坐标。
//	duration: 滑动持续时间（毫秒）。
//	fingerID: 触摸点的指针 ID（0-9）。
//	displayId: 屏幕ID。
func Swipe(x1, y1, x2, y2, duration, fingerID, displayId int) {

}

// Swipe2 使用贝塞尔曲线方式进行滑动
// 参数:
//
//	x1: 起始点的 X 坐标。
//	y1: 起始点的 Y 坐标。
//	x2: 结束点的 X 坐标。
//	y2: 结束点的 Y 坐标。
//	duration: 滑动持续时间（毫秒）。
//	fingerID: 触摸点的指针 ID（0-9）。
//	displayId: 屏幕ID。
func Swipe2(x1, y1, x2, y2, duration, fingerID, displayId int) {

}

// Home 模拟按下 Home 键。
// 参数:
//
//	displayId: 屏幕ID。
func Home(displayId int) {

}

// Back 模拟按下返回键。
// 参数:
//
//	displayId: 屏幕ID。
func Back(displayId int) {

}

// Recents 显示最近任务。
// 参数:
//
//	displayId: 屏幕ID。
func Recents(displayId int) {

}

// PowerDialog 弹出电源键菜单。
func PowerDialog() {

}

// Notifications 拉出通知栏。
func Notifications() {

}

// QuickSettings 显示快速设置(下拉通知栏到底)。
func QuickSettings() {

}

// VolumeUp 按下音量上键。
// 参数:
//
//	displayId: 屏幕ID。
func VolumeUp(displayId int) {

}

// VolumeDown 按下音量下键。
// 参数:
//
//	displayId: 屏幕ID。
func VolumeDown(displayId int) {

}

// Camera 模拟按下照相键。
func Camera() {

}

// KeyAction 模拟按键单击 code值参考KEYCODE_开头常量。
// 参数:
//
//	displayId: 屏幕ID。
func KeyAction(code, displayId int) {

}

// KeyActionDown 模拟按键按下 code值参考KEYCODE_开头常量。
// 参数:
//
//	displayId: 屏幕ID。
func KeyActionDown(code, displayId int) {

}

// KeyActionUp 模拟按键弹起 code值参考KEYCODE_开头常量。
// 参数:
//
//	displayId: 屏幕ID。
func KeyActionUp(code, displayId int) {

}
