package main

type Rend interface {
	Render(char rune) float64
}

type Generator struct {
	Renderer Rend
}

func New(renderer Rend) Generator {
	return Generator{
		Renderer: renderer,
	}
}

func (g *Generator) Generate(charset []rune, mode Mode) string {
	var luminosities map[rune]float64

	for _, char := range charset {
		luminosity := g.Renderer.Render(char)
		luminosities[char] = luminosity
	}

	return ""
}
