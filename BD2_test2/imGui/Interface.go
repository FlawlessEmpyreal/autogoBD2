package imGui

// =====================================================================================
// 依赖说明（必须先在 info 包里做以下修改，否则本文件编译不过）：
//
//   package info
//
//   var RegChAll bool
//   var RegCh = make([]bool, 30) // 下标 0~29 对应第 1~30 章的勾选状态
//
// 原来的 RegCh1 / RegCh2 两个独立布尔量，替换成了一个 []bool 切片，
// 这样无论是 2 章还是 30 章都是同一套逻辑，不用每加一章就加一个变量。
// mainProcess.MainProcess() 里如果直接用到了 info.RegCh1 / info.RegCh2，
// 也需要同步改成遍历 info.RegCh，判断哪些下标为 true。
// =====================================================================================

import (
	"app/info"
	"app/mainProcess"
	"fmt"

	"github.com/Dasongzi1366/AutoGo/imgui"
)

// ------------------------------- 基础配置 -------------------------------

var totalChapters = len(mainProcess.ChapterHandlers) // 主线章节总数
const (
	chaptersPerRow = 6   // 每行排几个，对应截图里的横排勾选框
	sidebarWidth   = 150 //右侧功能栏固定宽度
)

// ------------------------------- 配色方案 -------------------------------
// 整体走深色界面 + 青柚色（teal）点缀，比纯黑底配荧光绿更柔和、更耐看。
// 如果想换主题色，只改 colAccent 系列三个值即可，其余颜色会自动呼应。

var (
	colWindowBg  = imgui.Vec4{X: 0.10, Y: 0.11, Z: 0.15, W: 1.00} // 主窗口背景
	colLeftBg    = imgui.Vec4{X: 0.13, Y: 0.14, Z: 0.18, W: 1.00} // 左侧内容区背景
	colSidebarBg = imgui.Vec4{X: 0.08, Y: 0.09, Z: 0.12, W: 1.00} // 右侧功能栏背景，更深一层做区分

	colAccent       = imgui.Vec4{X: 0.30, Y: 0.80, Z: 0.70, W: 1.00} // 主题强调色：青柚绿
	colAccentSoft   = imgui.Vec4{X: 0.30, Y: 0.80, Z: 0.70, W: 0.55} // 半透明版本，给 hover/选中态用
	colAccentHover  = imgui.Vec4{X: 0.36, Y: 0.86, Z: 0.76, W: 1.00}
	colAccentActive = imgui.Vec4{X: 0.26, Y: 0.74, Z: 0.64, W: 1.00}

	colFrameBg        = imgui.Vec4{X: 0.19, Y: 0.21, Z: 0.26, W: 1.00} // 复选框未勾选时的底色
	colFrameBgHovered = imgui.Vec4{X: 0.24, Y: 0.27, Z: 0.33, W: 1.00}
	colFrameBgActive  = imgui.Vec4{X: 0.22, Y: 0.24, Z: 0.30, W: 1.00}

	colSeparator = imgui.Vec4{X: 0.30, Y: 0.80, Z: 0.70, W: 0.35} // 每行下方的分隔线，呼应主题色但偏淡
	colText      = imgui.Vec4{X: 0.93, Y: 0.94, Z: 0.96, W: 1.00}
	colTextDim   = imgui.Vec4{X: 0.60, Y: 0.62, Z: 0.68, W: 1.00} // 说明性文字，弱化显示
)

// ------------------------------- 功能页签 -------------------------------
// 右侧"功能选项"栏点击后切换左侧显示的内容，类似截图里的翻页效果。
// 以后想加新功能页，只要在 pageList 里加一行、再写一个 renderXxxPage() 函数即可。

type pageID int

const (
	pageChapters pageID = iota
	pageSettings
	pageAbout
)

var pageList = []struct {
	id    pageID
	label string
}{
	{pageChapters, "撞怪"},
	{pageSettings, "跑商"},
	{pageAbout, "日志"},
}

// ------------------------------- 运行时状态 -------------------------------

