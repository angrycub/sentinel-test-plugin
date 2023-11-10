package main

import (
	sdk "github.com/hashicorp/sentinel-sdk"
	"github.com/hashicorp/sentinel-sdk/framework"
	"github.com/hashicorp/sentinel-sdk/rpc"
)

func main() {
	rpc.Serve(&rpc.ServeOpts{
		ImportFunc: func() sdk.Import {
			return &framework.Import{Root: &root{}}
		},
	})
}
