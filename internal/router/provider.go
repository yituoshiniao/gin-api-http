package router

import (
	"github.com/google/wire"

	v1 "github.com/yituoshiniao/gin-api-http/internal/router/v1"
)

var WireSet = wire.NewSet(
	NewRouter,
	v1.NewGoodsCenterRouter,
	v1.NewGoodsCenterTestV2,
	// v1.NewIapRouter,
)
