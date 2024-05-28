package task

import (
	"github.com/google/wire"

	"github.com/yituoshiniao/gin-api-http/internal/task/queue"
	"github.com/yituoshiniao/gin-api-http/internal/task/scheduler"
)

var WireSet = wire.NewSet(
	NewEnter,
	NewCronSingleAppleProductPriceTask,
	queue.NewEmailDeliveryTask,
	queue.NewNewGroupEmailDeliveryTask,
	scheduler.NewTestSched,
	scheduler.NewDynamicPeriodicTaskSched,
	queue.NewRegisterTask,
	scheduler.NewRegisterSched,
)
