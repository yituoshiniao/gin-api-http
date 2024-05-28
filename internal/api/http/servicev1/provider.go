package servicev1

import (
	"github.com/google/wire"
)

var (
	WireSet = wire.NewSet(
		NewUserScoreSrv,
	)
)
