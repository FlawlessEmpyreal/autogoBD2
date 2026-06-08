package mainProcess

func MainProcess() {

	Ctrl := NewController()
	RegisterChapters(Ctrl)

	//执行全部启用章节
	Ctrl.RunAll()

	//执行指定章节
	//CtrlMainStoryline.RunByName("chapter2", "chapter3")
}
