# go-web-benchmark
Go version of the Swift Web Benchmark.  https://medium.com/@qutheory/server-side-swift-vs-the-other-guys-1-input-ec48d7be37b7

# Getting it up and going

`go get -u github.com/willie/go-web-benchmark`

# Building

`go build bench.go`

# Benchmarking

(On OS X, you need to raise ulimits for testing: `ulimit -n 2048`)

`wrk -d 10 -t 4 -c 128 http://localhost:8300/plaintext
wrk -d 10 -t 4 -c 128 http://localhost:8300/json
wrk -d 10 -t 4 -c 128 http://localhost:8300/sqlite-fetch
`

# Stuff I did to get the benchmark at Digital Ocean server setup (per the test setup)

Installed Go according to: https://www.digitalocean.com/community/tutorials/how-to-install-go-1-6-on-ubuntu-14-04

Then: `go get -u github.com/willie/go-web-benchmark`

`cd ~/go-willie/work/src/github.com/willie/go-web-benchmark; go build bench.go; cp bench ~/go-willie/bench; cd ~/go-willie`

# Acknowledgements

`https://mholt.github.io/json-to-go/
github.com/jmoiron/sqlx
github.com/mattn/go-sqlite3
`
