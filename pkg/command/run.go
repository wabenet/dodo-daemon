package command

import (
	api "github.com/dodo-cli/dodo-core/api/v1alpha2"
	"github.com/dodo-cli/dodo-core/pkg/core"
	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-core/pkg/plugin/configuration"
	"github.com/dodo-cli/dodo-core/pkg/plugin/runtime"
)

func RunContainer(m plugin.Manager, name string) error {
	config := configuration.AssembleBackdropConfig(m, name, &api.Backdrop{})
	config.ContainerName = config.Name

	if len(config.ImageId) == 0 {
		imageID, err := core.BuildByName(m, config.BuildInfo)
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
