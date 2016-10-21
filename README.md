This is a Go response for https://medium.com/@rymcol/benchmarks-for-the-top-server-side-swift-frameworks-vs-node-js-24460cfe0beb

Currently only `/json` handler is implemented. `/blog` implementation is left as an exercise for readers :)
Go version used - `go tip` aka `go 1.8`.


# Build Time Results

```
time go build

real	0m0.466s
user	0m0.544s
sys	0m0.068s
```

# Memory usage

1. After process start: 5Mb
2. Peak memory usage for 10 concurrent clients: 9.5Mb
3. Peak memory usage for 100 concurrent clients: 10Mb


# Thread usage

4 threads. It has little sense with [GOMAXPROCS](https://golang.org/pkg/runtime/) :)


# JSON benchmarks

Results: 160K qps

```
./wrk -t 1 -c 20 -d 30s http://localhost:8090/json
Running 30s test @ http://localhost:8090/json
  1 threads and 20 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   137.60us  518.44us  22.56ms   98.28%
    Req/Sec   160.97k     8.89k  171.46k    89.70%
  4820802 requests in 30.10s, 836.28MB read
Requests/sec: 160160.69
Transfer/sec:     27.78MB
```

# Conclusions

Go is much faster than Swift and node.js :)
