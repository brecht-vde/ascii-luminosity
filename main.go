package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func main() {

}

func drawCharacter(character string) *image.Gray {
	img := image.NewGray(image.Rect(0, 0, 50, 50))
	draw.Draw(img, img.Bounds(), image.White, img.Bounds().Min, draw.Src)

	fontFile, _ := os.Open("FiraCodeNerdFontMono-Regular.ttf")
	defer fontFile.Close()

	fontData, _ := io.ReadAll(fontFile)
	fontItem, _ := opentype.Parse(fontData)

	face, _ := opentype.NewFace(fontItem, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	bounds, _ := font.BoundString(face, character)

	imgWidth := fixed.I(50)
	imgHeight := fixed.I(50)

	cx := (imgWidth - bounds.Max.X) / 2
	cy := (imgHeight - bounds.Min.Y) / 2

	d := &font.Drawer{
		Dst:  img,
		Src:  image.Black,
		Face: face,
		Dot:  fixed.Point26_6{X: cx, Y: cy},
	}

	d.DrawString(character)

	file, _ := os.Create(fmt.Sprintf("%v.png", character))
	defer file.Close()
	png.Encode(file, img)

	return img
}
