package conn

import (
	"github.com/dghubble/sling"
	"github.com/opentracing/opentracing-go"
	"github.com/yituoshiniao/kit/xhttp/hclient"
	"go.uber.org/zap"

	"github.com/yituoshiniao/gin-api-http/config"
)

type IDClient struct {
	sling *sling.Sling
}

func NewIDClient(conf config.Config, logger *zap.Logger, tracer opentracing.Tracer) *IDClient {
	return &IDClient{
		sling: hclient.New(
			hclient.WithLogger(logger),
			hclient.WithTracer(tracer),
			hclient.WithTarget(conf.IDService.Host),
			// hclient.WithTimeout( time.Duration( conf.IDService.Timeout * int64(time.Second) )),
			hclient.WithInsecure(), hclient.WithServiceName("IDService"),
		),
	}
}
