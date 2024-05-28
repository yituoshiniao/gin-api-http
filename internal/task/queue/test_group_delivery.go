package queue

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtask"
	"github.com/yituoshiniao/kit/xtrace"
)

const (
	/*
		GroupTypeEmailDelivery task typename,
		注意每一个group的 都需要一个aggregate聚合函数处理，所以每一个聚合任务都需要全局初始化一个,否则就使用全局初始化的
	*/
	GroupTypeEmailDelivery = xtask.AggregateTypeName
)

// GroupEmailDeliveryPayload 异步任务需要传递的数据结构
type GroupEmailDeliveryPayload struct {
	UserID     int
	TemplateID string
	DataStr    string
	TraceID    string
}

type GroupEmailDeliveryTask struct{}

func NewNewGroupEmailDeliveryTask() *GroupEmailDeliveryTask {
	return &GroupEmailDeliveryTask{}
}

// Task 异步任务需要传递的数据
func (e *GroupEmailDeliveryTask) Task(ctx context.Context, userID int, tmplID, dataStr string) (*asynq.Task, error) {
	tmpPayload := GroupEmailDeliveryPayload{
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

	xlog.S(ctx).Infow("groupTask-任务信息", "tmpPayload", tmpPayload)

	return asynq.NewTask(GroupTypeEmailDelivery, payload), nil
}

func (e *GroupEmailDeliveryTask) AggregatedTask(ctx context.Context, task *asynq.Task) error {

	xlog.S(ctx).Infow("处理程序收到聚合的任务", "task.payload", string(task.Payload()), "task.payload", task.Type())
	return nil
}
