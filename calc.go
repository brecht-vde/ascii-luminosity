package main

type Meter interface {
	Measure(character rune) float64
}

type Calc struct {
	Meter      Meter
	Characters []rune
}

func NewCalc(characters []rune, meter Meter) *Calc {
	return &Calc{
		Characters: characters,
		Meter:      meter,
	}
}

func (c *Calc) Calculate() *Gradient {
	g := &Gradient{}

	for _, character := range c.Characters {
		luminosity := c.Meter.Measure(character)
		g.Add(character, luminosity)
	}

	return g
}
