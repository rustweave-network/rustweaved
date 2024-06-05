#!/bin/bash
rm -rf /tmp/rustweaved-temp

rustweaved --devnet --appdir=/tmp/rustweaved-temp --profile=6061 --loglevel=debug &
RUSTWEAVED_PID=$!
RUSTWEAVED_KILLED=0
function killRustweavedIfNotKilled() {
    if [ $RUSTWEAVED_KILLED -eq 0 ]; then
      kill $RUSTWEAVED_PID
    fi
}
trap "killRustweavedIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $RUSTWEAVED_PID

wait $RUSTWEAVED_PID
RUSTWEAVED_KILLED=1
RUSTWEAVED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Rustweaved exit code: $RUSTWEAVED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $RUSTWEAVED_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
