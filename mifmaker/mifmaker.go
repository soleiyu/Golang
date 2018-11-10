package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	writer(16, 16, "ipwm4b.mif")
}

func writer (width, depth int, fn string) {
	file, _ := os.Create(fn)
	defer file.Close()

	file.Write(([]byte)("--  A____A\n"))
	file.Write(([]byte)("-- |・ㅅ・|\n"))
	file.Write(([]byte)("-- |っ　ｃ|\n"))
	file.Write(([]byte)("-- |　　　|\n"))
	file.Write(([]byte)("--  U￣￣U\n\n"))

	file.Write(([]byte)("WIDTH="))
	file.Write(([]byte)(strconv.Itoa(width)))
	file.Write(([]byte)(";\n"))
	file.Write(([]byte)("DEPTH="))
	file.Write(([]byte)(strconv.Itoa(depth)))
	file.Write(([]byte)(";\n\n"))

	file.Write(([]byte)("ADDRESS_RADIX=HEX;\n"))
	file.Write(([]byte)("DATA_RADIX=BIN;\n\n"))

	file.Write(([]byte)("CONTENT BEGIN\n"))
	p := math.Log2((float64)(width))

	for i := 0; i < depth; i++ {
		h :=	fmt.Sprintf("  %04X  :  ", i)
		file.Write(([]byte)(h))

		if i < depth / 2 {
			pbs := pwmBits((int)(p), i)

			for j := 0; j < width; j++ {
				if pbs[j] {
					file.Write(([]byte)("1"))
				} else {
					file.Write(([]byte)("0"))
				}
			}
			file.Write(([]byte)(";\n"))
		} else {
			pbs := pwmBits((int)(p), depth - i - 1)

			for j := 0; j < width; j++ {
				if pbs[j] {
					file.Write(([]byte)("0"))
				} else {
					file.Write(([]byte)("1"))
				}
			}
			file.Write(([]byte)(";\n"))
		}	
	}

	file.Write(([]byte)("END;\n"))
}

func pwmBits (bit, val int) []bool {
	maxv := (int)(math.Pow(2, (float64)(bit)))
	bits := make([]bool, maxv)

	for i := 0; i < val; i++ {
		qv := i * maxv / val
		bits[qv] = true
	}

	return bits
}

