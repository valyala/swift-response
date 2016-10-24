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
$ wrk -t 1 -c 20 -d 60s --latency http://10.10.xx.xxx:8090/json
Running 1m test @ http://10.10.xx.xxx:8090/json
  1 threads and 20 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   248.10us  660.77us  29.34ms   98.64%
    Req/Sec    95.98k     4.22k  104.16k    85.02%
  Latency Distribution
     50%  176.00us
     75%  222.00us
     90%  346.00us
     99%    0.97ms
  5739608 requests in 1.00m, 1.96GB read
Requests/sec:  95501.72
Transfer/sec:     33.33MB
```

Results: 95500 qps, 248us average latency.
This is 17x faster than [the fastest Swift code](https://medium.com/@rymcol/benchmarks-for-the-top-server-side-swift-frameworks-vs-node-js-24460cfe0beb#7d04) with 5.7K qps and 3520us average latency :)


20 concurrent connections is laughable, so below are Go results for 20K
concurrent connections:

```
$ wrk -t 1 -c 20000 -d 60s --latency http://10.10.xx.xxx:8090/json
Running 1m test @ http://10.10.xx.xxx:8090/json
  1 threads and 20000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   142.28ms   54.24ms   1.01s    85.73%
    Req/Sec    76.41k    12.74k  109.64k    68.27%
  Latency Distribution
     50%  127.46ms
     75%  136.15ms
     90%  224.86ms
     99%  358.61ms
  4533633 requests in 1.00m, 1.54GB read
Requests/sec:  75427.92
Transfer/sec:     26.32MB
```

Peak memory usage for 20K concurrent connections: 500Mb.


# Just for fun: nodejs performance on the same hardware

Node.js code has been taken [from the original benchmark](https://github.com/rymcol/Server-Side-Swift-Benchmarking/tree/949e5e75ab9c2b9c1741c704db330e15817c85bb/NodeJSON). Nodejs version:

```
$ nodejs --version
v4.2.6
```

20 concurrent connections:

```
$ wrk -t 1 -c 20 -d 60s --latency http://10.10.xx.xxx:8091/json
Running 1m test @ http://10.10.xx.xxx:8091/json
  1 threads and 20 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.75ms  378.70us  24.16ms   93.40%
    Req/Sec     7.32k   431.18     8.05k    76.21%
  Latency Distribution
     50%    2.70ms
     75%    2.85ms
     90%    3.02ms
     99%    3.56ms
  437404 requests in 1.00m, 170.57MB read
Requests/sec:   7277.96
Transfer/sec:      2.84MB
```

20K concurrent connections:

```
$ wrk -t 1 -c 20000 -d 60s --latency http://10.10.xx.xxx:8091/json
Running 1m test @ http://10.10.xx.xxx:8091/json
  1 threads and 20000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.34s   423.74ms   2.00s    65.37%
    Req/Sec     9.33k     3.71k   21.79k    82.67%
  Latency Distribution
     50%    1.38s 
     75%    1.67s 
     90%    1.98s 
     99%    2.00s 
  390175 requests in 1.00m, 152.15MB read
  Socket errors: connect 0, read 7, write 62627, timeout 131594
Requests/sec:   6493.70
Transfer/sec:      2.53MB
```

Pay attention to latencies and the `errors` line :)

Peak memory usage for 20K concurrent connections: 700Mb.


# Conclusions

Go is much faster than Swift and node.js. And it uses less memory :)
