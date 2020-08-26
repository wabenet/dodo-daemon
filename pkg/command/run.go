package command

import (
	"fmt"

	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-core/pkg/plugin/configuration"
	"github.com/dodo-cli/dodo-core/pkg/plugin/runtime"
	"github.com/dodo-cli/dodo-core/pkg/types"
	log "github.com/hashicorp/go-hclog"
)

func RunContainer(name string) error {
	config := GetConfig(name)
	config.ContainerName = config.Name

	rt, err := GetRuntime()
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

	for _, p := range plugin.GetPlugins(configuration.Type.String()) {
		err := p.(configuration.Configuration).Provision(containerID)
		if err != nil {
			log.Default().Warn("could not provision", "error", err)
		}
	}

	return rt.StartContainer(containerID)
}

func StopContainer(name string) error {
	config := GetConfig(name)
	config.ContainerName = config.Name

	rt, err := GetRuntime()
	if err != nil {
		return err
	}

	return rt.RemoveContainer(config.ContainerName)
}

func RestartContainer(name string) error {
	config := GetConfig(name)
	config.ContainerName = config.Name

	rt, err := GetRuntime()
	if err != nil {
		return err
	}

	if err := rt.RemoveContainer(config.ContainerName); err != nil {
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

	for _, p := range plugin.GetPlugins(configuration.Type.String()) {
		err := p.(configuration.Configuration).Provision(containerID)
		if err != nil {
			log.Default().Warn("could not provision", "error", err)
		}
	}

	return rt.StartContainer(containerID)
}

func GetRuntime() (runtime.ContainerRuntime, error) {
	for _, p := range plugin.GetPlugins(runtime.Type.String()) {
		if rt, ok := p.(runtime.ContainerRuntime); ok {
			return rt, nil
		}
	}

	return nil, fmt.Errorf("could not find container runtime: %w", plugin.ErrNoValidPluginFound)
}

func GetConfig(name string) *types.Backdrop {
	config := &types.Backdrop{Name: name, Entrypoint: &types.Entrypoint{}}

	for _, p := range plugin.GetPlugins(configuration.Type.String()) {
		conf, err := p.(configuration.Configuration).UpdateConfiguration(config)
		if err != nil {
			log.L().Warn("could not get config", "error", err)
			continue
		}

		config.Merge(conf)
	}

	log.L().Debug("assembled configuration", "backdrop", config)
	return config
}
