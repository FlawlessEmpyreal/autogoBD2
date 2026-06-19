package info_chapter

import (
	"app/info"
	"image"
	"path"
)

var Ch3 = info.ChapterConfig{
	ChapterName:     "chapter3",
	ChapterImg_path: path.Join(info.ImgRoot, "misc/chapter3Select.png"),
	Type_:           "main",
	YoloLable:       "chapter3_monster",
	MapConfig: []info.MapConfig{
		{
			MapName:    1,
			TpPoint:    image.Point{614, 302},
			BigMapPath: path.Join(info.ImgRoot, "map/scaled_grey_Extend_chapter3_1.jpg"),
			//Bin_mapPath: path.Join(info.ImgRoot, "map/bin_map_chapter2_1.jpg"),
			Bin_mapPath: path.Join(info.ImgRoot, "map/bin_map_chapter3_1.jpg"),
			MapFind: info.MapFind{
				Function: "第三章地图1界面找色",
				MapColorsCmp: []info.StructColorCmp{
					{X: 567, Y: 385, Color: "938f82"},
					{X: 676, Y: 324, Color: "99978c"},
					{X: 751, Y: 176, Color: "ea5551"},
					{X: 531, Y: 394, Color: "54c279"},
				},
			},
			MonsterLocation: []image.Point{
				{X: 356, Y: 148},
				{X: 304, Y: 273},
				{X: 181, Y: 332},
				{X: 225, Y: 434},
				{X: 150, Y: 375},
			},
		},
		{
			MapName:     2,
			TpPoint:     image.Point{659, 214},
			BigMapPath:  path.Join(info.ImgRoot, "map/scaled_grey_Extend_chapter3_2.jpg"),
			Bin_mapPath: path.Join(info.ImgRoot, "map/bin_map_chapter3_2.jpg"),
			MapFind: info.MapFind{
				Function: "第三章地图2界面找色",
				MapColorsCmp: []info.StructColorCmp{
					{X: 741, Y: 174, Color: "898576"},
					{X: 572, Y: 227, Color: "a09e96"},
					{X: 569, Y: 469, Color: "df524d"},
					{X: 664, Y: 419, Color: "98978b"},
				},
			},
			MonsterLocation: []image.Point{
				{X: 400, Y: 98},
				{X: 193, Y: 150},
				{X: 240, Y: 250},
				{X: 371, Y: 350},
				{X: 254, Y: 362},
				{X: 212, Y: 313},
				{X: 174, Y: 422},
			},
		},
	},
}
