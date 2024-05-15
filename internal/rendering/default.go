package rendering

import (
	"image"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type DefaultRenderer struct {
	img  *image.Gray
	face font.Face
}

func NewDefaultRenderer(width, height, size int, face font.Face) (*DefaultRenderer, error) {
	img := image.NewGray(image.Rect(0, 0, width, height))

	return &DefaultRenderer{
		img:  img,
		face: face,
	}, nil
}

func (r *DefaultRenderer) Render(char rune) float64 {
	r.reset()
	r.render(char)
	return r.calculateLuminosity(r.img.Pix)
}

func (r *DefaultRenderer) reset() {
	draw.Draw(r.img, r.img.Bounds(), image.White, r.img.Bounds().Min, draw.Src)
}

func (r *DefaultRenderer) render(char rune) {
	center := r.getCenter(char)

	d := &font.Drawer{
		Dst:  r.img,
		Src:  image.Black,
		Face: r.face,
		Dot:  center,
	}

	d.DrawString(string(char))
}

func (r *DefaultRenderer) getCenter(char rune) fixed.Point26_6 {
	fw := fixed.I(r.img.Bounds().Max.X)
	fh := fixed.I(r.img.Bounds().Max.Y)

	bounds, _ := font.BoundString(r.face, string(char))

	cx := (fw - bounds.Max.X) / 2
	cy := (fh - bounds.Max.Y) / 2

	return fixed.Point26_6{X: cx, Y: cy}
}

func (r *DefaultRenderer) calculateLuminosity(pixels []uint8) float64 {
	length := float64(len(pixels))
	sum := 0.0

	for i := 0; i < len(pixels); i++ {
		sum += float64(pixels[i] / 255)
	}

	return sum / length
}
