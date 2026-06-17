package mainProcess

import (
	"app/info"
	"app/info/info_chapter"
)

func RegChapter1(Ctrl *Controller) {
	Ctrl.Register(&ChapterConfig{
		Name:       "chapter1",
		Enabled:    true,
		MaxRetries: 3,
		Stages: []Stage{
			{
				Name:  "FindChapter",
				type_: "FindChapter",
				Run: func() (info.RecoveryAction, error) {
					return FindChapter(
						info_chapter.Ch1.ChapterImg_path,
						info_chapter.Ch1.Type_,
					)
				},
			},
			{
				Name:  "Tp_Map1",
				type_: "Tp",
				Run: func() (info.RecoveryAction, error) {
					return Chapter_Tp(
						info_chapter.Ch1.MapConfig[0].MapFind.MapColorsCmp,
						info_chapter.Ch1.MapConfig[0].TpPoint,
					)
				},
			},
			{
				Name:  "RunChapter1_1",
				type_: "RunChapter",
				Run: func() (info.RecoveryAction, error) {
					return ChapterRun(
						info_chapter.Ch1.MapConfig[0].BigMapPath,
						info_chapter.Ch1.MapConfig[0].Bin_mapPath,
						info_chapter.Ch1.MapConfig[0].MonsterLocation,
						info_chapter.Ch1.YoloLable,
						"main",
					)
				},
			},
			{
				Name:  "Tp_Map2",
				type_: "Tp",
				Run: func() (info.RecoveryAction, error) {
					return Chapter_Tp(
						info_chapter.Ch1.MapConfig[1].MapFind.MapColorsCmp,
						info_chapter.Ch1.MapConfig[1].TpPoint,
					)
				},
			},
			{
				Name:  "RunChapter1_2",
				type_: "RunChapter",
				Run: func() (info.RecoveryAction, error) {
					return ChapterRun(
						info_chapter.Ch1.MapConfig[1].BigMapPath,
						info_chapter.Ch1.MapConfig[1].Bin_mapPath,
						info_chapter.Ch1.MapConfig[1].MonsterLocation,
						info_chapter.Ch1.YoloLable,
						"main",
					)
				},
			},
		},
	})
}

func RegChapter2(Ctrl *Controller) {
	Ctrl.Register(&ChapterConfig{
		Name:       "chapter2",
		Enabled:    true,
		MaxRetries: 3,
		Stages: []Stage{
			{
				Name:  "FindChapter",
				type_: "FindChapter",
				Run: func() (info.RecoveryAction, error) {
					return FindChapter(
						info_chapter.Ch2.ChapterImg_path,
						info_chapter.Ch2.Type_,
					)
				},
			},
			{
				Name:  "Tp_Map1",
				type_: "Tp",
				Run: func() (info.RecoveryAction, error) {
					return Chapter_Tp(
						info_chapter.Ch2.MapConfig[0].MapFind.MapColorsCmp,
						info_chapter.Ch2.MapConfig[0].TpPoint,
					)
				},
			},
			{
				Name:  "RunChapter2_1",
				type_: "RunChapter",
				Run: func() (info.RecoveryAction, error) {
					return ChapterRun(
						info_chapter.Ch2.MapConfig[0].BigMapPath,
						info_chapter.Ch2.MapConfig[0].Bin_mapPath,
						info_chapter.Ch2.MapConfig[0].MonsterLocation,
						info_chapter.Ch2.YoloLable,
						"main",
					)
				},
			},
			{
				Name:  "Tp_Map2",
				type_: "Tp",
				Run: func() (info.RecoveryAction, error) {
					return Chapter_Tp(
						info_chapter.Ch2.MapConfig[1].MapFind.MapColorsCmp,
						info_chapter.Ch2.MapConfig[1].TpPoint,
					)
				},
			},
			{
				Name:  "RunChapter2_2",
				type_: "RunChapter",
				Run: func() (info.RecoveryAction, error) {
					return ChapterRun(
						info_chapter.Ch2.MapConfig[1].BigMapPath,
						info_chapter.Ch2.MapConfig[1].Bin_mapPath,
						info_chapter.Ch2.MapConfig[1].MonsterLocation,
						info_chapter.Ch2.YoloLable,
						"main",
					)
				},
			},
		},
	})
}
