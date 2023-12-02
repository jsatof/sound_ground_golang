package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		want string
	}{
		{want: "Hello World"},
	}

	for _, tcase := range tests {
		got := HelloWorld()
		if tcase.want != got {
			t.Errorf("expected: %v, got: %v", tcase.want, got)
		}
	}
}
