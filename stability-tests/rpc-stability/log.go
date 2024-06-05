package main

import (
	"github.com/rustweave-network/rustweaved/infrastructure/logger"
	"github.com/rustweave-network/rustweaved/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
