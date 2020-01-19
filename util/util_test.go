package util

import "testing"

func TestMakeArray(t *testing.T) {
	_ary := MakeArray(5, 1)
	if len(_ary) != 5 {
		t.Error(`expected length 5, got `, len(_ary))
	}
	for i := range _ary {
		if _ary[i] != 1 {
			t.Error(`expected 1, got `, _ary[i])
		}
	}
}

func TestMakeRandom(t *testing.T) {
	_rand := MakeRandom().Intn(100000)
	_rand2 := MakeRandom().Intn(100000)
	if _rand == _rand2 {
		t.Error(_rand, `and`, _rand2, `should be differentiated`)
	}
}
