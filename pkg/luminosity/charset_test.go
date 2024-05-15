package luminosity

import "testing"

func TestCharset(t *testing.T) {
	tests := []struct {
		charset  Charset
		expected string
	}{
		{
			charset:  Ascii,
			expected: " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		},
		{
			charset:  Ansi,
			expected: " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~\u00a0¡¢£¤¥¦§¨©ª«¬\u00ad®¯°±²³´µ¶·¸¹º»¼½¾¿ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿ",
		},
	}

	for _, tt := range tests {
		result := LoadCharset(tt.charset)

		if string(result) != tt.expected {
			t.Errorf("expected: %s\nreceived: %s", tt.expected, string(result))
		}
	}
}
