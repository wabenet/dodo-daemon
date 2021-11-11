package command

import (
	"fmt"

	api "github.com/dodo-cli/dodo-core/api/v1alpha2"
	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-core/pkg/plugin/builder"
	"github.com/dodo-cli/dodo-core/pkg/plugin/configuration"
	"github.com/dodo-cli/dodo-core/pkg/plugin/runtime"
	"github.com/dodo-cli/dodo-core/pkg/ui"
)

func RunContainer(m plugin.Manager, name string) error {
	config := configuration.AssembleBackdropConfig(m, name, &api.Backdrop{})
	config.ContainerName = config.Name

	if len(config.ImageId) == 0 {
		imageID, err := buildByName(m, config.BuildInfo)
		if err != nil {
			return err
		}

		config.ImageId = imageID
	}

	rt, err := runtime.GetByName(m, config.Runtime)
	if err != nil {
		return err
	}

	imageID, err := rt.ResolveImage(config.ImageId)
	if err != nil {
		return err
	}

	config.ImageId = imageID

	containerID, err := rt.CreateContainer(config, false, false)
	if err != nil {
		return err
	}

	return rt.StartContainer(containerID)
}

func StopContainer(m plugin.Manager, name string) error {
	config := configuration.AssembleBackdropConfig(m, name, &api.Backdrop{})
	config.ContainerName = config.Name

	rt, err := runtime.GetByName(m, config.Runtime)
	if err != nil {
		return err
	}

	return rt.DeleteContainer(config.ContainerName)
}

func RestartContainer(m plugin.Manager, name string) error {
	config := configuration.AssembleBackdropConfig(m, name, &api.Backdrop{})
	config.ContainerName = config.Name

	rt, err := runtime.GetByName(m, config.Runtime)
	if err != nil {
		return err
	}

	if err := rt.DeleteContainer(config.ContainerName); err != nil {
		return err
	}

	imageID, err := rt.ResolveImage(config.ImageId)
	if err != nil {
		return err
	}

	config.ImageId = imageID

	containerID, err := rt.CreateContainer(config, false, false)
	if err != nil {
		return err
	}

	return rt.StartContainer(containerID)
}

func buildByName(m plugin.Manager, overrides *api.BuildInfo) (string, error) {
	config, err := configuration.FindBuildConfig(m, overrides.ImageName, overrides)
	if err != nil {
		return "", fmt.Errorf("error finding build config for %s: %w", overrides.ImageName, err)
	}

	for _, dep := range config.Dependencies {
		conf := &api.BuildInfo{}
		configuration.MergeBuildInfo(conf, overrides)
		conf.ImageName = dep

		if _, err := buildByName(m, conf); err != nil {
			return "", err
		}
	}

	return buildImage(m, config)
}

func buildImage(m plugin.Manager, config *api.BuildInfo) (string, error) {
	b, err := builder.GetByName(m, config.Builder)
	if err != nil {
		return "", fmt.Errorf("could not find build plugin for %s: %w", config.Builder, err)
	}

	imageID := ""

	err = ui.NewTerminal().RunInRaw(
		func(t *ui.Terminal) error {
			if id, err := b.CreateImage(config, &plugin.StreamConfig{
				Stdin:          t.Stdin,
				Stdout:         t.Stdout,
				Stderr:         t.Stderr,
				TerminalHeight: t.Height,
				TerminalWidth:  t.Width,
			}); err != nil {
				return fmt.Errorf("error in container I/O stream: %w", err)
			} else {
				imageID = id
			}

			return nil
		},
	)

	return imageID, err
}
