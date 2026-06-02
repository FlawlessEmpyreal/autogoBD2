package plugin

// ClassLoader APK类加载器
type ClassLoader struct {
}

// Instance Java对象实例
type Instance struct {
}

// Context 安卓Context参数标记,使用时会自动获取Application Context
type Context struct{}

// AssetManager 安卓AssetManager参数标记
type AssetManager struct{}

// NewContext 创建Context参数
func NewContext() Context {
	return Context{}
}

// NewAssetManager 创建AssetManager参数
func NewAssetManager() AssetManager {
	return AssetManager{}
}

// LoadApk 加载外部APK
func LoadApk(apkPath string) *ClassLoader {
	return nil
}

// NewInstance 创建类实例
func (cl *ClassLoader) NewInstance(className string, args ...interface{}) *Instance {
	return nil
}

// CallString 调用返回String的方法
func (inst *Instance) CallString(methodName string, args ...interface{}) (string, error) {
	return "", nil
}

// CallInt 调用返回int的方法
func (inst *Instance) CallInt(methodName string, args ...interface{}) (int, error) {
	return 0, nil
}

// CallLong 调用返回long的方法
func (inst *Instance) CallLong(methodName string, args ...interface{}) (int64, error) {
	return 0, nil
}

// CallFloat 调用返回float的方法
func (inst *Instance) CallFloat(methodName string, args ...interface{}) (float32, error) {
	return 0, nil
}

// CallDouble 调用返回double的方法
func (inst *Instance) CallDouble(methodName string, args ...interface{}) (float64, error) {
	return 0, nil
}

// CallBool 调用返回boolean的方法
func (inst *Instance) CallBool(methodName string, args ...interface{}) (bool, error) {
	return false, nil
}

// CallVoid 调用无返回值的方法
func (inst *Instance) CallVoid(methodName string, args ...interface{}) error {
	return nil
}

// Release 释放实例
func (inst *Instance) Release() {

}

// Release 释放类加载器
func (cl *ClassLoader) Release() {

}
