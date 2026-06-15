package imGui

import (
	"fmt"
	"time"

	"github.com/Dasongzi1366/AutoGo/imgui"
)

func ImGuiRun() {
	// 初始化 ImGui
	imgui.Init()

	// 1. 定义状态变量
	showWindow := true       // 控制窗口是否显示
	isRunning := false       // 控制脚本是否正在运行
	myScriptEnabled := false // 复选框绑定的布尔变量

	// 2. 主循环
	imgui.Run(func() {
		// 如果脚本正在运行，可以选择不显示窗口，或者显示一个“运行中”的提示
		if !isRunning && showWindow {
			// 设置窗口大小（宽 300，高 200），只在首次创建时生效
			imgui.SetNextWindowSizeV(imgui.Vec2{X: 300, Y: 200}, imgui.CondOnce)

			// 创建窗口。第二个参数传入 &showWindow，点击窗口右上角的 X 也会关闭
			imgui.BeginV("我的脚本控制台", &showWindow, 0)

			// 3. 添加复选框控件
			// 勾选时 myScriptEnabled 自动变为 true，取消勾选变为 false
			imgui.Checkbox("chapter1", &myScriptEnabled)

			imgui.Spacing()
			imgui.Separator()
			imgui.Spacing()

			// 4. 添加运行按钮
			if imgui.Button("点击运行") {
				isRunning = true
				showWindow = false // 点击后隐藏窗口

				// 启动你的脚本协程
				go func() {
					fmt.Println("脚本开始运行，启用状态:", myScriptEnabled)
					// 模拟脚本运行耗时
					time.Sleep(5 * time.Second)
					fmt.Println("脚本运行结束")

					// 脚本运行完毕后，可以重新显示窗口
					isRunning = false
					showWindow = true
				}()
			}

			imgui.End()
		} else if isRunning {
			// 脚本运行中，显示一个小的提示窗口（可选）
			imgui.SetNextWindowSizeV(imgui.Vec2{X: 200, Y: 80}, imgui.CondOnce)
			imgui.BeginV("运行状态", nil, 0)
			imgui.Text("脚本正在后台运行...")
			imgui.End()
		}
	})

	// 阻塞主进程
	select {}
}
