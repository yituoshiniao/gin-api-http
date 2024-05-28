package asynqdemo

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(

	NewTaskProducer,
	NewTaskGroupProducer,
)
