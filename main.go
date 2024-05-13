package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func main() {
	var font string
	flag.StringVar(&font, "font", "", "path to a font file to calculate the luminosity of characters with.")
	flag.Parse()

	face, err := initFontFace(font)

	if err != nil {
		fmt.Printf("%q\n", err)
		return
	}

	renderer, err := NewRenderer(50, 50, 16, face)

	if err != nil {
		fmt.Printf("%q\n", err)
		return
	}

	a := renderer.Render('A')

	fmt.Println(a)
}

func initFontFace(path string) (font.Face, error) {
	file, err := os.Open(path)

	if err != nil {
		err = fmt.Errorf("could not open font file %q", err)
		return nil, err
	}

	data, err := io.ReadAll(file)

	if err != nil {
		err = fmt.Errorf("could not read font file %q", err)
		return nil, err
	}

	fnt, err := opentype.Parse(data)

	if err != nil {
		err = fmt.Errorf("could not parse font file %q", err)
		return nil, err
	}

	face, err := opentype.NewFace(fnt, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	if err != nil {
		err = fmt.Errorf("could not create face %q", err)
		return nil, err
	}

	return face, nil
}
