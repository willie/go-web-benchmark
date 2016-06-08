# go-web-benchmark
Go version of the Swift Web Benchmark.  https://medium.com/@qutheory/server-side-swift-vs-the-other-guys-1-input-ec48d7be37b7

Props to Matt Holt for the always useful JSON to Go

`https://mholt.github.io/json-to-go/`

# Project Dependencies

`go install github.com/mattn/go-sqlite3`
`go get github.com/jmoiron/sqlx`

# Building

`go build bench.go`

# Benchmarking

On OS X, you need to raise ulimits for testing:
`ulimit -n 2048`

`wrk -d 10 -t 4 -c 128 http://localhost:8300/plaintext`
`wrk -d 10 -t 4 -c 128 http://localhost:8300/json`
`wrk -d 10 -t 4 -c 128 http://localhost:8300/sqlite-fetch`
