package pool

import (
	"errors"
	"fmt"

	"go.uber.org/zap"

	"contest-influence/container-pool/pkg/container"
)

type ContainerPoolOptions struct {
}

type ContainerPool struct {
	logger        *zap.Logger
	options       ContainerPoolOptions
	containerChan chan container.Container
}

func NewContainerPool(logger *zap.Logger, options ContainerPoolOptions) (*ContainerPool, error) {
	return &ContainerPool{
		logger:        logger,
		options:       options,
		containerChan: make(chan container.Container),
	}, nil
}

func safeCall(body func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// Преобразуем panic в error
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("%v", x)
			}
		}
	}()

	body()
	return nil
}

func (p *ContainerPool) pushContainer(container container.Container) error {
	select {
	case p.containerChan <- container:
		return nil
	default:
		return fmt.Errorf("Cannot push container, pool is closed")
	}
}

func (p *ContainerPool) stopContainer(container container.Container) {
	p.logger.Info("Stopping container", zap.String("container_name", container.Name()))
	container.Stop()
}

func (p *ContainerPool) RegisterContainer(container container.Container) error {
	return p.pushContainer(container)
}

func (p *ContainerPool) Exec(containerCount uint64, body func([]container.Container)) error {
	containers := make([]container.Container, 0)

	defer func() {
		for _, c := range containers {
			c.Restart()
			err := p.pushContainer(c)
			if err != nil {
				p.stopContainer(c)
			}
		}
	}()

	for range containerCount {
		container, ok := <-p.containerChan
		if !ok {
			return fmt.Errorf("Cannot get container: channel is closed")
		}
		container.Restart()
		containers = append(containers, container)
	}

	return safeCall(func() {
		body(containers)
	})
}

func (p *ContainerPool) Close() {
	close(p.containerChan)
	for c := range p.containerChan {
		p.stopContainer(c)
	}
}
