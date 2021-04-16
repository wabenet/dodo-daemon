package command

import (
	api "github.com/dodo-cli/dodo-core/api/v1alpha1"
	"github.com/dodo-cli/dodo-core/pkg/core"
)

func RunContainer(name string) error {
	config := core.AssembleBackdropConfig(name, &api.Backdrop{})
	config.ContainerName = config.Name

	if len(config.ImageId) == 0 {
		imageID, err := core.BuildByName(config.BuildInfo)
		if err != nil {
			return err
		}

		config.ImageId = imageID
	}

	rt, err := core.GetRuntime(config.Runtime)
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

func StopContainer(name string) error {
	config := core.AssembleBackdropConfig(name, &api.Backdrop{})
	config.ContainerName = config.Name

	rt, err := core.GetRuntime(config.Runtime)
	if err != nil {
		return err
	}

	return rt.DeleteContainer(config.ContainerName)
}

func RestartContainer(name string) error {
	config := core.AssembleBackdropConfig(name, &api.Backdrop{})
	config.ContainerName = config.Name

	rt, err := core.GetRuntime(config.Runtime)
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
