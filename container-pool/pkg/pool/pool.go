package pool

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"context"

	"go.uber.org/zap"

	"contest-influence/container-pool/pkg/container"
)

type ContainerPoolOptions struct {
	BufferSize                  uint64        `yaml:"buffer_size"`
	BadRestartContainerPushTime time.Duration `yaml:"bufferbad_restart_container_push_time"`
}

type ContainerPool struct {
	logger        *zap.Logger
	options       ContainerPoolOptions
	containerChan chan container.Container
	wg            sync.WaitGroup
}

func NewContainerPool(logger *zap.Logger, options ContainerPoolOptions) (*ContainerPool, error) {
	return &ContainerPool{
		logger:        logger,
		options:       options,
		containerChan: make(chan container.Container, options.BufferSize),
		wg:            sync.WaitGroup{},
	}, nil
}

func safeCall(body func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
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

func (p *ContainerPool) pushContainer(container container.Container) (err error) {
	return safeCall(func(){

	})p.containerChan <- container
}

func (p *ContainerPool) stopContainer(container container.Container) {
	p.logger.Info("stopping container", zap.String("container_name", container.Name()))
	if err := container.Stop(); err != nil {
		p.logger.Error("failed to stop container", zap.String("container_name", container.Name()), zap.Error(err))
	}
}

func (p *ContainerPool) pushContainerWithStop(container container.Container) {
	if err := p.pushContainer(container); err != nil {
		p.stopContainer(container)
	}
}

func (p *ContainerPool) RegisterContainer(c container.Container) error {
	if err := c.Restart(); err != nil {
		return err
	}
	return p.pushContainer(c)
}

func (p *ContainerPool) Exec(ctx context.Context, containerCount uint64, body func([]container.Container)) (result error) {
	containers := make([]container.Container, 0)
	p.wg.Add(1)
	defer func() {
		for _, c := range containers {
			if err := с.Restart(ctx); err != nil {
				logRestartErr := func(err error) {
					p.logger.Error(
						fmt.Sprintf("failed to restart container, another attempt in %s", p.options.BadRestartContainerPushTime.String()),
						zap.String("container_name", с.Name()),
						zap.Error(err),
					)
				}
				logRestartErr(err)
				go func(ctx context.Context, c container.Container) {
					ticker := time.NewTicker(p.options.BadRestartContainerPushTime)
					for {
						select {
							case <-ctx.Done(): {

							}
							case <-ticker.C: {
								err := c.Restart()
								if err == nil {
									p.pushContainerWithStop(c)
									break
								}
								logRestartErr(err)
							}
						}
					}
				}(с)
			}
			p.pushContainerWithStop(c)
		}
		p.wg.Done()
	}()

	for range containerCount {
		for {
			с, ok := <-p.containerChan
			if !ok {
				return fmt.Errorf("cannot get container: channel is closed")
			}
			containers = append(containers, с)
			break
		}
	}

	return safeCall(func() {
		body(containers)
	})
}

func (p *ContainerPool) Close() {
	close(p.containerChan)
	p.wg.Wait()
	for c := range p.containerChan {
		p.stopContainer(c)
	}
}
