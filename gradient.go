package main

import "sort"

type Gradient struct {
	ColorStops []ColorStop
}

type ColorStop struct {
	Character  rune
	Luminosity float64
}

type SortOrder string

const (
	L2D SortOrder = "light-to-dark"
	D2L SortOrder = "dark-to-light"
)

func (g *Gradient) ToString(order SortOrder) string {
	sort.Slice(g.ColorStops, func(i, j int) bool {
		switch order {
		default:
			return g.ColorStops[i].Luminosity < g.ColorStops[j].Luminosity
		case D2L:
			return g.ColorStops[i].Luminosity > g.ColorStops[j].Luminosity

		}
	})

	charset := ""

	for i := 0; i < len(g.ColorStops); i++ {
		charset += string(g.ColorStops[i].Character)
	}

	return charset
}
