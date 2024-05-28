package util

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewBigCache,
	NewSnowflakeIDClient,
	NewGob,
)