var (
	currentPage pageID = pageChapters //记录当前显示的是哪个页面
	showWindow         = true
	isRunning          = false //防止任务重复启动的锁
)

// ------------------------------- 入口函数 -------------------------------

func ImGuiRun() {
	imgui.Init()

	info.RegChAll = false
	if info.RegCh == nil || len(info.RegCh) != totalChapters {
		info.RegCh = make([]bool, totalChapters)
	}

	imgui.Run(func() {
		if isRunning || !showWindow {
			return
		}

		imgui.SetNextWindowSizeV(imgui.Vec2{X: 920, Y: 560}, imgui.CondOnce)

		// ---- 整体主题色入栈，作用到这一帧的整个窗口 ----
		imgui.PushStyleColorVec4(imgui.ColWindowBg, colWindowBg)
		imgui.PushStyleColorVec4(imgui.ColChildBg, colLeftBg)
		imgui.PushStyleColorVec4(imgui.ColText, colText)
		imgui.PushStyleColorVec4(imgui.ColSeparator, colSeparator)
		imgui.PushStyleColorVec4(imgui.ColButton, colAccent)
		imgui.PushStyleColorVec4(imgui.ColButtonHovered, colAccentHover)
		imgui.PushStyleColorVec4(imgui.ColButtonActive, colAccentActive)
		imgui.PushStyleColorVec4(imgui.ColHeader, colAccentSoft) // Selectable 选中态背景
		imgui.PushStyleColorVec4(imgui.ColHeaderHovered, colAccentHover)
		imgui.PushStyleColorVec4(imgui.ColHeaderActive, colAccent)

		imgui.BeginV("AutoBD2", &showWindow, 0)

		renderBody() // 核心布局渲染

		imgui.End()

		imgui.PopStyleColorV(10) // 对应上面 10 次 PushStyleColor，必须成对
	})

	select {}
}

// ------------------------------- 左右两栏布局 -------------------------------

func renderBody() {
	avail := imgui.ContentRegionAvail()
	leftWidth := avail.X - sidebarWidth - 12 // 12 是左右两栏之间留的空隙

	// 左侧面板
	if leftWidth < 200 {
		leftWidth = 200
	}
	imgui.BeginChildStrV("left_panel", imgui.Vec2{X: leftWidth, Y: 0}, imgui.ChildFlagsAlwaysAutoResize, 0) //创建子窗口容器。左侧宽度自适应，右侧固定宽度。
	switch currentPage {
	case pageChapters:
		renderChapterPage()
	case pageSettings:
		renderSettingsPage()
	case pageAbout:
		renderAboutPage()
	}
	imgui.EndChild()

	imgui.SameLine() //取消换行，让左右两个子窗口并排显示。

	// 右侧面板
	imgui.PushStyleColorVec4(imgui.ColChildBg, colSidebarBg)
	imgui.BeginChildStrV("right_panel", imgui.Vec2{X: sidebarWidth, Y: 0}, imgui.ChildFlagsAlwaysAutoResize, 0)
	renderSidebar()
	imgui.EndChild()
	imgui.PopStyleColor()
}

// ------------------------------- 右侧：功能选项栏 -------------------------------

func renderSidebar() {
	imgui.Text("选项")
	imgui.Spacing()
	imgui.Separator()
	imgui.Spacing()

	//点击按钮时，更新 currentPage 变量，从而触发左侧内容切换。
	for _, p := range pageList {
		selected := currentPage == p.id

		if imgui.SelectableBoolV(p.label, selected, 0, imgui.Vec2{X: 0, Y: 32}) { //创建可选中的按钮，高度固定为 32。
			currentPage = p.id
		}
		imgui.Spacing()
	}
}

// ------------------------------- 左侧：章节选择页 -------------------------------

