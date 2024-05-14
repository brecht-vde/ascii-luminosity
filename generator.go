package main

import "sort"

type Rend interface {
	Render(char rune) float64
}

type Generator struct {
	Renderer Rend
}

func NewGenerator(renderer Rend) Generator {
	return Generator{
		Renderer: renderer,
	}
}

func (g *Generator) Generate(charset []rune, mode Mode) string {
	c := make([]rune, len(charset))
	copy(c, charset)

	luminosities := g.calculateLuminosities(c)
	g.applyMode(c, luminosities, mode)

	return string(c)
}

func (g *Generator) calculateLuminosities(charset []rune) map[rune]float64 {
	luminosities := make(map[rune]float64)

	for _, char := range charset {
		luminosity := g.Renderer.Render(char)
		luminosities[char] = luminosity
	}

	return luminosities
}

func (g *Generator) applyMode(charset []rune, luminosities map[rune]float64, mode Mode) {
	sort.Slice(charset, func(i, j int) bool {
		switch mode {
		default:
			return luminosities[charset[i]] < luminosities[charset[j]]
		case Darkening:
			return luminosities[charset[i]] > luminosities[charset[j]]
		}
	})
}
