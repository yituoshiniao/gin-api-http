package asynqdemo

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/internal/api/dto"
	"github.com/yituoshiniao/gin-api-http/internal/api/http"
	"github.com/yituoshiniao/gin-api-http/internal/task/queue"
)

type TaskProducer struct {
	asynqClient       *asynq.Client
	response          *http.Response
	emailDeliveryTask *queue.EmailDeliveryTask
}

func NewTaskProducer(
	asynqClient *asynq.Client,
	response *http.Response,
	emailDeliveryTask *queue.EmailDeliveryTask,
) *TaskProducer {
	return &TaskProducer{
		asynqClient:       asynqClient,
		response:          response,
		emailDeliveryTask: emailDeliveryTask,
	}
}

// EmailDeliveryTaskAdd
//
//	@Summary		asynq-add异步任务
//	@Description	asynq-异步任务,可通过： http://localhost:7013/monitoring/ 查看dashbord报表
//	@Tags			asynq
//	@Param			object	query		dto.AsynqEmailDeliveryTaskAddRequest	1	"request"
//	@Success		200		{object}	dto.AsynqEmailDeliveryTaskAddResponse	"请求成功"
//	@Router			/asynq/v1/addTask [GET]
func (d *TaskProducer) EmailDeliveryTaskAdd(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.AsynqEmailDeliveryTaskAddRequest{}
	xlog.S(ctx).Infow("req-信息", "req", req)
	client := d.asynqClient

	// 初使货需要传递的数据
	task, err := d.emailDeliveryTask.Task(ctx, 42, fmt.Sprintf("some:template:id:%d", 1), `{"name":"李四"}`)
	if err != nil {
		xlog.S(ctx).Errorw("NewEmailDeliveryTask-错误", "err", err)
		d.response.Error(c, err.Error(), nil)
		return
	}
	// 任务入队 立即执行
	// info, err := client.Enqueue(task)

	// 任务入队 立即执行
	// info, err := client.Enqueue(task, time.Now())

	// 指定任务id, 相同的任务id会去重
	taskid := uuid.New().String()
	xlog.S(ctx).Infow("taskid-信息", "taskid", taskid)

	// 默认队列default
	// info, err := client.Enqueue(task, asynq.TaskID(taskid))

	// 延迟执行 3分钟后执行
	// info, err := client.Enqueue(task, asynq.ProcessIn(3*time.Minute))

	// 延迟执行 5分钟后执行
	// info, err := client.Enqueue(task, asynq.ProcessAt(time.Now().Add(5*time.Minute)))

	// MaxRetry 重度次数, Timeout超时时间
	// info, err = client.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(3*time.Second))

	// 指定任务保存时间;通过Retention函数指定一个时间段，代表当任务被处理完后，还可以保留一定的时间。如下是当task被成功消费后再保留2个小时。
	// info, err := client.Enqueue(task, asynq.Retention(10*time.Hour))

	// 使用Unique选项：让 Asynq 为任务创建唯一性锁; //使用此选项排队的任务保证在给定的ttl中是唯一的。在接下来的一个小时内持有唯一性锁。
	// info, err = client.Enqueue(task, asynq.Unique(time.Hour))

	// 指定入队的队列;通过Queue选项函数指定将消息发送至哪个队列。如下，将消息发送至redis的名为“high”的队列中。
	// info, err = client.Enqueue(task, asynq.Queue("high"))

	// 指定任务执行时间, 需要在10秒内完成的任务， 也可通过 asynq.Deadline(xmas)定义
	info, err := client.Enqueue(task, asynq.Timeout(10*time.Second))

	if err != nil {
		xlog.S(ctx).Errorw("Enqueue-错误", "err", err)
		d.response.Error(c, err.Error(), nil)
		return
	}
	d.response.Success(c, info)
}
