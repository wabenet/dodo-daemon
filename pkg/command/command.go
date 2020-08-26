package command

import (
	"github.com/spf13/cobra"
)

func NewDaemonCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "daemon",
		Short:            "run backdrops in daemon mode",
		TraverseChildren: true,
		SilenceUsage:     true,
	}

	cmd.AddCommand(NewDaemonStartCommand())
	cmd.AddCommand(NewDaemonStopCommand())
	cmd.AddCommand(NewDaemonRestartCommand())
	return cmd
}

func NewDaemonStartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "run a backdrop in daemon mode",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunContainer(args[0])
		},
	}
}

func NewDaemonStopCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "stop a daemon backdrop",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return StopContainer(args[0])
		},
	}
}

func NewDaemonRestartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "restart",
		Short: "restart a daemon backdrop",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RestartContainer(args[0])
		},
	}
}
