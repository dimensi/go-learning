package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

var cachedSymbols []byte

func getSliceSymbols() []byte {

	if len(cachedSymbols) != 0 {
		return cachedSymbols
	}

	str := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	result := make([]byte, 0, len(str))

	for _, val := range str {
		result = append(result, byte(val))
	}

	cachedSymbols = result
	return result
}

func getRotSymbol(symbol byte) byte {
	var index int
	symbols := getSliceSymbols()

	for idx, val := range symbols {
		if symbol == val {
			index = idx
			break
		}
	}

	if index == 0 {
		return symbol
	}

	if index+13 > len(symbols) {
		return symbols[index-13]
	}

	return symbols[index+13]
}

func (r13 rot13Reader) Read(bytes []byte) (int, error) {
	n, _ := r13.r.Read(bytes)
	str := bytes[:n]

	for index, val := range str {
		str[index] = getRotSymbol(val)
	}

	return n, io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
