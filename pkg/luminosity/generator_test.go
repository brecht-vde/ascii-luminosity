package luminosity

import (
	"testing"
)

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

	tests := []struct {
		mode     Mode
		expected string
	}{
		{
			mode:     Darkening,
			expected: "BA",
		},
		{
			mode:     Lightening,
			expected: "AB",
		},
	}

	renderer := &MockRenderer{}
	generator := NewGenerator(renderer)

	for _, tt := range tests {
		result := generator.Generate([]rune{'A', 'B'}, tt.mode)

		if result != tt.expected {
			t.Errorf("expected: %s\nreceived: %s", tt.expected, result)
		}
	}
}
