package queue

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/internal/task/scheduler"
)

type RegisterTask struct {
	mux                      *asynq.ServeMux
	emailDeliveryTask        *EmailDeliveryTask
	groupEmailDeliveryTask   *GroupEmailDeliveryTask
	testSched                *scheduler.TestSched
	dynamicPeriodicTaskSched *scheduler.DynamicPeriodicTaskSched
}

func NewRegisterTask(
	mux *asynq.ServeMux,
	emailDeliveryTask *EmailDeliveryTask,
	groupEmailDeliveryTask *GroupEmailDeliveryTask,
	testSched *scheduler.TestSched,
	dynamicPeriodicTaskSched *scheduler.DynamicPeriodicTaskSched,
) *RegisterTask {
	return &RegisterTask{
		mux:                      mux,
		emailDeliveryTask:        emailDeliveryTask,
		groupEmailDeliveryTask:   groupEmailDeliveryTask,
		testSched:                testSched,
		dynamicPeriodicTaskSched: dynamicPeriodicTaskSched,
	}
}

// Register 注册任务
func (r *RegisterTask) Register() {
	// r.mux.HandleFunc(TypeEmailDelivery, loggingMiddleware(r.emailDeliveryTask.Handle)) //中间件使用

	r.mux.HandleFunc(TypeEmailDelivery, r.emailDeliveryTask.Handle)

	// 任务超时&撤销
	// r.mux.HandleFunc(TypeEmailDelivery, r.emailDeliveryTask.CancelHandle)

	// 组聚合任务
	// r.mux.HandleFunc(xtask.AggregateTypeName, r.groupEmailDeliveryTask.AggregatedTask)
	r.mux.HandleFunc(GroupTypeEmailDelivery, r.groupEmailDeliveryTask.AggregatedTask)

	// 调度定时任务
	r.mux.HandleFunc(r.testSched.TypeName(), r.testSched.Handle)

	// 动态调度任务
	r.mux.HandleFunc(r.dynamicPeriodicTaskSched.TypeName(), r.dynamicPeriodicTaskSched.Handle)
}

// loggingMiddleware 记录任务日志中间件
func loggingMiddleware(h asynq.HandlerFunc) asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		start := time.Now()
		xlog.S(ctx).Infow("Start processing", "task.payload", string(task.Payload()), "task", "task.Type", task.Type())
		err := h.ProcessTask(ctx, task)
		if err != nil {
			return err
		}
		xlog.S(ctx).Infow("Finished processing", "task.payload", string(task.Payload()), "task耗时", time.Since(start).String())
		return nil
	}
}
