package info

//第一章信息

import (
	"image"
	"path"
)

// 第一章信息
var Ch1 = ChapterConfig{
	ChapterName:     "chapter1",
	ChapterImg_path: path.Join(imgRoot, "misc/chapter1Selcet.png"),
	Type_:           "main",
	MapConfig: []MapConfig{
		{
			MapName:     1,
			TpPoint:     image.Point{525, 372},
			BigMapPath:  path.Join(imgRoot, "map/scaled_grey_Extend_chapter1_1.jpg"),
			Bin_mapPath: path.Join(imgRoot, "map/bin_map_chapter1_1.jpg"),
			MapFind: MapFind{
				Function: "第一章地图1界面找色",
				MapColorsCmp: []StructColorCmp{
					{X: 562, Y: 231, Color: "a19f94"},
					{X: 752, Y: 348, Color: "a09e93"},
					{X: 604, Y: 405, Color: "9f9d92"},
					{X: 709, Y: 390, Color: "9f9d92"},
				},
			},
			MonsterLocation: []image.Point{
				{X: 256, Y: 397},
				{X: 338, Y: 380},
				{X: 420, Y: 332},
				{X: 304, Y: 233},
				{X: 175, Y: 191},
			},
		},
		{
			MapName:     2,
			TpPoint:     image.Point{543, 294},
			BigMapPath:  path.Join(imgRoot, "map/scaled_grey_Extend_chapter1_2.jpg"),
			Bin_mapPath: path.Join(imgRoot, "map/bin_map_chapter1_2.jpg"),
			MapFind: MapFind{
				Function: "第一章地图2界面找色",
				MapColorsCmp: []StructColorCmp{
					{X: 508, Y: 231, Color: "9e9b93"},
					{X: 648, Y: 317, Color: "9b9a8e"},
					{X: 503, Y: 420, Color: "9c9a90"},
					{X: 716, Y: 296, Color: "9a9990"},
				},
			},
			MonsterLocation: []image.Point{
				{X: 244, Y: 223},
				{X: 296, Y: 264},
				{X: 203, Y: 419},
				{X: 391, Y: 228},
				{X: 416, Y: 132},
			},
		},
	},
}
