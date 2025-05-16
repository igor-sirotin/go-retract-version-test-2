package foo

import "testing"

func TestFoo(t *testing.T) {
	const want = 42
	if got := Foo(); got != want {
		t.Errorf("Foo() = %d, want %d", got, want)
	}
}
