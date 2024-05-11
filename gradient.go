package main

import (
	"sort"
)

type Gradient struct {
	Stops Stops
}

type Stops map[rune]float64

type Mode string

const (
	Lightening Mode = "Lightening"
	Darkening  Mode = "Darkening"
)

func (g *Gradient) Render(mode Mode) string {
	keys := g.getKeys()
	g.applyMode(keys, mode)
	return string(keys)
}

func (g *Gradient) getKeys() []rune {
	keys := make([]rune, 0, len(g.Stops))
	for key := range g.Stops {
		keys = append(keys, key)
	}
	return keys
}

func (g *Gradient) applyMode(keys []rune, mode Mode) {
	sort.Slice(keys, func(i, j int) bool {
		switch mode {
		default:
			return g.Stops[keys[i]] < g.Stops[keys[j]]
		case Darkening:
			return g.Stops[keys[i]] > g.Stops[keys[j]]
		}
	})
}
