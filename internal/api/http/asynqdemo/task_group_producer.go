package asynqdemo

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtask"

	"github.com/yituoshiniao/gin-api-http/internal/api/dto"
	"github.com/yituoshiniao/gin-api-http/internal/api/http"
	"github.com/yituoshiniao/gin-api-http/internal/task/queue"
)

type TaskGroupProducer struct {
	asynqClient            *asynq.Client
	response               *http.Response
	emailDeliveryTask      *queue.EmailDeliveryTask
	groupEmailDeliveryTask *queue.GroupEmailDeliveryTask
}

func NewTaskGroupProducer(
	asynqClient *asynq.Client,
	response *http.Response,
	emailDeliveryTask *queue.EmailDeliveryTask,
	groupEmailDeliveryTask *queue.GroupEmailDeliveryTask,
) *TaskGroupProducer {
	return &TaskGroupProducer{
		asynqClient:            asynqClient,
		response:               response,
		emailDeliveryTask:      emailDeliveryTask,
		groupEmailDeliveryTask: groupEmailDeliveryTask,
	}
}

// GroupDeliveryTaskAdd
//
//	@Summary		asynq-添加聚合任务
//	@Description	asynq-添加聚合任务
//	@Tags			asynq
//	@Param			object	query		dto.GroupDeliveryTaskAddRequest		1	"request"
//	@Success		200		{object}	dto.GroupDeliveryTaskAddResponse	"请求成功"
//	@Router			/asynq/v1/addAggTask [GET]
func (d *TaskGroupProducer) GroupDeliveryTaskAdd(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.GroupDeliveryTaskAddRequest{}
	xlog.S(ctx).Infow("req-信息", "req", req)

	client := d.asynqClient

	// 初使货需要传递的数据
	task, err := d.groupEmailDeliveryTask.Task(ctx, 42, fmt.Sprintf("some:template:id:%d", 1), `{"name":"李四1111"}`)

	// task := asynq.NewTask("aggregation-tutorial", []byte("1111"))

	if err != nil {
		xlog.S(ctx).Errorw("GroupDeliveryTaskAdd-错误", "err", err)
		d.response.Error(c, err.Error(), nil)
		return
	}

	// 指定任务id
	taskid := uuid.New().String()
	xlog.S(ctx).Infow("taskid-信息", "taskid", taskid)

	// 任务组
	groupName := "example-group-test" // 任务组名

	// 任务队列名, 注意 如果使用聚合组这里的 队列名必须使用Queues配置中的 队列名;
	queueName := string(xtask.GroupQueue)
	info, err := client.Enqueue(task, asynq.Queue(queueName), asynq.Group(groupName), asynq.TaskID(taskid))
	// info, err := client.Enqueue(task, asynq.Queue(queueName), asynq.Group(groupName))

	if err != nil {
		xlog.S(ctx).Errorw("Enqueue-错误", "err", err)
		d.response.Error(c, err.Error(), nil)
		return
	}
	d.response.Success(c, info)
}

// 任务聚合, 任务组
func (t *TaskGroupProducer) asynqGroup(ctx context.Context, client *asynq.Client, task *asynq.Task) (info *asynq.TaskInfo, err error) {
	groupName := "example-group-test" // 任务组名
	queueName := "groupQueue"         // 任务队列名
	info, err = client.Enqueue(task, asynq.Queue(queueName), asynq.Group(groupName))
	return
}
