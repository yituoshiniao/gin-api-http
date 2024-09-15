package dto

import (
	"github.com/yituoshiniao/gin-api-http/internal/handler"
)

type GroupDeliveryTaskAddRequest struct {
	// 环境变量,默认线上; sandbox 沙箱环境, production 生产环境
	Env string `form:"env,default=production"`
}

type GroupDeliveryTaskAddResponse struct {
	handler.ResponseData
}

type AsynqEmailDeliveryTaskAddRequest struct {
	// 环境变量,默认线上; sandbox 沙箱环境, production 生产环境
	Env string `form:"env,default=production"`
}

type AsynqEmailDeliveryTaskAddResponse struct {
	handler.ResponseData
}
