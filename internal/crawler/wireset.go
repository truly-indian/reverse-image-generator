package crawler

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewCrawler,
)
