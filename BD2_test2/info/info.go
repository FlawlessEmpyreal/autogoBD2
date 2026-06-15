package info

import (
	"image"
	"path"
	"time"
)

// 使用技能

var (
	//━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// ====================Dir======================
	ImgRoot = "/mnt/shared/Pictures/img"

	YoloParamPath_Main   = path.Join(ImgRoot, "misc/best_main_1to17.param") //yolo param文件路径
	YoloBinPath_Main     = path.Join(ImgRoot, "misc/best_main_1to17.bin")   //yolo bin文件路径
	YoloParamPath_Branch = path.Join(ImgRoot, "misc/")
	YoloBinPath_Branch   = path.Join(ImgRoot, "misc/")

	ChapterSelectInterfaceImg_path = path.Join(ImgRoot, "misc/chapterSelectDetect.png")
	BattleInterfaceDetect_path     = path.Join(ImgRoot, "misc/BattleInterfaceDetect.jpg")
	BackButton_path                = path.Join(ImgRoot, "misc/backButton.png")
	EscapeImg_path                 = path.Join(ImgRoot, "misc/escapeDetect.png")
	ChapterSelectButtonImg_path    = path.Join(ImgRoot, "misc/ChapterSelectButton")
	//━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	//=====================labels===============

	//━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	//===================skill para=============

	Accelerate bool = true  //加速
	Subdue     bool = true  //压制
	Stealth    bool = false //隐身

	//━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	//================reg para==============

	//是否注册章节
	RegCh1 = true
	RegCh2 = true

	//━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	//================astar para==============
	MaxDist     = 20.0 //路径简化后,路径之间的最大像素距离
	SuccessDist = 15.0 //目标在这个范围内看作寻路完成
	ErrDist     = 40.0 //离下一个目标超过这个距离算作寻路超距, 注意: MaxDist + SuccessDist < ErrDist

	PathfindingRetryCount = 15                    //寻路重试次数,超过这个次数 超时计数+1
	TimeoutCount          = 2                     //超时计数大于这个数返回超时
	NavigationInterval    = 16 * time.Millisecond //寻路间隔,默认一秒六十次

	IsolateRadius = 2
)

//=================================================================================

type ButtenPoint struct {
	Accelerate image.Point
	Stealth    image.Point
	Subdue     image.Point
}

var BP = ButtenPoint{
	Accelerate: image.Point{1000, 650},
	Subdue:     image.Point{967, 561},
	Stealth:    image.Point{1000, 480},
}

type StructColorCmp struct { //找色点
	X     int
	Y     int
	Color string
}

// 地图配置（颜色识别 + 传送点）
type MapConfig struct { //每章节的地图设置
	MapName         int           //该章节的第几张地图
	TpPoint         image.Point   //传送点坐标
	BigMapPath      string        //匹配坐标的地图路径
	Bin_mapPath     string        //二值化地图路径
	MapFind         MapFind       //找地图用的找色信息
	MonsterLocation []image.Point //怪物位置
}

type MapFind struct { //找地图用的找色点
	Function     string           //说明性质的文本
	MapColorsCmp []StructColorCmp //找色点
}

// 每个章节的信息汇总
type ChapterConfig struct {
	ChapterName     string      //章节名称
	ChapterImg_path string      //找章节时用的图片
	MapConfig       []MapConfig //地图信息
	Type_           string      //主线还是支线
	YoloLable       string
}

type SceneConfig struct {
	ColorsCmp []StructColorCmp
}
type InterfaceJudgment struct { // 界面识别配置
	IF_Battle           SceneConfig    // 判断是否在战斗界面
	IF_TpInterface      [2]SceneConfig // tp界面有两种                     判断时使用MyOpencv.If_TpInterface
	If_Map              SceneConfig    // 判断是否在跑图界面
	If_GenerateTP       SceneConfig    // 判断是否生成传送
	If_LoadingInterface SceneConfig    // 判断是否在加载界面
	If_TpLeftButton     SceneConfig    // 地图中选择地图的左按钮找色
	If_TpRightButton    SceneConfig    // 地图中选择地图的右按钮找色
	If_Backbutton       [2]SceneConfig // 判断是否有返回按钮，返回按钮有两种   判断时使用MyOpencv.If_Backbutton
	If_BattleFieldRole  SceneConfig    // 判断是否在战场角色设置界面
	If_Pause            SceneConfig    //判断是否在战斗暂停页面
	If_escape           SceneConfig
}

type RecoveryAction int

