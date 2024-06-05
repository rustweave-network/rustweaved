package handshake

import (
	"github.com/rustweave-network/rustweaved/infrastructure/logger"
	"github.com/rustweave-network/rustweaved/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
