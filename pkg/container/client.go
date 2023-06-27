package container

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/docker/docker/api/types"
	sdkClient "github.com/docker/docker/client"
)

type Client interface {
	StartImage(Container) (string, error)
	StopImage(string) error
	ListContainers() ([]types.Container, error)
	CreateUpdateContainer(Container) (string, error)
	UpdateContainer(Container) (error)
	RenameImage()
}

const defaultStopSignal = "SIGTERM"

func NewClient() Client {
	cli, err := sdkClient.NewClientWithOpts(sdkClient.FromEnv)

	if err != nil {
		log.Fatalf("Error instantiating docker client: %s", err)
	}

	return dockerClient{
		api: cli,
	}
}

type dockerClient struct {
	api sdkClient.CommonAPIClient
}

func (client dockerClient) ListContainers() ([]types.Container, error) {
	bg := context.Background()
	containers, err := client.api.ContainerList(bg, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func (client dockerClient) UpdateContainer(c Container) error {
	containers, err := client.ListContainers()
	if err != nil {
		return err
	}
	for _, container := range containers {
		if container.Image == c.config.Image {
			if err := client.PullImage(c.config.Image); err != nil {
				log.Fatal(err)
				return err
			}
			if err := client.StopImage(container.ID); err != nil {
				log.Fatal(err)
				return err
			} 
		}
	}
	return nil
}

func (client dockerClient) CreateUpdateContainer(c Container) (string, error) {
	containers, err := client.ListContainers()
	if err != nil {
		return "", err
	}
	for _, container := range containers {
		if container.Image == c.config.Image {
			if err := client.PullImage(c.config.Image); err != nil {
				log.Fatal(err)
				return "", err
			}
			if err := client.StopImage(container.ID); err != nil {
				log.Fatal(err)
				return "", err
			} else {
				id, err := client.StartImage(c)
				return id, err
			}
		}
	}
	id, err := client.StartImage(c)
	return "", nil
}

func (client dockerClient) StartImage(c Container) (string, error) {
	bg := context.Background()
	resp, err := client.api.ContainerCreate(bg, &c.config, &c.hostConfig, nil, nil, "blog")
	if err != nil {
		return "", err
	}

	if err := client.api.ContainerStart(bg, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}

	return resp.ID, nil
}

func (client dockerClient) StopImage(cID string) error {
	bg := context.Background()
	signal := defaultStopSignal
	if err := client.api.ContainerKill(bg, cID, signal); err != nil {
		return err
	}

	return nil
}

func (client dockerClient) RenameImage() {
	
}

func (client dockerClient) PullImage(image string) error {
	ctx := context.Background()
	response, err := client.api.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		log.Fatalf("Error pulling image %s, %s", image, err)
		return err
	}

	defer response.Close()
	if _, err = ioutil.ReadAll(response); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