const (
	StageDone           RecoveryAction = iota //阶段成功
	RetryStage                                // 重试当前阶段
	RetryChapter                              // 从第一个阶段重来
	GoToRunMapInterface                       // 回战斗界面，处理完后重试当前阶段
	SkipStage                                 // 跳过当前阶段，继续下一阶段
	SkipChapter                               // 跳过整个章节
	AbortAll                                  // 终止所有
	WayHandleDone                             //寻路异常处理完成
)

type WayFindState int

const (
	STATE_Done             WayFindState = iota // 正常到达终点
	STATE_Movement_timeout                     // 寻路超时
	STATE_CDT_Useless                          // 坐标异常
	STATE_Loss_Map                             //小地图丢失
	STATE_Fixed_succes                         //handleLossMap 修理成功
	STATE_Fixed_Failed                         //handleLossMap 修理失败
)

// 判断是否在战斗界面
//var BattleInterface = []StructColorCmp{
//	{X: 1122, Y: 626, Color: "dedad6"},
//	{X: 194, Y: 618, Color: "ffffff"},
//	{X: 1186, Y: 42, Color: "ffffff"},
//	{X: 937, Y: 42, Color: "fefefd"},
//	{X: 854, Y: 42, Color: "fefefd"},
//}

// tp界面有两种
//var TpInterface1 = []StructColorCmp{
//	{X: 169, Y: 41, Color: "dedad2"},
//	{X: 182, Y: 53, Color: "dedad2"},
//	{X: 182, Y: 28, Color: "dad6ce"},
//}
//var TpInterface2 = []StructColorCmp{
//	{X: 182, Y: 49, Color: "d9d6cd"},
//	{X: 182, Y: 34, Color: "dcdad2"},
//	{X: 187, Y: 42, Color: "dedad2"},
//}

// 第一章地图1界面找色
//var Chapter1_1ColorCmp = []StructColorCmp{
//	{X: 562, Y: 231, Color: "a19f94"},
//	{X: 752, Y: 348, Color: "a09e93"},
//	{X: 604, Y: 405, Color: "9f9d92"},
//	{X: 709, Y: 390, Color: "9f9d92"},
//}

//// 第一章地图1界面传送点
//var Chapter1_1TpPoint = image.Point{525, 372}
//
//// 第一章地图2界面找色
//var Chapter1_2ColorCmp = []StructColorCmp{
//	{X: 508, Y: 231, Color: "9e9b93"},
//	{X: 648, Y: 317, Color: "9b9a8e"},
//	{X: 503, Y: 420, Color: "9c9a90"},
//	{X: 716, Y: 296, Color: "9a9990"},
//}
//
//// 第一章地图2界面传送点
//var Chapter1_2TpPoint = image.Point{543, 294}
//
//var ChapterSelectDetect = []StructColorCmp{
//	{X: 103, Y: 702, Color: "101821"},
//	{X: 605, Y: 704, Color: "101821"},
//	{X: 1163, Y: 703, Color: "101821"},
//}
//
//var If_Map = []StructColorCmp{ //判断是否在跑图界面
//	{X: 321, Y: 269, Color: "ffffff"},
//	{X: 1096, Y: 562, Color: "dedad6"},
//	{X: 1091, Y: 54, Color: "ffffff"},
//	{X: 317, Y: 83, Color: "ffffff"},
//}
//
//var If_GenerateTP = []StructColorCmp{ //判断是否生成传送
//	{X: 477, Y: 442, Color: "dedad6"},
//	{X: 614, Y: 446, Color: "dedad6"},
//	{X: 660, Y: 443, Color: "dedad6"},
//	{X: 802, Y: 443, Color: "dedad6"},
//}
//
//// 判断是否在加载界面
//var LoadingInterface = []StructColorCmp{
//	{X: 191, Y: 149, Color: "000000"},
//	{X: 190, Y: 343, Color: "000000"},
//	{X: 299, Y: 400, Color: "000000"},
//	{X: 383, Y: 163, Color: "000000"},
//	{X: 63, Y: 42, Color: "000000"},
//}
//
//// 地图中选择地图的左按钮找色
//var TpLeftButten = []StructColorCmp{
//	{X: 300, Y: 344, Color: "dedad6"},
//	{X: 305, Y: 338, Color: "dedace"},
//	{X: 304, Y: 349, Color: "dddbd5"},
//}
//
//// 地图中选择地图的右按钮找色
//var TpRightButten = []StructColorCmp{
//	{X: 979, Y: 344, Color: "dedad6"},
//	{X: 974, Y: 338, Color: "dedace"},
//	{X: 975, Y: 349, Color: "dddbd5"},
//}
