package main

import (
	"os"

	"github.com/dodo-cli/dodo-daemon/plugin"
)

func main() {
	os.Exit(plugin.RunMe())
}
