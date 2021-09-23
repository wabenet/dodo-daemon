package plugin

import (
	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-daemon/pkg/command"
)

func RunMe() int {
	m := plugin.Init()

	if err := command.New(m).GetCobraCommand().Execute(); err != nil {
		return 1
	}

	return 0
}

func IncludeMe(m plugin.Manager) {
	m.IncludePlugins(command.New(m))
}
