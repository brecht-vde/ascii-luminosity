package main

import (
	"image"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

type Renderer struct {
	img  *image.Gray
	face *font.Face
}

func NewRenderer(width, height, size int, fnt *sfnt.Font) (*Renderer, error) {
	face, err := opentype.NewFace(fnt, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})

	if err != nil {
		return nil, err
	}

	img := image.NewGray(image.Rect(0, 0, width, height))

	return &Renderer{
		img:  img,
		face: &face,
	}, nil
}

func (r *Renderer) Render(char rune) []uint8 {
	r.reset()
	r.render(char)
	return r.img.Pix
}

func (r *Renderer) reset() {
	draw.Draw(r.img, r.img.Bounds(), image.White, r.img.Bounds().Min, draw.Src)
}

func (r *Renderer) render(char rune) {
	center := r.getCenter(char)

	d := &font.Drawer{
		Dst:  r.img,
		Src:  image.Black,
		Face: *r.face,
		Dot:  center,
	}

	d.DrawString(string(char))
}

func (r *Renderer) getCenter(char rune) fixed.Point26_6 {
	fw := fixed.I(r.img.Bounds().Max.X)
	fh := fixed.I(r.img.Bounds().Max.Y)

	bounds, _ := font.BoundString(*r.face, string(char))

	cx := (fw - bounds.Max.X) / 2
	cy := (fh - bounds.Max.Y) / 2

	return fixed.Point26_6{X: cx, Y: cy}
}
