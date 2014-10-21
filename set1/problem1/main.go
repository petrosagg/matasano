package problem1

import (
	"bytes"
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

func B64Encode(src []byte) string {
	var buf = make([]byte, (len(src)+2)/3*4)

	for i := 0; i < len(src)/3; i++ {
		c1 := (src[i*3] & 0xfc) >> 2
		c2 := ((src[i*3] & 0x03) << 4) + ((src[i*3+1] & 0xf0) >> 4)
		c3 := ((src[i*3+1] & 0x0f) << 2) + ((src[i*3+2] & 0xc0) >> 6)
		c4 := src[i*3+2] & 0x3f

		buf[i*4+0] = getB64Char(c1)
		buf[i*4+1] = getB64Char(c2)
		buf[i*4+2] = getB64Char(c3)
		buf[i*4+3] = getB64Char(c4)
	}

	return string(buf)
}

func HexDecode(hex string) []byte {
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

func HexEncode(src []byte) string {
	hex := make([]byte, len(src)*2)

	for i := 0; i < len(src); i++ {
		b1 := src[i] & 0xf0 >> 4
		b2 := src[i] & 0x0f

		hex[i*2] = b1 + '0'
		hex[i*2+1] = b2 + '0'

		if b1 > 9 {
			hex[i*2] = src[i]&0x0f + 'a' - 10
		}

		if b2 > 9 {
			hex[i*2+1] = src[i]&0x0f + 'a' - 10
		}
	}

	return string(hex)
}

func Run() {
	flag.Parse()

	hex := flag.Arg(0)

	fmt.Println(B64Encode(HexDecode(hex)))
}
