package luminosity

type Charset string

const (
	Ascii Charset = "Ascii"
	Ansi  Charset = "Ansi"
)

func LoadCharset(charset Charset) []rune {
	switch charset {
	case Ansi:
		return genAnsi()
	default:
		return genAscii()
	}
}

func genAscii() []rune {
	length := 127 - 32
	chars := make([]rune, length)

	for i := 0; i < length; i++ {
		chars[i] = rune(i + 32)
	}

	return chars
}

func genAnsi() []rune {
	length := 256 - 160
	chars1 := genAscii()
	chars2 := make([]rune, length)

	for i := 0; i < length; i++ {
		chars2[i] = rune(i + 160)
	}

	chars1 = append(chars1, chars2...)
	return chars1
}
