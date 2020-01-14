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
		_answer := MySum(v.data...)
		if _answer != v.answer {
			t.Error(`expect`, v.answer, `got`, _answer)
		}
	}
}

// https://godoc.org/github.com/rabelais88/go-docker/sum 이 주소에서 내용 확인 가능
func ExampleMySum() {
	MySum(1, 2)
}
