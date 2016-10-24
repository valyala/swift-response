This is a Go response to https://medium.com/@rymcol/benchmarks-for-the-top-server-side-swift-frameworks-vs-node-js-24460cfe0beb

Currently only `/json` handler is implemented. `/blog` implementation is left as an exercise for readers :)
Go version used - `go tip` aka `go 1.8`.

# Hosting & Environment

Test server configuration:
```
$ uname -a
Linux xxx 4.4.0-45-generic #66-Ubuntu SMP Wed Oct 19 14:12:37 UTC 2016 x86_64 x86_64 x86_64 GNU/Linux

$ cat /proc/cpuinfo | grep name | head -1
model name	: Intel(R) Xeon(R) CPU E3-1230 v3 @ 3.30GHz
```

`wrk` was run from a separate server with identical configuration.

Both servers are connected over 1Gbps network.


# Build Time Results

```
time go build

real	0m0.466s
user	0m0.544s
sys	0m0.068s
```

This is 68x faster than the best build time for Swift code - [32 seconds](https://medium.com/@rymcol/benchmarks-for-the-top-server-side-swift-frameworks-vs-node-js-24460cfe0beb#4039).

# Memory Usage

1. After process start: 5Mb
2. Peak memory usage for 10 concurrent clients: 9.5Mb
3. Peak memory usage for 100 concurrent clients: 10Mb

Compare to [Swift memory usage](https://medium.com/@rymcol/benchmarks-for-the-top-server-side-swift-frameworks-vs-node-js-24460cfe0beb#8d9b).


# Thread usage

It has little sense with [GOMAXPROCS](https://golang.org/pkg/runtime/) :)


# JSON benchmarks

```
wrk -t 1 -c 20 -d 60s --latency http://10.10.xx.xx:8090/json
Running 1m test @ http://10.10.xx.xxx:8090/json
  1 threads and 20 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   131.55us  241.49us  16.35ms   99.53%
    Req/Sec   151.30k     8.94k  161.57k    94.83%
  Latency Distribution
     50%  116.00us
     75%  136.00us
     90%  160.00us
     99%  273.00us
  9031856 requests in 1.00m, 1.53GB read
Requests/sec: 150530.40
Transfer/sec:     26.11MB
```

Results: 150K qps, 132us average latency:
This is 26x faster than [the fastest Swift code](https://medium.com/@rymcol/benchmarks-for-the-top-server-side-swift-frameworks-vs-node-js-24460cfe0beb#7d04) with 5.7K qps and 3520us average latency :)


20 concurrent connections is laughable, so below are Go results for 20K
concurrent connections:

```
wrk -t 1 -c 20000 -d 60s --latency http://10.10.xx.xxx:8090/json
Running 1m test @ http://10.10.xx.xxx:8090/json
  1 threads and 20000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   101.56ms   10.43ms 260.03ms   70.53%
    Req/Sec    98.60k    10.97k  133.15k    86.98%
  Latency Distribution
     50%  100.92ms
     75%  108.40ms
     90%  113.86ms
     99%  123.18ms
  5867954 requests in 1.00m, 0.99GB read
Requests/sec:  97734.68
Transfer/sec:     16.95MB
```

Peak memory usage for 20K concurrent connections: 360Mb.


# Conclusions

Go is much faster than Swift and node.js. And it uses less memory :)
