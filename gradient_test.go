package main

import "testing"

func TestGradientRenders(t *testing.T) {
	tests := []struct {
		direction GradientDirection
		expected  string
		items     []struct {
			character  rune
			luminosity float64
		}
	}{
		{
			direction: GradientDirection(D2L),
			expected:  "ABCDE",
			items: []struct {
				character  rune
				luminosity float64
			}{
				{
					character:  'A',
					luminosity: 0,
				},
				{
					character:  'B',
					luminosity: 0.25,
				},
				{
					character:  'C',
					luminosity: 0.50,
				},
				{
					character:  'D',
					luminosity: 0.75,
				},
				{
					character:  'E',
					luminosity: 1,
				},
			},
		},
		{
			direction: GradientDirection(L2D),
			expected:  "EDCBA",
			items: []struct {
				character  rune
				luminosity float64
			}{
				{
					character:  'A',
					luminosity: 0,
				},
				{
					character:  'B',
					luminosity: 0.25,
				},
				{
					character:  'C',
					luminosity: 0.50,
				},
				{
					character:  'D',
					luminosity: 0.75,
				},
				{
					character:  'E',
					luminosity: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		g := &Gradient{}

		for _, item := range tt.items {
			g.Add(item.character, item.luminosity)
		}

		if charset := g.Render(tt.direction); charset != tt.expected {
			t.Errorf("Output '%q' does not equal expected '%q'", charset, tt.expected)
		}
	}
}
