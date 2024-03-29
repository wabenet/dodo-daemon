package plugin

import (
	"github.com/wabenet/dodo-core/pkg/plugin"
	"github.com/wabenet/dodo-core/pkg/plugin/command"
	impl "github.com/wabenet/dodo-daemon/internal/plugin/command"
)

func RunMe() int {
	m := plugin.Init()

	if err := NewCommand(m).GetCobraCommand().Execute(); err != nil {
		return 1
	}

	return 0
}

func IncludeMe(m plugin.Manager) {
	m.IncludePlugins(NewCommand(m))
}

func NewCommand(m plugin.Manager) command.Command {
	return impl.New(m)
}
