#txcomment API

search all tx-comments for a particular string

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
