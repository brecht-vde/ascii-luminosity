package rendering

import (
	"testing"

	"golang.org/x/image/font/basicfont"
)

func TestDefaultRenderer(t *testing.T) {
	tests := []struct {
		char     rune
		expected float64
	}{
		{
			char:     'A',
			expected: 0.991200,
		},
		{
			char:     'B',
			expected: 0.989600,
		},
	}

	renderer, err := NewDefaultRenderer(50, 50, 16, basicfont.Face7x13)

	if err != nil {
		t.Errorf(err.Error())
	}

	for _, tt := range tests {
		result := renderer.Render(tt.char)

		if result != tt.expected {
			t.Errorf("expected: %f\nreceived: %f", tt.expected, result)
		}
	}
}
