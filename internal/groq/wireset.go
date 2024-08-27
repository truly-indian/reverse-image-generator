package groq

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewGroq,
)
