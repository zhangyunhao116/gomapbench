


BENCH_CMD="go test -bench=. -count=10 -timeout=10h -benchmem"

env BENCH_TYPE=runtime go generate
$BENCH_CMD > runtime.txt
env GOEXPERIMENT=swisstable $BENCH_CMD > runtime-swisstable.txt

env BENCH_TYPE=swiss0 go generate
$BENCH_CMD > swiss0.txt
env BENCH_TYPE=swiss1 go generate
$BENCH_CMD > swiss1.txt