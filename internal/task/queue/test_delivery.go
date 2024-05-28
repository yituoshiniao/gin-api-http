package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtrace"
)

const (
	// TypeEmailDelivery task typename
	TypeEmailDelivery = "email:deliver"
)

// EmailDeliveryPayload 异步任务需要传递的数据结构
type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
	DataStr    string
	TraceID    string
}

type EmailDeliveryTask struct{}

func NewEmailDeliveryTask() *EmailDeliveryTask {
	return &EmailDeliveryTask{}
}

// Task 异步任务需要传递的数据
func (e *EmailDeliveryTask) Task(ctx context.Context, userID int, tmplID, dataStr string) (*asynq.Task, error) {
	tmpPayload := EmailDeliveryPayload{
		UserID:     userID,
		TemplateID: tmplID,
		DataStr:    dataStr,
		TraceID:    xtrace.GetTraceID(ctx),
	}
	payload, err := json.Marshal(tmpPayload)
	if err != nil {
		xlog.S(ctx).Infow("Marshal-错误", "err", err)
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

// Handle 发送email处理逻辑
func (e *EmailDeliveryTask) Handle(ctx context.Context, t *asynq.Task) error {
	// 接收任务数据
	var payload EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		xlog.S(context.Background()).Infow("Unmarshal-错误", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// 逻辑处理start...
	xlog.S(ctx).Infow("payload-信息", "payload", payload)

	// asynq.ErrDuplicateTask
	// 返回错误, 那么将启动重试
	// return  errors.New("错误")

	// 当配置了Retention时，可以将处理结果也暂存在redis中，方便回溯查看
	res, err := e.DoStuff(ctx, t)
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

// DoStuff 处理task
func (e *EmailDeliveryTask) DoStuff(ctx context.Context, t *asynq.Task) (res []byte, err error) {
	traceID := xtrace.GetTraceID(ctx)
	response := struct {
		TractID string
	}{
		TractID: traceID,
	}
	res, err = json.Marshal(response)
	return
}

// CancelHandle 如果 Handler.ProcessTask 返回一个 SkipRetry 错误，不管剩余的重试次数是多少，任务都将被归档。
func (e *EmailDeliveryTask) CancelHandle(ctx context.Context, t *asynq.Task) error {
	// 接收任务数据
	var payload EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		xlog.S(context.Background()).Infow("Unmarshal-错误", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// 逻辑处理start...
	xlog.S(ctx).Infow("payload-信息", "payload", payload)

	// 如果 Handler.ProcessTask 返回一个 SkipRetry 错误，不管剩余的重试次数是多少，任务都将被归档。
	// err := asynq.SkipRetry

	// 任务超时&撤销  任务设置了超时时间或截止时间，如何处理取消操作。
	c := make(chan error, 1)
	go func() {
		c <- e.doWork(ctx, t)
	}()
	select {
	case <-ctx.Done():
		xlog.S(ctx).Warnw("ctx.Done", "err", ctx.Err())
		// cancelation signal received, abandon this work.
		return ctx.Err()
	case res := <-c:
		return res
	}

}

func (e *EmailDeliveryTask) doWork(ctx context.Context, t *asynq.Task) error {
	// 接收任务数据
	var payload EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		xlog.S(context.Background()).Infow("Unmarshal-错误", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// 逻辑处理start...
	xlog.S(ctx).Infow("payload-信息", "payload", payload)

	// 模拟任务超时
	time.Sleep(10 * time.Second)

	return nil
}
