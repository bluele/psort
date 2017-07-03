# psort

## Install

```
$ go get github.com/bluele/psort
```

## Benchmark

```
$ go test -bench . -benchmem
BenchmarkSort-4             1000           2714116 ns/op              64 B/op          2 allocs/op
BenchmarkPSort-4            5000            302611 ns/op              64 B/op          2 allocs/op
PASS
ok      github.com/bluele/psort 19.015s
```

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>
