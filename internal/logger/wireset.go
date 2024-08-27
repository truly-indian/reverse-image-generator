package logger

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewLogger,
)
