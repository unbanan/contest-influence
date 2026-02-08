package container

import (
	"github.com/stretchr/testify/mock"
)

type MockContainer struct {
	mock.Mock
}

var _ Container = (*MockContainer)(nil)

func (m *MockContainer) Name() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockContainer) Init() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockContainer) Start() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockContainer) Stop() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockContainer) Restart() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockContainer) Exec(commands ...string) ContainerExecResult {
	args := m.Called(commands)
	return args.Get(0).(ContainerExecResult)
}
