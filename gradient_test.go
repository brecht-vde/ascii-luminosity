package main

import "testing"

func TestGradientRender(t *testing.T) {
	data := map[rune]float64{
		'A': 0.00,
		'B': 0.25,
		'C': 0.50,
		'D': 0.75,
		'E': 1.00,
	}

	tests := []struct {
		name     string
		data     map[rune]float64
		mode     Mode
		expected string
	}{
		{
			name:     "Light mode gradient",
			data:     data,
			mode:     Lightening,
			expected: "ABCDE",
		},
		{
			name:     "Dark mode gradient",
			data:     data,
			mode:     Darkening,
			expected: "EDCBA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := &Gradient{
				Stops: tt.data,
			}

			result := sut.Render(tt.mode)

			if result != tt.expected {
				t.Errorf("Test '%s' failed. expected: '%s', received: '%s'", tt.name, tt.expected, result)
			}
		})
	}
}
