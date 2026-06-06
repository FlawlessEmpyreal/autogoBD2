package mainProcess

//错误类型定义

import (
	"app/info"
	"fmt"
)

type StageError struct {
	Chapter string
	Stage   string
	Retries int //已重试次数
	Action  info.RecoveryAction
	Err     error
}

func (e *StageError) Error() string {
	return fmt.Sprintf("[%s → %s] (已重试%d次) %v", e.Chapter, e.Stage, e.Retries, e.Err)
}
