package common

import (
	"fmt"
	"github.com/rustweave-network/rustweaved/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunRustweavedForTesting runs rustweaved for testing purposes
func RunRustweavedForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	rustweavedRunCommand, err := StartCmd("KASPAD",
		"rustweaved",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("Rustweaved started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := rustweavedRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("Rustweaved closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := rustweavedRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("Rustweaved stopped")
	}
}
