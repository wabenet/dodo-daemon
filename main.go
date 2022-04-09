package main

import (
	"os"

	"github.com/dodo-cli/dodo-daemon/pkg/plugin"
)

func main() {
	os.Exit(plugin.RunMe())
}
