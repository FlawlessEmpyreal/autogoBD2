package mainProcess

import (
	"app/MyOpenCV"
	"app/info"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Dasongzi1366/AutoGo/motion"
	"github.com/Dasongzi1366/AutoGo/opencv"
)

// 单个阶段
type Stage struct {
	type_ string
	Name  string
	Run   func() (info.RecoveryAction, error)
}

// 章节配置
type ChapterConfig struct {
	Name       string
	Enabled    bool                                  // 章节是否可执行
	MaxRetries int                                   // 单个阶段最大重试次数
	Stages     []Stage                               // 阶段列表，按顺序执行
	OnRecover  func(*StageError) info.RecoveryAction // 出错时的恢复策略
}

// 中控
type Controller struct {
	chapters []*ChapterConfig
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Register(cfg *ChapterConfig) {
	if cfg.MaxRetries == 0 {
		cfg.MaxRetries = 3
	}
	if cfg.OnRecover == nil {
		// 默认策略：重试当前阶段
		cfg.OnRecover = func(e *StageError) info.RecoveryAction {
			return info.RetryStage
		}
	}
	c.chapters = append(c.chapters, cfg)
}

// 执行单个章节（按阶段遍历 + 出错恢复）
func (c *Controller) runChapter(cfg *ChapterConfig) error {
	log.Printf("━━━ 开始章节: %s ━━━", cfg.Name)
	retries := 0
	GoToRunMap := 0
	RetryChapter := 0
	for i := 0; i < len(cfg.Stages); {

		stage := cfg.Stages[i]
		log.Printf("  ▶ 阶段 [%d/%d] %s", i+1, len(cfg.Stages), stage.Name)

		action, err := stage.Run()

		if err != nil {
			// 包装错误
			if stage.type_ != "FindChapter" && GoToRunMap > 2 { //找章节不能跳过
				if stage.type_ == "Tp" {
					i++
				} //tp报错太多这张图也不用跑了
				action = info.SkipStage
			} else if stage.type_ == "FindChapter" && GoToRunMap > 2 { //找章节老失败直接跳过本章
				log.Printf("  ⏭ 寻找章节失败三次跳过本章节")
				action = info.SkipChapter
			}
			if RetryChapter > 2 {
				action = info.SkipChapter
			}
			stageErr := &StageError{
				Chapter: cfg.Name,
				Stage:   stage.Name,
				Retries: retries,
				Action:  action,
				Err:     err,
			}
			log.Printf("  ❌ 错误: %v", stageErr)

			switch stageErr.Action {
			case info.RetryStage:
				retries++
				if retries > cfg.MaxRetries {
					log.Printf("  ⛔ %s 超过最大重试次数(%d)，升级为 goToRunMapInterface", stage.Name, cfg.MaxRetries)
					GoToRunMapInterface()
					GoToRunMap++
					retries = 0
					// i 不变，主菜单处理完后重试当前阶段
				} else {
					log.Printf("  🔄 重试阶段 %s (%d/%d)", stage.Name, retries, cfg.MaxRetries)
					time.Sleep(2 * time.Second)
					// i 不变
				}

			case info.RetryChapter:
				log.Printf("  🔄 从头重跑章节 %s", cfg.Name)
				GoToRunMap = 0
				RetryChapter++
				retries = 0
				i = 0 // 回到第一个阶段

			case info.GoToRunMapInterface:
				log.Printf("  🏠 回跑图界面...")
				GoToRunMapInterface() // 阻塞，直到返回跑图界面
				log.Printf("已返回跑图界面次数:%d", GoToRunMap)
				GoToRunMap++
				log.Printf("  ↩ 返回跑图界面完毕，重试阶段 %s", stage.Name)
				time.Sleep(500 * time.Millisecond)
				motion.Click(645, 384, 0, 0) //点一下防止持续跑步
				time.Sleep(200 * time.Millisecond)
				motion.Click(645, 384, 0, 0) //点一下防止持续跑步
				time.Sleep(500 * time.Millisecond)
				retries = 0
				// i 不变，重试当前阶段

			case info.SkipStage:
				log.Printf("  ⏭ 跳过阶段 %s", stage.Name)
				GoToRunMap = 0 //重置单个阶段回跑图界面的次数
				retries = 0
				i++ // 跳到下一阶段

			case info.SkipChapter:
				log.Printf("  ⏭ 章节重复次数太多,跳过整个章节 %s", cfg.Name)
				return nil

			case info.AbortAll:
				log.Printf("  🛑 终止所有执行")
				return fmt.Errorf("abort: %v", stageErr)
			}
		} else {
			log.Printf("  ✅ %s 完成", stage.Name)
			retries = 0 // 进入下一阶段时重置重试计数
			GoToRunMap = 0
			RetryChapter = 0
			i++ // 前进到下一阶段
			continue
		}
	}
	log.Printf("━━━ 章节完成: %s ━━━", cfg.Name)
	return nil
}

// 执行所有启用的章节
func (c *Controller) RunAll() error {
	for _, cfg := range c.chapters {
		if !cfg.Enabled {
			log.Printf("⏸ 跳过(disabled): %s", cfg.Name)
			continue
		}
		if err := c.runChapter(cfg); err != nil {
			return err
		}
	}
	return nil
}

// 只执行指定名称的章节（按注册顺序）
func (c *Controller) RunByName(names ...string) error {
	nameSet := make(map[string]bool)
	for _, n := range names {
		nameSet[n] = true
	}
	for _, cfg := range c.chapters {
		if nameSet[cfg.Name] {
			if err := c.runChapter(cfg); err != nil {
				return err
			}
		}
	}
	return nil
}

func GoToRunMapInterface() {
	imgCS := MyOpenCV.ChapterSelectInterface()
	x, y := opencv.FindImage(37, 670, 1202, 720, imgCS, false, false, 0.5, 0)
	for i := 0; i < 3; i++ {
		if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //先判断是不是已经在跑图界面
			return
		}
		time.Sleep(500 * time.Millisecond)

		if MyOpenCV.If_BattleInterface(0.85) { //进战斗退出
			for i := 0; i < 10; i++ {
				MyOpenCV.WaitLoading(10)                              //等待十秒加载
				if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
					return
				}

				motion.Click(1178, 42, 0, 0)
				if MyOpenCV.ColorCmp(info.IF.If_Pause.ColorsCmp, 0.7) {
					time.Sleep(1 * time.Second)
					motion.Click(510, 660, 0, 0)
					time.Sleep(1 * time.Second)
					motion.Click(721, 433, 0, 0)
				}

				time.Sleep(1000 * time.Millisecond)
			}
		}

		if MyOpenCV.ColorCmp(info.IF.If_Pause.ColorsCmp, 0.8) { //如果在暂停页面
			time.Sleep(1 * time.Second)
			motion.Click(510, 660, 0, 0)
			time.Sleep(1 * time.Second)
			motion.Click(721, 433, 0, 0)
			time.Sleep(1500 * time.Millisecond)
			MyOpenCV.If_LoadingInterface(10)
			if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
				return
			}
		}

		if MyOpenCV.If_escape() { //如果在逃跑页面
			motion.Click(721, 433, 0, 0)
			time.Sleep(1500 * time.Millisecond)
			MyOpenCV.If_LoadingInterface(10)
			if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
				return
			}
		}

		//返回按钮找色
		if MyOpenCV.ListColorsCmp(info.IF.If_Backbutton[:], 0.7) { //如果有返回按钮
			motion.Click(info.MiscPoint.BackButton.X, info.MiscPoint.BackButton.Y, 0, 0)
			time.Sleep(1500 * time.Millisecond)
			MyOpenCV.If_LoadingInterface(10)
			if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
				return
			}
		}

		//返回按钮找图
		if MyOpenCV.If_BackButton() {
			motion.Click(info.MiscPoint.BackButton.X, info.MiscPoint.BackButton.Y, 0, 0)
			time.Sleep(1500 * time.Millisecond)
			MyOpenCV.If_LoadingInterface(10)
			if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
				return
			}
		}

		//点战场角色界面按钮,这里其实是退出战场角色界面
		time.Sleep(500 * time.Millisecond)
		if MyOpenCV.ColorCmp(info.IF.If_BattleFieldRole.ColorsCmp, 0.8) && x == -1 && y == -1 {
			motion.Click(644, 652, 0, 0)
			time.Sleep(1500 * time.Millisecond)
			MyOpenCV.If_LoadingInterface(10)
			if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
				return
			}
		}

		time.Sleep(500 * time.Millisecond)
		//如果还不在那可能是打开钻石金币弹窗了点一下战场界面按钮
		if x == -1 && y == -1 {
			motion.Click(644, 652, 0, 0) //点战场角色界面按钮,关闭金币弹窗
			time.Sleep(1500 * time.Millisecond)
			MyOpenCV.If_LoadingInterface(10)
			if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
				return
			}
		}
	}

	//如果还不行大概率是背景太白了,点一下返回键试试
	motion.Click(info.MiscPoint.BackButton.X, info.MiscPoint.BackButton.Y, 0, 0)
	time.Sleep(1500 * time.Millisecond)
	MyOpenCV.If_LoadingInterface(10)
	if MyOpenCV.ColorCmp(info.IF.If_Map.ColorsCmp, 0.8) { //判断是不是已经在跑图界面
		return
	}

	println("无法返回跑图界面")
	os.Exit(0)
}
