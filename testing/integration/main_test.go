package integration

import (
	"github.com/rustweave-network/rustweaved/infrastructure/logger"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	logger.SetLogLevels(logger.LevelDebug)
	logger.InitLogStdout(logger.LevelDebug)

	os.Exit(m.Run())
}
