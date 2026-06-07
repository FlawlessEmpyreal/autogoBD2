package mainProcess

func MainProcess() {

	CtrlMainStoryline := NewController()
	RegisterChapters(CtrlMainStoryline)

	//执行全部启用章节
	CtrlMainStoryline.RunAll()

	//执行指定章节
	//CtrlMainStoryline.RunByName("chapter2", "chapter3")
}
