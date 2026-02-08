package pool

import (
	"context"
	"testing"
	"time"

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

func testWithTimeout(t *testing.T, timeout time.Duration, test func(t *testing.T)) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan bool)
}

func TestPoolJustWorks(t *testing.T) {
	context := initContext()
	defer context.logger.Sync()

	pool, err := NewContainerPool(context.logger, ContainerPoolOptions{})

	require.NoError(t, err)

	container1 := &container.MockContainer{}

	container1.On("Restart").Return(nil)
	container1.On("Stop").Return(nil)

	require.NoError(t, pool.RegisterContainer(container1))

	pool.Acquire
	require.NoError(t, pool.Stop())

}
