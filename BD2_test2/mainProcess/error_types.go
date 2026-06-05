package mainProcess

//错误类型定义

import "fmt"

type RecoveryAction int

const (
	StageDone         RecoveryAction = iota //阶段成功
	RetryStage                              // 重试当前阶段
	RetryChapter                            // 从第一个阶段重来
	GoBattleInterface                       // 回战斗界面，处理完后重试当前阶段
	SkipStage                               // 跳过当前阶段，继续下一阶段
	SkipChapter                             // 跳过整个章节
	AbortAll                                // 终止所有
)

type StageError struct {
	Chapter string
	Stage   string
	Retries int //已重试次数
	Action  RecoveryAction
	Err     error
}

func (e *StageError) Error() string {
	return fmt.Sprintf("[%s → %s] (已重试%d次) %v", e.Chapter, e.Stage, e.Retries, e.Err)
}