func renderChapterPage() {
	imgui.Text("章节选择")
	imgui.Spacing()

	if imgui.Checkbox("All", &info.RegChAll) { //全选功能
		applySelectAllToIndividual()
	}
	imgui.SameLineV(0, 24)
	runClicked := imgui.Button("Run")

	imgui.Spacing()
	imgui.Separator()
	imgui.Spacing()

	renderChapterGrid()

	if runClicked {
		isRunning = true
		showWindow = false

		go func() {
			mainProcess.MainProcess()
			isRunning = false
			showWindow = true
		}()
	}
}

// renderChapterGrid 画出截图里那种横排小勾选框、数字写在框下面、每行一条分隔线的效果。
func renderChapterGrid() {
	// 缩小复选框尺寸 + 换主题色，只在这个区域内生效，画完就弹出，不影响其他控件
	imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 4, Y: 4})
	imgui.PushStyleColorVec4(imgui.ColFrameBg, colFrameBg)
	imgui.PushStyleColorVec4(imgui.ColFrameBgHovered, colFrameBgHovered)
	imgui.PushStyleColorVec4(imgui.ColFrameBgActive, colFrameBgActive)
	imgui.PushStyleColorVec4(imgui.ColCheckMark, colAccent)

	boxSize := imgui.FrameHeight() // 当前样式下复选框的边长，用来把数字居中对齐到框下面

	for idx := 0; idx < totalChapters; idx++ {
		imgui.PushIDInt(int32(idx))
		imgui.BeginGroup()

		groupStartX := imgui.CursorPosX()

		checked := info.RegCh[idx]
		if imgui.Checkbox("", &checked) {
			info.RegCh[idx] = checked
			syncSelectAllFromIndividual()
		}

		label := fmt.Sprintf("%s", mainProcess.ChapterHandlers[idx].RegName)
		//		label := fmt.Sprintf("%d", idx+1)
		textSize := imgui.CalcTextSize(label)
		offsetX := (boxSize - textSize.X) / 2
		if offsetX < 0 {
			offsetX = 0
		}
		imgui.SetCursorPos(imgui.Vec2{X: groupStartX + offsetX, Y: imgui.CursorPosY()})
		imgui.Text(label)

		imgui.EndGroup()
		imgui.PopID()

		isRowEnd := (idx+1)%chaptersPerRow == 0
		isLast := idx == totalChapters-1

		if !isRowEnd {
			imgui.SameLine()
		} else if !isLast {
			imgui.Spacing()
			imgui.Separator()
			imgui.Spacing()
		}
	}

	imgui.PopStyleColorV(4)
	imgui.PopStyleVar()
}

// ------------------------------- 左侧：运行设置页（占位，按需扩展） -------------------------------

func renderSettingsPage() {
	imgui.Text("跑商")
	imgui.Spacing()
	imgui.Separator()
	imgui.Spacing()

	imgui.PushStyleColorVec4(imgui.ColText, colTextDim)
	imgui.Text("这里可以放延迟时间、循环次数之类的参数设置。")
	imgui.Text("先留空，等需要的时候再加对应的控件。")
	imgui.PopStyleColor()
}

// ------------------------------- 左侧：关于页（占位） -------------------------------

func renderAboutPage() {
	imgui.Text("Log")
	imgui.Spacing()
	imgui.Separator()
	imgui.Spacing()
	imgui.Text("AutoBD2 自动化脚本")

	imgui.PushStyleColorVec4(imgui.ColText, colTextDim)
	imgui.Text("此页面仅作占位展示，方便后续扩展更多功能选项。")
	imgui.PopStyleColor()
}

// ------------------------------- 全选 <-> 单项 状态同步 -------------------------------

// applySelectAllToIndividual 在用户点击"全选"复选框后，把状态广播给所有单项。
func applySelectAllToIndividual() {
	for i := range info.RegCh {
		info.RegCh[i] = info.RegChAll
	}
}

// syncSelectAllFromIndividual 在任意单项变化后，重新计算"全选"该不该是勾选状态。
// 只要有一个没勾上，"全选"就要变回未勾选；全部勾上才让"全选"变成勾选。
func syncSelectAllFromIndividual() {
	for _, v := range info.RegCh {
		if !v {
			info.RegChAll = false
			return
		}
	}
	info.RegChAll = true
}
