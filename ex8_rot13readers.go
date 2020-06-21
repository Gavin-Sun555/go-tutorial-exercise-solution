package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func Rot13 (input byte) byte{
	
	if (input >= 'A' && input <= 'M') || (input >= 'a' && input <= 'm') {
		input += 13
	}else if(input >= 'N' && input <= 'Z') || (input >= 'n' && input <= 'z') {
		input -= 13
	}
	return input
}

func (rot rot13Reader) Read(b []byte) (int, error){
	n, err := rot.r.Read(b)

	for i := 0; i < len(b); i++ {
		b[i] = Rot13(b[i])
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
