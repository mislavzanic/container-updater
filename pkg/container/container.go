package container

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	t "github.com/mislavzanic/container-updater/pkg/types"
)

type Container struct {
	config container.Config
	id string
	hostConfig container.HostConfig
	name string
}

func NewContainer(payload t.Payload) Container {
	ps := nat.PortSet{}
	pm := nat.PortMap{}
	for _, pb := range payload.PortBindings {
		ps[nat.Port(pb.ContainerPort)] = struct{}{}
		pm[nat.Port(pb.ContainerPort)] = append(pm[nat.Port(pb.ContainerPort)], nat.PortBinding{HostIP: pb.HostIP, HostPort: pb.HostPort})
	}
	return Container{
		config: container.Config{
			Image: payload.Image,
			ExposedPorts: ps,
		},
		hostConfig: container.HostConfig{
			PortBindings: pm,
		},
		id: "",
		name: payload.Name,
	}
}
