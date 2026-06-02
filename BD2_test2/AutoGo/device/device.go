package device

var (
	CpuAbi        string // 设备的CPU架构，如"arm64-v8a", "x86", "x86_64"
	BuildId       string // 修订版本号，或者诸如"M4-rc20"的标识
	Broad         string // 设备的主板型号
	Brand         string // 与产品或硬件相关的厂商品牌，如"Xiaomi", "Huawei"等
	Device        string // 设备在工业设计中的名称
	Model         string // 设备型号
	Product       string // 整个产品的名称
	Bootloader    string // 设备Bootloader的版本
	Hardware      string // 设备的硬件名称
	Fingerprint   string // 构建(build)的唯一标识码
	Serial        string // 硬件序列号
	SdkInt        int    // 安卓系统API版本。例如安卓4.4的sdkInt为19
	Incremental   string // 设备构建的内部版本号
	Release       string // Android系统版本号。例如"5.0", "7.1.1"
	BaseOS        string // 设备的基础操作系统版本
	SecurityPatch string // 安全补丁程序级别
	Codename      string // 开发代号，例如发行版是"REL"
)

// Notification 通知信息
type Notification struct {
	PackageName string // 包名
	Id          string // ID
	Tag         string // 标签
	Title       string // 标题
	Text        string // 内容
}

// GetDisplayInfo 获取设备分辨率信息。
func GetDisplayInfo(displayId int) (width, height, dpi, rotation int) {
	return 0, 0, 0, 0
}

// GetImei 返回设备的IMEI。
// 返回:
//
//	string: 设备的IMEI。
func GetImei() string {
	return ""
}

// GetAndroidId 返回设备的Android ID。
// 返回:
//
//	string: 设备的Android ID。
func GetAndroidId() string {
	return ""
}

// GetWifiMac 获取设备WIFI-MAC。
// 返回:
//
//	string: 设备的WIFI MAC地址。
func GetWifiMac() string {
	return ""
}

// GetWlanMac 获取设备以太网MAC。
// 返回:
//
//	string: 设备的以太网MAC地址。
func GetWlanMac() string {
	return ""
}

// GetIp 获取设备局域网IP地址。
// 返回:
//
//	string: 设备的局域网IP地址。
func GetIp() string {
	return ""
}

// GetBrightness 返回当前的(手动)亮度。范围为0~255。
// 返回:
//
//	string: 当前的亮度值。
func GetBrightness() string {
	return ""
}

// GetBrightnessMode 返回当前亮度模式，0为手动亮度，1为自动亮度。
// 返回:
//
//	string: 当前的亮度模式。
func GetBrightnessMode() string {
	return ""
}

// GetMusicVolume 返回当前媒体音量。
// 返回:
//
//	int: 当前的媒体音量。
func GetMusicVolume() int {
	return 0
}

// GetNotificationVolume 返回当前通知音量。
// 返回:
//
//	int: 当前的通知音量。
func GetNotificationVolume() int {
	return 0
}

// GetAlarmVolume 返回当前闹钟音量。
// 返回:
//
//	int: 当前的闹钟音量。
func GetAlarmVolume() int {
	return 0
}

// GetMusicMaxVolume 返回媒体音量的最大值。
// 返回:
//
//	int: 媒体音量的最大值。
func GetMusicMaxVolume() int {
	return 0
}

// GetNotificationMaxVolume 返回通知音量的最大值。
// 返回:
//
//	int: 通知音量的最大值。
func GetNotificationMaxVolume() int {
	return 0
}

// GetAlarmMaxVolume 返回闹钟音量的最大值。
// 返回:
//
//	int: 闹钟音量的最大值。
func GetAlarmMaxVolume() int {
	return 0
}

// SetMusicVolume 设置当前媒体音量。
// 参数:
//
//	volume: 要设置的媒体音量。
func SetMusicVolume(volume int) {

}

// SetNotificationVolume 设置当前通知音量。
// 参数:
//
//	volume: 要设置的通知音量。
func SetNotificationVolume(volume int) {

}

// SetAlarmVolume 设置当前闹钟音量。
// 参数:
//
//	volume: 要设置的闹钟音量。
func SetAlarmVolume(volume int) {

}

// GetBattery 返回当前电量百分比。
// 返回:
//
//	int: 当前的电量百分比。
func GetBattery() int {
	return 0
}

// GetBatteryStatus 获取电池状态。 1：没有充电；2：正充电；3：没插充电器；4：不充电； 5：电池充满
// 返回:
//
//	int: 当前的电池状态。
func GetBatteryStatus() int {
	return 0
}

// SetBatteryStatus 模拟电池状态。 1：没有充电；2：正充电；5：电池充满
// 参数:
//
//	value: 要设置的电池状态。
func SetBatteryStatus(value int) {

}

// SetBatteryLevel 模拟电池电量百分百 0-100
// 参数:
//
//	value: 要设置的电池电量百分比。
func SetBatteryLevel(value int) {

}

// GetTotalMem 返回设备内存总量，单位(KB)。1MB = 1024KB。
// 返回:
//
//	int: 设备的内存总量。
func GetTotalMem() int {
	return 0
}

// GetAvailMem 返回设备当前可用的内存，单位字节(KB)。
// 返回:
//
//	int: 设备当前可用的内存。
func GetAvailMem() int {
	return 0
}

// IsScreenOn 返回设备屏幕是否是亮着的。如果屏幕亮着，返回true; 否则返回false。
// 返回:
//
//	bool: 屏幕是否亮着。
func IsScreenOn() bool {
	return false
}

// IsScreenUnlock 返回屏幕锁是否已经解开。已经解开返回true; 否则返回false。
// 返回:
//
//	bool: 屏幕锁是否已经解开。
func IsScreenUnlock() bool {
	return false
}

// SetDisplayPower 设置屏幕电源模式，不影响脚本运行。
//
// 参数:
//
//	on: 是否点亮。
func SetDisplayPower(on bool) {

}

// WakeUp 唤醒设备，包括唤醒设备CPU、屏幕等，可以用来点亮屏幕。
func WakeUp() {

}

// Reboot 重启设备。
func Reboot() {

}

// KeepScreenOn 保持屏幕常亮。
func KeepScreenOn() {

}

// Vibrate 使设备震动一段时间，单位毫秒，需要root权限。
// 参数:
//
//	ms: 要震动的时间（毫秒）。
func Vibrate(ms int) {

}

// CancelVibration 如果设备处于震动状态，则取消震动，需要root权限。
func CancelVibration() {

}

// GetNotification 获取当前所有通知消息。
func GetNotification() []Notification {
	return nil
}
