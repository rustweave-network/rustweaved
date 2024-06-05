#!/bin/bash
rm -rf /tmp/rustweaved-temp

NUM_CLIENTS=128
rustweaved --devnet --appdir=/tmp/rustweaved-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
RUSTWEAVED_PID=$!
RUSTWEAVED_KILLED=0
function killRustweavedIfNotKilled() {
  if [ $RUSTWEAVED_KILLED -eq 0 ]; then
    kill $RUSTWEAVED_PID
  fi
}
trap "killRustweavedIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $RUSTWEAVED_PID

wait $RUSTWEAVED_PID
RUSTWEAVED_EXIT_CODE=$?
RUSTWEAVED_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Rustweaved exit code: $RUSTWEAVED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $RUSTWEAVED_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
