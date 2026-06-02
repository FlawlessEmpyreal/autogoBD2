package apkctl

// LifecycleEvent 表示脚本运行过程中的控制事件。
//
// 可用事件:
//   - EventPause: 暂停脚本。
//   - EventResume: 恢复脚本。
//   - EventStop: 停止脚本。
type LifecycleEvent string

const (
	// EventPause 表示脚本暂停事件。
	EventPause LifecycleEvent = "pause"

	// EventResume 表示脚本恢复事件。
	EventResume LifecycleEvent = "resume"

	// EventStop 表示脚本停止事件。
	EventStop LifecycleEvent = "stop"
)

// Eval 执行指定的 JavaScript 脚本并返回结果。
//
// 参数:
//   - contextID: 执行上下文 ID。相同的 contextID 会复用同一个 JS 作用域，
//     因此可以在多次 Eval 之间保留变量、控件引用和页面状态。
//   - script: 需要执行的 JavaScript 代码字符串。
//
// 默认注入到 JS 作用域中的对象:
//   - appContext: Application Context。
//   - go: 一个桥接对象，可调用 `go.send("消息")` 把字符串消息发回 Go。
//   - scriptUi: 原生脚本页面控制器，用来打开/关闭一个专门给 JS 使用的空 Activity。
//   - uiActivity / uiContext / uiRoot: 当前脚本页面已经就绪时，可直接拿来创建原生控件。
//
// 当前 APK 已内置一些常用原生 View 库，JS 中可直接 importClass 使用，例如:
//   - Packages.com.google.android.material.button.MaterialButton
//   - Packages.com.google.android.material.textfield.TextInputLayout
//   - Packages.com.google.android.material.textfield.TextInputEditText
//   - Packages.com.google.android.material.card.MaterialCardView
//   - Packages.androidx.recyclerview.widget.RecyclerView
//   - Packages.androidx.swiperefreshlayout.widget.SwipeRefreshLayout
//   - Packages.androidx.viewpager2.widget.ViewPager2
//
// scriptUi 常用方法:
//   - scriptUi.open()                      打开原生脚本页面
//   - scriptUi.getActivity() / getContext()
//   - scriptUi.getRoot()
//   - scriptUi.runOnUiThread(runnable)     在主线程创建/更新原生控件
//   - scriptUi.clear()                     清空页面上的所有原生控件
//   - scriptUi.close()                     关闭UI页面
//
// 创建原生 UI 的典型流程:
//  1. 先调用 `scriptUi.open()`
//  2. 再获取 `scriptUi.getActivity()` 和 `scriptUi.getRoot()`
//  3. 在 `scriptUi.runOnUiThread(...)` 中创建 Button / EditText / Spinner 等原生控件
//
// JS 侧示例:
//
//	if (!scriptUi.open()) {
//	    throw new Error("ScriptUiActivity not ready");
//	}
//	importClass(Packages.com.google.android.material.button.MaterialButton);
//	var activity = scriptUi.getActivity();
//	var root = scriptUi.getRoot();
//	scriptUi.runOnUiThread(new java.lang.Runnable({
//	    run: function () {
//	        root.removeAllViews();
//	        var button = new MaterialButton(activity);
//	        button.setText("点我");
//	        root.addView(button);
//	    }
//	}));
//
// 返回值:
//   - string: 脚本执行后的返回值。若 JS 最终没有显式返回内容，通常会得到空字符串。
func Eval(contextID, script string) string {
	return ""
}

// SetCallback 设置一个接收 JS 主动消息的回调。
//
// 当 APK 内执行的 JavaScript 调用 `go.send("你的消息")` 时，
// Java 层会把这条消息主动推回 Go，
// 然后这里注册的 callback 就会收到：
//   - contextID: 发送这条消息的 JS 执行上下文 ID
//   - message: JS 里传给 `go.send(...)` 的原始字符串
//
// JS 侧示例:
//
//	go.send("button_clicked")
//	go.send("seekbar:" + value)
//
// Go 侧示例:
//
//	apkctl.SetCallback(func(contextID, message string) {
//	    fmt.Println("from js:", contextID, message)
//	})
//
// 说明:
//   - 只有通过 APK 的 JS 引擎执行的脚本，调用 `go.send(...)` 才会触发这里的回调。
//   - 如果没有设置 callback，JS 发来的消息会被接收但不会进一步处理。
//   - message 是普通字符串；如果要传复杂结构，建议 JS 自己先 JSON.stringify 后再发送。
func SetCallback(callback func(contextID, message string)) {

}

// RegEvent 注册脚本控制事件回调。
//
// 参数:
//   - event: 要监听的事件，可使用 EventPause、EventResume、EventStop。
//   - callback: 事件触发时执行的 Go 回调；传 nil 表示取消该事件的回调注册。
//
// Go 侧示例:
//
//	apkctl.RegEvent(apkctl.EventPause, func() {
//	    apkctl.Toast("准备暂停")
//	})
//
// 注意:
//   - EventPause、EventStop 的 callback 执行完成后，脚本才会继续完成对应操作。
//   - EventResume 是恢复后的通知事件，不会等待 callback 返回。
func RegEvent(event LifecycleEvent, callback func()) {

}
