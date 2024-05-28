package cron

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewTestSrv,
	NewFactoryTemplateSrv,
)
