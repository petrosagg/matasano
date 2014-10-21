package problem3

import (
	"flag"
	"fmt"

	"petrosagg/matasano/set1/problem1"
	"petrosagg/matasano/set1/problem2"
)

func Run() {
	flag.Parse()

	ciphertext := problem1.HexDecode(flag.Arg(0))

	key := make([]byte, len(ciphertext))

	freqs := make([]byte, 256)

	common_letters := []byte{' ', 'e', 't', 'a', 'o', 'i', 'n', 's', 'h', 'r', 'd', 'l', 'u'}

	cleartext := ""
	max_score := 0

	for k := 0; k < 256; k++ {
		// Fill the key buffer with the current guess
		for i := 0; i < len(key); i++ {
			key[i] = byte(k)
		}

		// Compute candidate cleartext
		candidate := problem2.Xor(ciphertext, key)

		// Zero out frequency buffer
		for i := 0; i < 256; i++ {
			freqs[i] = 0
		}

		for _, char := range candidate {
			freqs[char]++
		}

		max := byte(0)
		max_char := byte(0)
		for char, freq := range freqs {
			if freq > max {
				max = freq
				max_char = byte(char)
			}
		}

		score := 1 << uint(len(common_letters))
		for _, letter := range common_letters {
			if letter != max_char {
				score = score >> 1
			} else {
				break
			}
		}

		if score > max_score {
			cleartext = string(problem2.Xor(ciphertext, key))
			max_score = score
		}
	}

	fmt.Println(cleartext)
}
