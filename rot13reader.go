package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func generateRot(count byte, offset byte) func(byte, byte) byte {
	return func(base byte, character byte) (newChar byte) {
		charRelativeToBase := character - base
		charWithOffset := charRelativeToBase + offset
		charWithingScope := charWithOffset % count
		goBackToBase := charWithingScope + base
		return goBackToBase
	}
}

func (self rot13Reader) Read(p []byte) (n int, err error) {
	n, err = self.r.Read(p)

	offset := byte(13)
	A := byte(65)
	Z := byte(90)
	a := byte(97)
	z := byte(122)
	charsInRange := byte(26)

	offsetChar := generateRot(charsInRange, offset)

	for i, char := range p {
		switch {
		case char >= A && char <= Z:
			p[i] = offsetChar(A, char)
		case char >= a && char <= z:
			p[i] = offsetChar(a, char)
		}
	}
	return
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
