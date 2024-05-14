package main

import "testing"

type MockRenderer struct{}

func (m *MockRenderer) Render(char rune) float64 {
	switch char {
	case 'A':
		return 0.25
	case 'B':
		return 0.50
	default:
		return 0.00
	}
}

func TestGenerator(t *testing.T) {
	expected := "AB"

	renderer := &MockRenderer{}
	generator := NewGenerator(renderer)

	gradient := generator.Generate([]rune{'A', 'B'}, Lightening)

	if gradient != expected {
		t.Errorf("expected: %q, received: %q", expected, gradient)
	}
}
