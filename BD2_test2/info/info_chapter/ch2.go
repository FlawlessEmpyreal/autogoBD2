package info_chapter

import (
	"app/info"
	"image"
	"path"
)

// 第二章信息
var Ch2 = info.ChapterConfig{
	ChapterName:     "chapter2",
	ChapterImg_path: path.Join(info.ImgRoot, "misc/chapter2Select.png"),
	Type_:           "main",
	YoloLable:       "chapter2_monster",
	MapConfig: []info.MapConfig{
		{
			MapName:     1,
			TpPoint:     image.Point{577, 324},
			BigMapPath:  path.Join(info.ImgRoot, "map/scaled_grey_Extend_chapter2_1.jpg"),
			Bin_mapPath: path.Join(info.ImgRoot, "map/bin_map_chapter2_1.jpg"),
			MapFind: info.MapFind{
				Function: "第二章地图1界面找色",
				MapColorsCmp: []info.StructColorCmp{
					{X: 735, Y: 202, Color: "9b9a90"},
					{X: 630, Y: 219, Color: "9c998f"},
					{X: 551, Y: 294, Color: "9d9d94"},
					{X: 601, Y: 449, Color: "98958a"},
				},
			},
			MonsterLocation: []image.Point{
				{X: 138, Y: 205},
				{X: 217, Y: 287},
				{X: 308, Y: 328},
				{X: 283, Y: 83},
				{X: 400, Y: 106},
			},
		},
		{
			MapName:     2,
			TpPoint:     image.Point{596, 371},
			BigMapPath:  path.Join(info.ImgRoot, "map/scaled_grey_Extend_chapter2_2.jpg"),
			Bin_mapPath: path.Join(info.ImgRoot, "map/bin_map_chapter2_2.jpg"),
			MapFind: info.MapFind{
				Function: "第二章地图2界面找色",
				MapColorsCmp: []info.StructColorCmp{
					{X: 559, Y: 445, Color: "969183"},
					{X: 600, Y: 254, Color: "99978b"},
					{X: 535, Y: 465, Color: "58cd80"},
					{X: 694, Y: 233, Color: "9d9d94"},
				},
			},
			MonsterLocation: []image.Point{
				{X: 193, Y: 357},
				{X: 120, Y: 417},
				{X: 213, Y: 211},
				{X: 236, Y: 104},
				{X: 354, Y: 98},
			},
		},
	},
}
