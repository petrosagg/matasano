package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
)

func getB64Char(c byte) byte {
	switch {
	case 0 <= c && c <= 25:
		return c + 'A'
	case 26 <= c && c <= 51:
		return c + 'a' - 26
	case 52 <= c && c <= 61:
		return c + '0' - 52
	case c == 62:
		return '+'
	case c == 63:
		return '/'
	}
	return '*'
}

func b64Encode(src []byte) string {
	var ret = ""
	for i := 0; i < len(src)/3; i++ {
		c1 := (src[i*3] & 0xfc) >> 2
		c2 := ((src[i*3] & 0x03) << 4) + ((src[i*3+1] & 0xf0) >> 4)
		c3 := ((src[i*3+1] & 0x0f) << 2) + ((src[i*3+2] & 0xc0) >> 6)
		c4 := src[i*3+2] & 0x3f

		ret += c1
		fmt.Printf("%c", getB64Char(c1))
		fmt.Printf("%c", getB64Char(c2))
		fmt.Printf("%c", getB64Char(c3))
		fmt.Printf("%c", getB64Char(c4))
	}
	fmt.Printf("\n")

	return hex.Dump(src)
}

func hexDecode(hex string) []byte {
	var buff bytes.Buffer

	for i := 0; i < len(hex)/2; i++ {
		b1 := hex[i*2] - '0'
		b2 := hex[i*2+1] - '0'

		if b1 > 9 {
			b1 = b1 - 'a' + '0' + 10
		}

		if b2 > 9 {
			b2 = b2 - 'a' + '0' + 10
		}

		buff.WriteByte(byte(b1<<4 + b2))
	}

	return buff.Bytes()
}

func main() {
	flag.Parse()

	hex := flag.Arg(0)

	fmt.Println(b64Encode(hexDecode(hex)))
}
