package main

import (
	"github.com/kaytu-io/kaytu/pkg/plugin/sdk"
	"github.com/kaytu-io/plugin-gcp/plugin"
)

func main() {
	sdk.New(plugin.NewPlugin(), 4).Execute()
}
