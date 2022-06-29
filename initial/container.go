package initial

import "github.com/sarulabs/di"

func (initial *Initial) InitContainer(Container di.Container) *Initial {
	initial.Container = Container
	return initial
}
