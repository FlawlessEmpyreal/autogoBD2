package info

import "image"

type StructColorCmp struct {
	X     int
	Y     int
	Color string
}

var BattleInterface = []StructColorCmp{
	{X: 1122, Y: 626, Color: "dedad6"},
	{X: 194, Y: 618, Color: "ffffff"},
	{X: 1186, Y: 42, Color: "ffffff"},
	{X: 937, Y: 42, Color: "fefefd"},
	{X: 854, Y: 42, Color: "fefefd"},
}

var TpInterface1 = []StructColorCmp{
	{X: 169, Y: 41, Color: "dedad2"},
	{X: 182, Y: 53, Color: "dedad2"},
	{X: 182, Y: 28, Color: "dad6ce"},
}
var TpInterface2 = []StructColorCmp{
	{X: 182, Y: 49, Color: "d9d6cd"},
	{X: 182, Y: 34, Color: "dcdad2"},
	{X: 187, Y: 42, Color: "dedad2"},
}

var Chapter1_1ColorCmp = []StructColorCmp{
	{X: 562, Y: 231, Color: "a19f94"},
	{X: 752, Y: 348, Color: "a09e93"},
	{X: 604, Y: 405, Color: "9f9d92"},
	{X: 709, Y: 390, Color: "9f9d92"},
}

var Chapter1_1TpPoint = image.Point{525, 372}

var Chapter1_2ColorCmp = []StructColorCmp{
	{X: 508, Y: 231, Color: "9e9b93"},
	{X: 648, Y: 317, Color: "9b9a8e"},
	{X: 503, Y: 420, Color: "9c9a90"},
	{X: 716, Y: 296, Color: "9a9990"},
}
var Chapter1_2TpPoint = image.Point{543, 294}

var ChapterSelectDetect = []StructColorCmp{
	{X: 103, Y: 702, Color: "101821"},
	{X: 605, Y: 704, Color: "101821"},
	{X: 1163, Y: 703, Color: "101821"},
}

var If_Map = []StructColorCmp{ //判断是否在跑图界面
	{X: 321, Y: 269, Color: "ffffff"},
	{X: 1096, Y: 562, Color: "dedad6"},
	{X: 1091, Y: 54, Color: "ffffff"},
	{X: 317, Y: 83, Color: "ffffff"},
}

var If_GenerateTP = []StructColorCmp{ //判断是否在跑图界面
	{X: 477, Y: 442, Color: "dedad6"},
	{X: 614, Y: 446, Color: "dedad6"},
	{X: 660, Y: 443, Color: "dedad6"},
	{X: 802, Y: 443, Color: "dedad6"},
}

var LoadingInterface = []StructColorCmp{
	{X: 191, Y: 149, Color: "000000"},
	{X: 190, Y: 343, Color: "000000"},
	{X: 299, Y: 400, Color: "000000"},
	{X: 383, Y: 163, Color: "000000"},
	{X: 63, Y: 42, Color: "000000"},
}

var TpLeftButten = []StructColorCmp{
	{X: 300, Y: 344, Color: "dedad6"},
	{X: 305, Y: 338, Color: "dedace"},
	{X: 304, Y: 349, Color: "dddbd5"},
}

var TpRightButten = []StructColorCmp{
	{X: 979, Y: 344, Color: "dedad6"},
	{X: 974, Y: 338, Color: "dedace"},
	{X: 975, Y: 349, Color: "dddbd5"},
}
