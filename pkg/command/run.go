package command

import (
	"fmt"

	api "github.com/dodo-cli/dodo-core/api/v1alpha1"
	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-core/pkg/plugin/configuration"
	"github.com/dodo-cli/dodo-core/pkg/plugin/runtime"
	"github.com/dodo-cli/dodo-core/pkg/types"
	log "github.com/hashicorp/go-hclog"
)

func RunContainer(name string) error {
	config := GetConfig(name)
	config.ContainerName = config.Name

	rt, err := GetRuntime(config.Runtime)
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
	config := GetConfig(name)
	config.ContainerName = config.Name

        rt, err := GetRuntime(config.Runtime)
	if err != nil {
		return err
	}

	return rt.DeleteContainer(config.ContainerName)
}

func RestartContainer(name string) error {
	config := GetConfig(name)
	config.ContainerName = config.Name

        rt, err := GetRuntime(config.Runtime)
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

func GetRuntime(name string) (runtime.ContainerRuntime, error) {
	for n, p := range plugin.GetPlugins(runtime.Type.String()) {
		if name != "" && name != n {
			continue
		}

		if rt, ok := p.(runtime.ContainerRuntime); ok {
			return rt, nil
		}
	}

	return nil, fmt.Errorf("could not find container runtime: %w", plugin.ErrNoValidPluginFound)
}

func GetConfig(name string) *api.Backdrop {
	config := &api.Backdrop{Name: name, Entrypoint: &api.Entrypoint{}}

	for _, p := range plugin.GetPlugins(configuration.Type.String()) {
		info, err := p.PluginInfo()
		if err != nil {
			log.L().Warn("could not read plugin info")
			continue
		}

		log.L().Debug("Fetching configuration from plugin", "name", info.Name)

		conf, err := p.(configuration.Configuration).GetBackdrop(name)
		if err != nil {
			log.L().Warn("could not get config", "error", err)
			continue
		}

		types.Merge(config, conf)
	}

	log.L().Debug("assembled configuration", "backdrop", config)
	return config
}
