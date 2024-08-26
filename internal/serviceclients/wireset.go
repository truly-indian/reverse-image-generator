package serviceclients

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewSerpAPIClient,
)
