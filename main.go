package main

import (
	"os"

	"github.com/wabenet/dodo-daemon/pkg/plugin"
)

func main() {
	os.Exit(plugin.RunMe())
}
