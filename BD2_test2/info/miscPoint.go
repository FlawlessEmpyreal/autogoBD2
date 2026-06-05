package info

import "image"

type pointList struct {
	MainChapterChoose   image.Point
	BranchChapterChoose image.Point
}

var MiscPoint = pointList{
	MainChapterChoose:   image.Point{273, 558},
	BranchChapterChoose: image.Point{450, 558},
}
