# Installation

`go get -u github.com/zeropsio/zerops-go`

# Basic example

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/zeropsio/zerops-go/sdk"
	"github.com/zeropsio/zerops-go/sdkBase"
)

func main() {
	ctx := context.Background()

	zdk := sdk.New(
		sdkBase.DefaultConfig(),
		http.DefaultClient,
	)

	authorizedSdk := sdk.AuthorizeSdk(zdk, os.Getenv("token"))

	info, err := authorizedSdk.GetUserInfo(ctx)
	if err != nil {
		panic(err)
	}

	i, err := info.Output()
	if err != nil {
		panic(err)
	}

	clientId := i.ClientUserList[0].ClientId
	fmt.Println(clientId)
}
```