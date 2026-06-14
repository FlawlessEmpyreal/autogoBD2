package info

import "image"

type pointList struct {
	MainChapterChoose   image.Point //章节选择时选主线
	BranchChapterChoose image.Point //章节选择时选支线
	BackButton1         image.Point //返回按钮1
	BackButton2         image.Point //返回按钮2
	BackButton          image.Point
}

var MiscPoint = pointList{
	MainChapterChoose:   image.Point{273, 558},
	BranchChapterChoose: image.Point{450, 558},
	BackButton1:         image.Point{136, 40},
	BackButton2:         image.Point{105, 40},
	BackButton:          image.Point{165, 40},
}
