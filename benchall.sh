


BENCH_CMD="go test -bench=. -count=10 -timeout=100h -benchmem"

# env BENCH_TYPE=runtime go generate
# $BENCH_CMD > runtime.txt
# $BENCH_CMD -opt=beast > runtime-beast.txt
# env GOEXPERIMENT=swisstable $BENCH_CMD > runtime-swisstable.txt
# env GOEXPERIMENT=swisstable $BENCH_CMD -opt=beast > runtime-swisstable-beast.txt

# nohup go test -bench=".*/Int64/.*" -count=10 -timeout=100h -benchmem > a.txt & 

# env BENCH_TYPE=swiss0 go generate
# $BENCH_CMD > swiss0.txt
# env BENCH_TYPE=swiss1 go generate
# $BENCH_CMD > swiss1.txt
