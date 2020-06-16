package round

import (
	"testing"
)

func TestRounding(t *testing.T) {
	f := Do(3.4444444444444444444445)
	if f != 4.0 {
		t.Fatalf("unexpected, actual: `%f`, expected: `%f`", f, 4.0)
	}
	f = Do(3.123412341234123412341234213412341234)
	if f != 3.0 {
		t.Fatalf("unexpected, actual: `%f`, expected: `%f`", f, 3.0)
	}
}
