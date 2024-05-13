package main

import (
	"reflect"
	"testing"
)

type MockMeter struct {
	Values map[rune]float64
}

func NewMockMeter() *MockMeter {
	return &MockMeter{
		map[rune]float64{
			'A': 0.0,
			'B': 0.25,
			'C': 0.50,
		},
	}
}

func (m *MockMeter) Measure(character rune) float64 {
	if v, ok := m.Values[character]; ok {
		return v
	}

	return 0.0
}

func TestCalc(t *testing.T) {
	m := NewMockMeter()

	tests := []struct {
		characters []rune
		meter      *MockMeter
		expected   *Gradient
	}{
		{
			characters: []rune{
				'A',
				'B',
				'C',
			},
			meter: m,
			expected: &Gradient{
				GradientStops: []GradientStop{
					{
						Character:  'A',
						Luminosity: 0.0,
					},
					{
						Character:  'B',
						Luminosity: 0.25,
					},
					{
						Character:  'C',
						Luminosity: 0.5,
					},
				},
			},
		},
		{
			characters: []rune{
				'B',
				'C',
				'1',
			},
			meter: m,
			expected: &Gradient{
				GradientStops: []GradientStop{
					{
						Character:  'B',
						Luminosity: 0.25,
					},
					{
						Character:  'C',
						Luminosity: 0.5,
					},
					{
						Character:  '1',
						Luminosity: 0.0,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		c := NewCalc(tt.characters, tt.meter)
		g := c.Calculate()

		if !reflect.DeepEqual(g, tt.expected) {
			t.Errorf("Output '%+v' does not equal expected '%+v'", g, tt.expected)
		}
	}
}
