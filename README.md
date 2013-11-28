unqlitego
=========

UnQLite Binding for golang.

Install
---------

```sh
go get -d github.com/nobonobo/unqlitego
cd $GOPATH/src/github.com/nobonobo/unqlitego
make install
```

Benchmark
----------

```
BenchmarkFileStore	  200000	      9667 ns/op
BenchmarkFileFetch	  500000	      7928 ns/op
BenchmarkMemStore	  500000	      3824 ns/op
BenchmarkMemFetch	 1000000	      3448 ns/op
```
