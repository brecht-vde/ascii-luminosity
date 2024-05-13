package main

import "testing"

func TestGenerator(t *testing.T) {
	expected := "AB"

	generator := NewGenerator(renderer, []rune{'A', 'B'}, "light")

	gradient := generator.Generate("mode")
}
