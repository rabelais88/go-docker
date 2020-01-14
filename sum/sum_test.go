package sum

import (
	"testing"
)

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
	}

	for _, v := range tests {
		_answer := mySum(v.data...)
		if _answer != v.answer {
			t.Error(`expect`, v.answer, `got`, _answer)
		}
	}
}
