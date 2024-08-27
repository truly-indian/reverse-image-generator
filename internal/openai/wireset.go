package openai

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewOpenAIClient,
)
