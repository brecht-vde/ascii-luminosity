package main

import "sort"

type Gradient struct {
	GradientStops []GradientStop
}

type GradientStop struct {
	Character  rune
	Luminosity float64
}

type GradientDirection string

const (
	L2D GradientDirection = "light-to-dark"
	D2L GradientDirection = "dark-to-light"
)

func (g *Gradient) Add(character rune, luminosity float64) {
	stop := &GradientStop{
		Character:  character,
		Luminosity: luminosity,
	}

	g.GradientStops = append(g.GradientStops, *stop)
}

func (g *Gradient) Render(direction GradientDirection) string {
	g.sort(direction)
	return g.generate()
}

func (g *Gradient) generate() string {
	charset := ""

	for i := 0; i < len(g.GradientStops); i++ {
		charset += string(g.GradientStops[i].Character)
	}

	return charset
}

func (g *Gradient) sort(direction GradientDirection) {
	sort.Slice(g.GradientStops, func(i, j int) bool {
		switch direction {
		default:
			return g.GradientStops[i].Luminosity > g.GradientStops[j].Luminosity
		case D2L:
			return g.GradientStops[i].Luminosity < g.GradientStops[j].Luminosity
		}
	})
}
