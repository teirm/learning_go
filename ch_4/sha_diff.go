package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: ./sha_diff <arg1> <arg2>\n")
		os.Exit(1)
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	c1 := sha256.Sum256([]byte(arg1))
	c2 := sha256.Sum256([]byte(arg2))

	fmt.Printf("%x\n%x\n", c1, c2)

	bitDifference := compareBits(c1, c2)
	fmt.Printf("SHA bit difference: %d\n", bitDifference)
}

func compareBits(c1, c2 [32]uint8) int32 {

	var bitDifference int32

	for index, byte1 := range c1 {
		byte2 := c2[index]
		for i := 0; i < 8; i++ {
			bit1 := byte1 & 1
			bit2 := byte2 & 1
			if bit1 != bit2 {
				bitDifference++
			}
			byte1 = byte1 >> 1
			byte2 = byte2 >> 1
		}
	}
	return bitDifference
}
