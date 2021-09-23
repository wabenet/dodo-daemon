package command

import (
	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/spf13/cobra"
)

func New(m plugin.Manager) *Command {
	cmd := &cobra.Command{
		Use:              "daemon",
		Short:            "run backdrops in daemon mode",
		TraverseChildren: true,
		SilenceUsage:     true,
	}

	cmd.AddCommand(NewDaemonStartCommand(m))
	cmd.AddCommand(NewDaemonStopCommand(m))
	cmd.AddCommand(NewDaemonRestartCommand(m))

	return &Command{cmd: cmd}
}

func NewDaemonStartCommand(m plugin.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "run a backdrop in daemon mode",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunContainer(m, args[0])
		},
	}
}

func NewDaemonStopCommand(m plugin.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "stop a daemon backdrop",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return StopContainer(m, args[0])
		},
	}
}

func NewDaemonRestartCommand(m plugin.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "restart",
		Short: "restart a daemon backdrop",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RestartContainer(m, args[0])
		},
	}
}
