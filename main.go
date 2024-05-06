package main

import (
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
	ascii := populateAscii()
	calculateLuminosity(ascii)
}

func populateAscii() []string {
	var ascii []string

	for i := 32; i < 127; i++ {
		ascii = append(ascii, string(rune(i)))
	}

	return ascii
}

func calculateLuminosity(ascii []string) {
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

	bounds, _ := font.BoundString(face, "A")

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

	d.DrawString(".")

	file, _ := os.Create("output.png")
	defer file.Close()
	png.Encode(file, img)
}
