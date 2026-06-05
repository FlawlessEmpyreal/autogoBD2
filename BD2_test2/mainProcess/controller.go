package mainProcess

import (
	"fmt"
	"log"
	"time"
)

// 单个阶段
type Stage struct {
	Name string
	Run  func() (RecoveryAction, error)
}

// 章节配置
type ChapterConfig struct {
	Name       string
	Enabled    bool                             // 章节是否可执行
	MaxRetries int                              // 单个阶段最大重试次数
	Stages     []Stage                          // 阶段列表，按顺序执行
	OnRecover  func(*StageError) RecoveryAction // 出错时的恢复策略
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
		cfg.OnRecover = func(e *StageError) RecoveryAction {
			return RetryStage
		}
	}
	c.chapters = append(c.chapters, cfg)
}

// 执行单个章节（按阶段遍历 + 出错恢复）
func (c *Controller) runChapter(cfg *ChapterConfig) error {
	log.Printf("━━━ 开始章节: %s ━━━", cfg.Name)
	retries := 0

	for i := 0; i < len(cfg.Stages); {
		stage := cfg.Stages[i]
		log.Printf("  ▶ 阶段 [%d/%d] %s", i+1, len(cfg.Stages), stage.Name)

		action, err := stage.Run()

		if err != nil {
			// 包装错误
			stageErr := &StageError{
				Chapter: cfg.Name,
				Stage:   stage.Name,
				Retries: retries,
				Action:  action,
				Err:     err,
			}
			log.Printf("  ❌ 错误: %v", stageErr)

			switch stageErr.Action {
			case RetryStage:
				retries++
				if retries > cfg.MaxRetries {
					log.Printf("  ⛔ %s 超过最大重试次数(%d)，升级为 goToBattleInterface", stage.Name, cfg.MaxRetries)
					goToBattleInterface()
					retries = 0
					// i 不变，主菜单处理完后重试当前阶段
				} else {
					log.Printf("  🔄 重试阶段 %s (%d/%d)", stage.Name, retries, cfg.MaxRetries)
					time.Sleep(2 * time.Second)
					// i 不变
				}

			case RetryChapter:
				log.Printf("  🔄 从头重跑章节 %s", cfg.Name)
				retries = 0
				i = 0 // 回到第一个阶段

			case GoBattleInterface:
				log.Printf("  🏠 回跑图界面...")
				goToBattleInterface() // 阻塞，直到返回跑图界面
				log.Printf("  ↩ 返回跑图界面完毕，重试阶段 %s", stage.Name)
				retries = 0
				// i 不变，重试当前阶段

			case SkipStage:
				log.Printf("  ⏭ 跳过阶段 %s", stage.Name)
				retries = 0
				i++ // 跳到下一阶段

			case SkipChapter:
				log.Printf("  ⏭ 跳过整个章节 %s", cfg.Name)
				return nil

			case AbortAll:
				log.Printf("  🛑 终止所有执行")
				return fmt.Errorf("abort: %v", stageErr)
			}
		} else {
			log.Printf("  ✅ %s 完成", stage.Name)
			retries = 0 // 进入下一阶段时重置重试计数
			i++         // 前进到下一阶段
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

func goToBattleInterface() {

}
