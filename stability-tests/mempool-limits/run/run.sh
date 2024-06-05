#!/bin/bash

APPDIR=/tmp/rustweaved-temp
RUSTWEAVED_RPC_PORT=29587

rm -rf "${APPDIR}"

rustweaved --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${RUSTWEAVED_RPC_PORT}" --profile=6061 &
RUSTWEAVED_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${RUSTWEAVED_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $RUSTWEAVED_PID

wait $RUSTWEAVED_PID
RUSTWEAVED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Rustweaved exit code: $RUSTWEAVED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $RUSTWEAVED_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
