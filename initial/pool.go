package initial

import "cloud/internal/ants"

func (initial *Initial) InitPool(poolSize int) *Initial {
	// 创建一个 Ants 池
	pool, err := ants.NewAnts(poolSize)
	if err != nil {
		panic(err)
	}
	initial.Pool = pool
	return initial
}
