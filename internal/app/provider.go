package app

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewApp,
	NewEnter,
	InitGin,
	InitEnv,
	InitCtx,
	NewAsynqEnter,
)
