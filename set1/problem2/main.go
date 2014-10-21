package problem2

import (
	"flag"
	"fmt"

	"petrosagg/matasano/set1/problem1"
)

func Xor(hex1, hex2 []byte) []byte {
	ret := make([]byte, len(hex1))

	for i := 0; i < len(hex1); i++ {
		ret[i] = hex1[i] ^ hex2[i]
	}
	return ret
}

func Run() {
	flag.Parse()

	hex1 := problem1.HexDecode(flag.Arg(0))
	hex2 := problem1.HexDecode(flag.Arg(1))

	fmt.Println(hex1)
	fmt.Println(hex2)
	fmt.Println(problem1.HexEncode(Xor(hex1, hex2)))
}
