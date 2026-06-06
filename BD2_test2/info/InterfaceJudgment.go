package info

// 判断界面用
var IF = InterfaceJudgment{
	IF_Battle: SceneConfig{
		ColorsCmp: []StructColorCmp{
			{X: 1122, Y: 626, Color: "dedad6"},
			{X: 194, Y: 618, Color: "ffffff"},
			{X: 1186, Y: 42, Color: "ffffff"},
			{X: 937, Y: 42, Color: "fefefd"},
			{X: 854, Y: 42, Color: "fefefd"},
		},
	},
	IF_TpInterface: [2]SceneConfig{
		{
			ColorsCmp: []StructColorCmp{
				{X: 169, Y: 41, Color: "dedad2"},
				{X: 182, Y: 53, Color: "dedad2"},
				{X: 182, Y: 28, Color: "dad6ce"},
			},
		},
		{
			ColorsCmp: []StructColorCmp{
				{X: 182, Y: 49, Color: "d9d6cd"},
				{X: 182, Y: 34, Color: "dcdad2"},
				{X: 187, Y: 42, Color: "dedad2"},
			},
		},
	},
	If_Map: SceneConfig{
		ColorsCmp: []StructColorCmp{ //判断是否在跑图界面
			{X: 321, Y: 269, Color: "ffffff"},
			{X: 1096, Y: 562, Color: "dedad6"},
			{X: 1091, Y: 54, Color: "ffffff"},
			{X: 317, Y: 83, Color: "ffffff"},
		},
	},
	If_GenerateTP: SceneConfig{
		ColorsCmp: []StructColorCmp{ //判断是否生成传送
			{X: 477, Y: 442, Color: "dedad6"},
			{X: 614, Y: 446, Color: "dedad6"},
			{X: 660, Y: 443, Color: "dedad6"},
			{X: 802, Y: 443, Color: "dedad6"},
		},
	},
	If_LoadingInterface: SceneConfig{
		ColorsCmp: []StructColorCmp{
			{X: 191, Y: 149, Color: "000000"},
			{X: 190, Y: 343, Color: "000000"},
			{X: 299, Y: 400, Color: "000000"},
			{X: 383, Y: 163, Color: "000000"},
			{X: 63, Y: 42, Color: "000000"},
		},
	},
	If_TpLeftButton: SceneConfig{
		ColorsCmp: []StructColorCmp{
			{X: 300, Y: 344, Color: "dedad6"},
			{X: 305, Y: 338, Color: "dedace"},
			{X: 304, Y: 349, Color: "dddbd5"},
		},
	},
	If_TpRightButton: SceneConfig{
		ColorsCmp: []StructColorCmp{
			{X: 979, Y: 344, Color: "dedad6"},
			{X: 974, Y: 338, Color: "dedace"},
			{X: 975, Y: 349, Color: "dddbd5"},
		},
	},
	If_Backbutton: [2]SceneConfig{
		{
			ColorsCmp: []StructColorCmp{
				{X: 134, Y: 39, Color: "ffffff"},
				{X: 136, Y: 36, Color: "fefefe"},
				{X: 138, Y: 45, Color: "ededed"},
			},
		},
		{
			ColorsCmp: []StructColorCmp{
				{X: 103, Y: 40, Color: "ffffff"},
				{X: 108, Y: 35, Color: "f4f4f4"},
				{X: 108, Y: 45, Color: "f2f2f2"},
			},
		},
	},
	If_BattleFieldRole: SceneConfig{
		ColorsCmp: []StructColorCmp{
			{X: 728, Y: 575, Color: "dedad6"},
			{X: 879, Y: 581, Color: "dedad6"},
			{X: 948, Y: 578, Color: "dedad6"},
			{X: 1094, Y: 579, Color: "dedad6"},
		},
	},
}
