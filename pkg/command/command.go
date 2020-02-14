package command

import (
	"github.com/oclaussen/dodo/pkg/config"
	"github.com/oclaussen/dodo/pkg/container"
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
			conf, err := config.LoadBackdrop(args[0])
			if err != nil {
				return err
			}

			c, err := container.NewContainer(conf, config.LoadAuthConfig(), true)
			if err != nil {
				return err
			}

			return c.Run()
		},
	}
}

func NewDaemonStopCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "stop a daemon backdrop",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := config.LoadBackdrop(args[0])
			if err != nil {
				return err
			}

			c, err := container.NewContainer(conf, config.LoadAuthConfig(), true)
			if err != nil {
				return err
			}

			return c.Stop()
		},
	}
}

func NewDaemonRestartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "restart",
		Short: "restart a daemon backdrop",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := config.LoadBackdrop(args[0])
			if err != nil {
				return err
			}

			c, err := container.NewContainer(conf, config.LoadAuthConfig(), true)
			if err != nil {
				return err
			}

			if err := c.Stop(); err != nil {
				return err
			}

			return c.Run()
		},
	}
}
