#txcomment API

search all tx-comments for a particular string

#install

#### foundation

txcomment-search-api requires [foundation][4], a simple interface that allows easy access to the RPC commands offered by flojson. It also abstracts the RPC username/password away from sourcecode/config files. The best way to pass those values into foundation is by exporting them into an environment variable. We will come back to this later.

To install foundation, simply use `go get`.

```
go get github.com/metacoin/foundation
```

#### dbr

To install dbr, simply use `go get`.

```
go get github.com/gocraft/dbr
```

#### txcomment-search-api

To install dbr, simply use `go get`.

```
go get github.com/dloa/txcomment-search-api
```

## Building

Navigate to the directory containing `dloa/txcomment-search-api`. You'll build txcomment-search-api as you would any other go project:

```
go build -a main.go
```

This should create a binary `txcomment-search-api` in the Go bin directory.


**requires floblockexplorer**

## setup

open port 5831 in ACL or iptables

```bash
$ iptables -A INPUT -p tcp --dport 5831 -j ACCEPT
```

set environment variables

```bash
$ export DB_USER='dbusername'
$ export DB_PASS='reallygoodpassword10020229294#33334$$'
```

run the program

```bash
$ go run main.go
```

on your web browser or whatever client:

send a POST to the endpoint `/searchTxComment` to get a response

```json
{
	"search":"Alexandria",
	"page":0,
	"results-per-page":10
}
```

max results-per-page is 30
