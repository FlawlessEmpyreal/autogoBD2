package imGui

import (
	"app/info"
	"app/mainProcess"

	"github.com/Dasongzi1366/AutoGo/imgui"
)

func ImGuiRun() {
	// 初始化 ImGui
	imgui.Init()

	// 1. 定义状态变量
	showWindow := true // 控制窗口是否显示
	isRunning := false // 控制脚本是否正在运行
	info.RegChAll = false
	info.RegCh1 = false // 复选框绑定的布尔变量
	info.RegCh2 = false

	// 2. 主循环
	imgui.Run(func() {
		// 如果脚本正在运行，可以选择不显示窗口，或者显示一个“运行中”的提示
		if !isRunning && showWindow {
			// 设置窗口大小，只在首次创建时生效
			imgui.SetNextWindowSizeV(imgui.Vec2{X: 900, Y: 500}, imgui.CondOnce)

			// 创建窗口。第二个参数传入 &showWindow，点击窗口右上角的 X 也会关闭
			imgui.BeginV("AutoBD2", &showWindow, 0)

			// 添加复选框控件
			// 勾选时自动变为 true，取消勾选变为 false
			imgui.Checkbox("All", &info.RegChAll)
			if info.RegChAll {
				info.RegCh1 = true
				info.RegCh2 = true
			}

			//渲染子选项复选框
			imgui.Checkbox("ch1", &info.RegCh1)
			imgui.Checkbox("ch2", &info.RegCh2)

			//子选项状态 -> 全选状态（向上同步）
			// 如果用户手动取消了 ch1 或 ch2，全选状态必须立刻变为 false
			if !info.RegCh1 || !info.RegCh2 {
				info.RegChAll = false
			}

			imgui.Spacing()
			imgui.Separator()
			imgui.Spacing()

			// 4. 添加运行按钮
			if imgui.Button("点击运行") {
				isRunning = true
				showWindow = false // 点击后隐藏窗口

				// 启动你的脚本协程
				go func() {
					// 模拟脚本运行耗时
					mainProcess.MainProcess()
					//fmt.Println("脚本运行结束")

					// 脚本运行完毕后，可以重新显示窗口
					isRunning = false
					showWindow = true
				}()
			}

			imgui.End()
		}
		//} else if isRunning {
		// 脚本运行中，显示一个小的提示窗口（可选）
		//imgui.SetNextWindowSizeV(imgui.Vec2{X: 200, Y: 80}, imgui.CondOnce)
		//imgui.BeginV("运行状态", nil, 0)
		//imgui.Text("脚本正在后台运行...")
		//imgui.End()
		//}
	})

	// 阻塞主进程
	select {}
}
