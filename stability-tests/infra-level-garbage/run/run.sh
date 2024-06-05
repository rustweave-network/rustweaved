#!/bin/bash
rm -rf /tmp/rustweaved-temp

rustweaved --devnet --appdir=/tmp/rustweaved-temp --profile=6061 &
RUSTWEAVED_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $RUSTWEAVED_PID

wait $RUSTWEAVED_PID
RUSTWEAVED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Rustweaved exit code: $RUSTWEAVED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $RUSTWEAVED_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
