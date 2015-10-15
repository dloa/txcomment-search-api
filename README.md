#txcomment API

open port 5831 in ACL or iptables

```bash
$ iptables -A INPUT -p tcp --dport 5831 -j ACCEPT
```

send a POST to the endpoint `/searchTxComment` to get a response

```json
{
	"search":"Alexandria",
	"page":0,
	"results-per-page":10
}
```

max results-per-page is 30
