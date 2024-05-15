package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/brecht-vde/ascii-luminosity/pkg/luminosity"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Options struct {
	Font    string
	Mode    luminosity.Mode
	Charset luminosity.Charset
}

func main() {
	options := parseOptions()

	face, err := initFontFace(options.Font)

	if err != nil {
		fmt.Printf("%q\n", err)
		return
	}

	generator, err := luminosity.NewDefaultGenerator(&face)

	if err != nil {
		fmt.Printf("%q", err)
	}

	charset := luminosity.LoadCharset(options.Charset)
	gradient := generator.Generate(charset, options.Mode)

	fmt.Println(gradient)
}

func parseOptions() Options {
	options := Options{}

	var mode string
	var charset string

	flag.StringVar(&options.Font, "font", "", "path to a font file to calculate the luminosity of characters with.")
	flag.StringVar(&mode, "mode", "", "mode in which the gradient should be displayed ('Lightening' or 'Darkening').")
	flag.StringVar(&charset, "charset", "", "charset to use in the gradient ('Ascii' or 'Emoji').")

	flag.Parse()

	options.Mode = luminosity.Mode(mode)
	options.Charset = luminosity.Charset(charset)

	return options
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
