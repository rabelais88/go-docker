package util

import (
	"math/rand"
	"time"
)

func MakeArray(size int, value int) []int {
	_ary := make([]int, size)
	for i := 0; i < size; i++ {
		_ary[i] = value
	}
	return _ary
}

func MakeRandom() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source)
}
