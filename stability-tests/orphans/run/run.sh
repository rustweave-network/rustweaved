#!/bin/bash
rm -rf /tmp/rustweaved-temp

rustweaved --simnet --appdir=/tmp/rustweaved-temp --profile=6061 &
RUSTWEAVED_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $RUSTWEAVED_PID

wait $RUSTWEAVED_PID
RUSTWEAVED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Rustweaved exit code: $RUSTWEAVED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $RUSTWEAVED_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
