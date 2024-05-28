package mockv2

import (
	"github.com/google/wire"

	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2/application/repository"
	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2/application/service"
)

var WireSet = wire.NewSet(
	service.NewUserScoreSrv,
	repository.NewUserScoreRepo,
)
