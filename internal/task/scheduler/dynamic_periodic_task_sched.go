package scheduler

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtrace"
)

type DynamicPeriodicTaskSched struct {
}

func NewDynamicPeriodicTaskSched() *DynamicPeriodicTaskSched {
	return &DynamicPeriodicTaskSched{}

}

// TypeName 调度任务type名
func (t *DynamicPeriodicTaskSched) TypeName() (taskName string) {
	return "dynamicPeriodicTask"
}

// Handle 调度任务处理
func (e *DynamicPeriodicTaskSched) Handle(ctx context.Context, t *asynq.Task) error {
	// 接收任务数据

	payload := string(t.Payload())

	// 逻辑处理start...
	xlog.S(ctx).Infow("DynamicPeriodicTaskSched-调度任务处理-信息", "payload", payload, " t.Type", t.Type())

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
func (e *DynamicPeriodicTaskSched) Do(ctx context.Context, t *asynq.Task) (res []byte, err error) {
	traceID := xtrace.GetTraceID(ctx)
	response := struct {
		TractID string
	}{
		TractID: traceID,
	}
	res, err = json.Marshal(response)
	return
}
