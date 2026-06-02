package app

// AppInfo 应用信息结构体
type AppInfo struct {
	PackageName string // 包名
	AppName     string // 应用显示名称
	VersionName string // 版本名称
	VersionCode string // 版本代码
	IsSystemApp bool   // 是否系统应用
	Enabled     bool   // 是否启用
}

// IntentOptions 用于构建Intent的选项结构体。
type IntentOptions struct {
	Action      string            // Intent的动作
	Type        string            // Intent的数据类型
	Data        string            // Intent的数据
	Category    []string          // Intent的类别
	PackageName string            // 应用包名
	Extras      map[string]string // Intent的额外参数
	Flags       []string          // Intent的标志
}

// CurrentPackage 获取当前页面应用包名。
// 返回:
//
//	string: 应用包名。
func CurrentPackage() string {
	return ""
}

// CurrentActivity 获取当前页面应用类名。
// 返回:
//
//	string: 应用类名。
func CurrentActivity() string {
	return ""
}

// Launch 通过应用包名启动应用。
// 参数:
//
//	packageName: 应用包名。
//	displayId: 屏幕ID。
//
// 返回:
//
//	bool: 是否成功启动应用。
func Launch(packageName string, displayId int) bool {
	return false
}

// SelfPackage 获取当前脚本/宿主自身的包名。
func SelfPackage() string {
	return ""
}

// GetList 获取手机中所有应用列表。
// 参数:
//
//	includeSystemApps: 是否包含系统应用。
//
// 返回:
//
//	[]AppInfo: 应用信息数组。
func GetList(includeSystemApps bool) []AppInfo {
	return nil
}

// GetName 获取应用名称。
// 参数:
//
//	packageName: 应用包名。
//
// 返回:
//
//	string: 应用名称。
func GetName(packageName string) string {
	return ""
}

// GetIcon 获取应用图标。
// 参数:
//
//	packageName: 应用包名。
//
// 返回:
//
//	[]byte: 图标字节数据。
func GetIcon(packageName string) []byte {
	return nil
}

// GetVersion 获取应用版本号。
// 参数:
//
//	packageName: 应用包名。
//
// 返回:
//
//	string: 版本号。
func GetVersion(packageName string) string {
	return ""
}

// OpenSetting 打开应用的详情页(设置页)。
// 参数:
//
//	packageName: 应用包名。
//
// 返回:
//
//	bool: 是否成功打开应用设置页。
func OpenSetting(packageName string) bool {
	return false
}

// ViewFile 用其他应用查看文件。文件不存在的情况由查看文件的应用处理。
// 参数:
//
//	path: 文件路径。
func ViewFile(path string) {

}

// EditFile 用其他应用编辑文件。文件不存在的情况由编辑文件的应用处理。
// 参数:
//
//	path: 文件路径。
func EditFile(path string) {

}

// Uninstall 卸载应用。
// 参数:
//
//	packageName: 应用包名。
func Uninstall(packageName string) {

}

// Install 安装应用。
// 参数:
//
//	path: APK文件路径。
func Install(path string) {

}

// IsInstalled 判断是否已经安装某个应用。
// 参数:
//
//	packageName: 应用包名。
//
// 返回:
//
//	bool: 应用是否已安装。
func IsInstalled(packageName string) bool {
	return false
}

// Clear 清除应用数据。
// 参数:
//
//	packageName: 应用包名。
func Clear(packageName string) {

}

// ForceStop 强制停止应用。
// 参数:
//
//	packageName: 应用包名。
func ForceStop(packageName string) {

}

// Disable 禁用应用。
// 参数:
//
//	packageName: 应用包名。
func Disable(packageName string) {

}

// Enable 启用应用。
// 参数:
//
//	packageName: 应用包名。
func Enable(packageName string) {

}

// EnableAccessibility 启用无障碍服务
// 参数:
//
//	packageName: 应用包名。
func EnableAccessibility(packageName string) {

}

// DisableAccessibility 禁用无障碍服务
// 参数:
//
//	packageName: 应用包名。
func DisableAccessibility(packageName string) {

}

// IgnoreBattOpt 忽略电池优化
// 参数:
//
//	packageName: 应用包名。
func IgnoreBattOpt(packageName string) {

}

// GetBrowserPackage 获取系统默认浏览器包名
// 返回:
//
//	string: 应用包名。
func GetBrowserPackage() string {
	return ""
}

// OpenUrl 用浏览器打开网站url。
// 参数:
//
//	url: 网站URL。
func OpenUrl(url string) {

}

// StartActivity 根据选项构造一个Intent，并启动该Activity。
// 参数:
//
//	options: Intent选项。
func StartActivity(options IntentOptions) {

}

// SendBroadcast 根据选项构造一个Intent，并发送该广播。
// 参数:
//
//	options: Intent选项。
func SendBroadcast(options IntentOptions) {

}

// StartService 根据选项构造一个Intent，并启动该服务。
// 参数:
//
//	options: Intent选项。
func StartService(options IntentOptions) {

}
