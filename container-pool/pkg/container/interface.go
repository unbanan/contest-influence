package container

type ContainerExecResult struct {
	ReturnCode int
	Trace      string
}

type Container interface {
	Name() string
	Init() error
	Start() error
	Stop() error
	Restart() error
	Exec(...string) ContainerExecResult
}
