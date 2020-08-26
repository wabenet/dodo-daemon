package plugin

import (
	"github.com/dodo-cli/dodo-core/pkg/appconfig"
	dodo "github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-daemon/pkg/command"
	log "github.com/hashicorp/go-hclog"
)

func RunMe() int {
	log.SetDefault(log.New(appconfig.GetLoggerOptions()))
	p := &command.Command{}
	if err := p.GetCobraCommand().Execute(); err != nil {
		return 1
	}
	return 0
}

func IncludeMe() {
	dodo.IncludePlugins(&command.Command{})
}
