package cron

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/yituoshiniao/gin-api-http/internal/metrics"
)

// Enter 手动命令
type Enter struct {
	counterMetrics     *metrics.CounterMetrics
	factoryTemplateSrv *FactoryTemplateSrv
}

func NewEnter(
	counterMetrics *metrics.CounterMetrics,
	factoryTemplateSrv *FactoryTemplateSrv,
) *Enter {
	return &Enter{
		counterMetrics:     counterMetrics,
		factoryTemplateSrv: factoryTemplateSrv,
	}
}

func (e *Enter) RunCommandTask(ctx context.Context, taskName string) error {
	e.factoryTemplateSrv.RegisterTemplate(ctx)
	if _, ok := Template[taskName]; !ok {
		return errors.New(fmt.Sprintf("task 不存在: %s", taskName))
	}
	return Template[taskName].Handle()
}
