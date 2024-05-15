package luminosity

type Renderer interface {
	Render(char rune) float64
}
