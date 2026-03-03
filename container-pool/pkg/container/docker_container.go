package container

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	docker_client "github.com/docker/docker/client"
)

type DockerContainerOptions struct {
	Name        string        `yaml:"container_name"`
	HttpTimeout time.Duration `yaml:"http_timeout"`
	Host        string        `yaml:"host"`
	Version     string        `yaml:"version"`
}

type DockerContainer struct {
	client  *docker_client.Client
	logger  *zap.Logger
	options DockerContainerOptions
}

// var _ Container = (*DockerContainer)(nil)

func NewDockerContainer(logger *zap.Logger, options DockerContainerOptions) (*DockerContainer, error) {
	var err error
	container := &DockerContainer{}

	if options.Name == "" {
		return nil, fmt.Errorf("docker container must have not empty name")
	}

	if logger == nil {
		return nil, fmt.Errorf("docker container must have logger")
	}

	docker_opts := make([]docker_client.Opt, 0)
	if options.HttpTimeout != 0 {
		docker_opts = append(docker_opts, docker_client.WithTimeout(options.HttpTimeout))
	}

	if options.Host != "" {
		docker_opts = append(docker_opts, docker_client.WithHost(options.Host))
	}

	if options.Version != "" {
		docker_opts = append(docker_opts, docker_client.WithVersion(options.Version))
	}

	container.client, err = docker_client.NewClientWithOpts(docker_opts...)

	if err != nil {
		return nil, err
	}

	container.logger = logger
	container.options = options

	return container, nil
}
