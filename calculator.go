package main

type Calculator struct {
	renderer *Renderer
}

func NewCalculator(renderer *Renderer) *Calculator {
	return &Calculator{
		renderer: renderer,
	}
}

func (c *Calculator) Calculate() Gradient {
	var colorStops []ColorStop

	for i := 32; i < 127; i++ {
		char := rune(i)

		pixels := c.renderer.Render(char)
		average := c.average(pixels)

		colorStops = append(colorStops, ColorStop{
			Character:  char,
			Luminosity: average,
		})
	}

	return Gradient{
		ColorStops: colorStops,
	}
}

func (c *Calculator) average(pixels []uint8) float64 {
	length := float64(len(pixels))
	sum := 0.0

	for i := 0; i < len(pixels); i++ {
		sum += float64(pixels[i] / 255)
	}

	return sum / length
}
