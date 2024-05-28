package scheduler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtask"
	"github.com/yituoshiniao/kit/xtrace"
)

type TestSched struct {
	scheduler *asynq.Scheduler
}

func NewTestSched(
	scheduler *asynq.Scheduler,
) *TestSched {
	return &TestSched{
		scheduler: scheduler,
	}

}

// TypeName 调度任务名
func (t *TestSched) TypeName() (taskName string) {
	return "test_sched_example_task"
}

// QueueName 队列名
func (t *TestSched) QueueName() (taskName string) {
	return string(xtask.GroupSchedulerQueue)
}

// Cronspec 定时任务时间分布
func (t *TestSched) Cronspec() (cronspec string) {
	// 你还可以使用"@every "来指定间隔。
	spec := "@every 30s" // 每隔30秒

	// spec := "14 11 * * ?" //每三十分钟执行一次
	// spec := "*/1 * * * ?" //每三十分钟执行一次
	// spec := "10 3,5 * * ?" //每日凌晨3点 和5点10分执行

	return spec

}

func (t *TestSched) Register() (entryID string, err error) {
	ctx := xtrace.NewCtxWithTraceId(context.Background())
	// 添加task任务，可以添加payload 处理数据，如果没有可以不添加
	task := asynq.NewTask(
		t.TypeName(),
		nil,

		// 当使用调度Scheduler时,task的option属性，尽可能放在 scheduler.Register 中处理
		// //asynq.Unique(time.Hour*10),
		// asynq.MaxRetry(2),
		// asynq.Retention(time.Hour*4),
		// asynq.Queue(string(xtask.HighQueue)),
		// //asynq.TaskID(t.TaskName()),
	)

	// 指定调度时间
	entryID, err = t.scheduler.Register(
		t.Cronspec(),
		task,
		asynq.Retention(time.Hour*10),
		asynq.Queue(t.QueueName()),
		// asynq.Unique(time.Hour*10),
		asynq.MaxRetry(3),
		// asynq.TaskID(t.TaskName()),
	)
	if err != nil {
		xlog.S(ctx).Errorw("TestSched-Register-错误", "err", err)
	}
	xlog.S(ctx).Infow("TestSched-Register", "err", err, "entryID", entryID)
	return entryID, err
}

// Handle 调度任务处理
func (e *TestSched) Handle(ctx context.Context, t *asynq.Task) error {
	// 接收任务数据

	payload := string(t.Payload())

	// 逻辑处理start...
	xlog.S(ctx).Infow("调度任务处理-信息", "payload", payload, " t.Type", t.Type())

	// 当配置了Retention时，可以将处理结果也暂存在redis中，方便回溯查看
	res, err := e.Do(ctx, t)
	if err != nil {
		xlog.S(ctx).Errorw("DoStuff-错误", "payload", payload)
		return err
	}
	if _, err = t.ResultWriter().Write(res); err != nil {
		xlog.S(ctx).Errorw("ResultWriter-错误", "payload", payload)
		return err
	}

	return nil
}

// Do 处理task
func (s *TestSched) Do(ctx context.Context, t *asynq.Task) (res []byte, err error) {
	traceID := xtrace.GetTraceID(ctx)
	response := struct {
		TractID string
	}{
		TractID: traceID,
	}
	res, err = json.Marshal(response)
	return
}
