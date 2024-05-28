package auth

import (
	"github.com/google/wire"
	"github.com/yituoshiniao/gin-api-http/internal/module/auth/application/service"
)

var WireSet = wire.NewSet(
	service.NewJwtTokenSrv,
)
