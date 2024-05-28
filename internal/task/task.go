package task

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/internal/metrics"
)

type Task interface {
	Handle() error
	Name() string
}

type TaskCron struct {
	task           Task
	ctx            context.Context
	counterMetrics *metrics.CounterMetrics
}

func NewTaskCron(
	ctx context.Context,
	task Task,
	counterMetrics *metrics.CounterMetrics,
) *TaskCron {
	return &TaskCron{
		task:           task,
		ctx:            ctx,
		counterMetrics: counterMetrics,
	}
}

func (t *TaskCron) Run() {
	defer func() {
		err := recover()
		if err != nil {
			xlog.S(t.ctx).Errorw("taskCron有不可避免的异常", "err", errors.WithStack(fmt.Errorf("%v", err)))
			go func() {
				t.counterMetrics.TaskStatusCounter.With(metrics.TaskName, t.task.Name(), metrics.TaskType, "panic").Add(1)
			}()
		}
	}()

	go func() {
		t.counterMetrics.TaskStatusCounter.With(metrics.TaskName, t.task.Name(), metrics.TaskType, "count").Add(1)
	}()

	err := t.task.Handle()
	if err != nil {
		xlog.S(t.ctx).Errorw("taskCron执行handle发生异常", "err", err)
		go func() {
			t.counterMetrics.TaskStatusCounter.With(metrics.TaskName, t.task.Name(), metrics.TaskType, "fail").Add(1)
		}()
	} else {
		go func() {
			t.counterMetrics.TaskStatusCounter.With(metrics.TaskName, t.task.Name(), metrics.TaskType, "success").Add(1)
		}()
	}
}
