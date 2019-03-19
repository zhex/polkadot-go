# polkadot-go

## install

```
$ go get github.com/zhex/polkadot-go
```

## use

```
import (
    "fmt"
    
    "github.com/zhex/polkadot-go/client" 
)

c := client.New("wss://poc3-rpc.polkadot.io/")
name, err := c.RPC.System.Name()
fmt.Println(name, err)
```