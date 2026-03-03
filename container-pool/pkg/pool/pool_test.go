package pool

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"contest-influence/container-pool/pkg/container"
)

type TestPoolContext struct {
	logger *zap.Logger
}

func initContext() (context TestPoolContext) {
	context.logger, _ = zap.NewDevelopment()
	return
}

func TestPoolJustWorks(t *testing.T) {
	context := initContext()
	defer context.logger.Sync()

	pool, err := NewContainerPool(context.logger, ContainerPoolOptions{
		BufferSize:                  10,
		BadRestartContainerPushTime: time.Millisecond,
	})

	require.NoError(t, err)

	container1 := &container.MockContainer{}

	container1.On("Restart").Return(nil)
	container1.On("Stop").Return(nil)
	container1.On("Name").Return("C1")

	require.NoError(t, pool.RegisterContainer(container1))

	for range 2 {
		pool.Exec(1, func(containers []container.Container) {
			require.Len(t, containers, 1)
			assert.Equal(t, "C1", containers[0].Name())
		})
	}

	pool.Close()
	container1.AssertExpectations(t)
}

func TestPoolClose(t *testing.T) {
	context := initContext()
	defer context.logger.Sync()

	const containerCount = 10
	pool, err := NewContainerPool(context.logger, ContainerPoolOptions{
		BufferSize:                  containerCount,
		BadRestartContainerPushTime: time.Millisecond,
	})

	require.NoError(t, err)

	containers := make([]*container.MockContainer, 0)

	for i := range containerCount {
		mockContainer := &container.MockContainer{}
		mockContainer.On("Restart").Return(nil)
		mockContainer.On("Stop").Return(nil).Once()
		mockContainer.On("Name").Return(fmt.Sprintf("C%d", i))
		containers = append(containers, mockContainer)
	}

	for _, c := range containers {
		require.NoError(t, pool.RegisterContainer(c))
	}

	wg := sync.WaitGroup{}
	cnt := [2]int{0, 0}

	for i := range 2 {
		go wg.Go(func() {
			for {
				number := rand.Int()%containerCount + 1

				err := pool.Exec(uint64(number), func(containers []container.Container) {
					uniqueContainers := make(map[string]struct{})
					for _, c := range containers {
						uniqueContainers[c.Name()] = struct{}{}
					}
					require.Len(t, uniqueContainers, number)
				})

				if err != nil {
					return
				}
				cnt[i]++
			}
		})
	}

	time.Sleep(time.Second)
	pool.Close()
	wg.Wait()

	assert.InDelta(t, float64(cnt[0])/float64(cnt[0]+cnt[1]), 0.5, 0.1)
	assert.Greater(t, cnt[0]+cnt[1], 1000)
	for _, c := range containers {
		c.AssertExpectations(t)
	}
}
