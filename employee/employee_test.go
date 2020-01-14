package employee

import (
	"testing"
)

func TestNew(t *testing.T) {
	myEmployee := New(`james`, `jack`, 3, 1)
	if myEmployee.firstName != `james` {
		t.Error(`wrong`)
	}
}
