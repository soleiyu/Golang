package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("ARG ERR E C R")
		return
	}

	e, _ := strconv.ParseFloat(os.Args[1], 64)
	c, _ := strconv.ParseFloat(os.Args[2], 64)
	r, _ := strconv.ParseFloat(os.Args[3], 64)
	fmt.Println("#e :", e)
	fmt.Println("#c :", c)
	fmt.Println("#r :", r)

	for i := 0; i < 100; i++ {
//		fmt.Println(float64(i) * 0.01, vc2(e, c, r, float64(i) * 0.00001))
	}

	mul := 2
	max := 100 * mul
	div := 10
	mxt := 1 * mul //ms
	tdv := float64(mxt) / (1000.0 * float64(max))
	dtm := tdv * 1000.0
	conv := 0.0

	for j := 0; j < max / (2 * div) ; j++ {
		for i := 0; i < div; i++ {
			fmt.Println(float64(j * div * 2 + i) * dtm,  conv + vcu(e - conv, c, r, float64(i) * tdv))
		}
		conv = conv + vcu(e - conv, c, r, float64(div) * tdv)

		for i := 0; i < div; i++ {
			fmt.Println(float64(j * div * 2 + div + i) * dtm, vcd(conv, c, r, float64(i) * tdv))
		}
		conv = vcd(conv, c, r, float64(div) * tdv)
	}
}

func vcu (v, c, r, t float64) float64 {
	return v * (1.0 - math.Exp(-t / (r * c)))
}

func vcd (v, c, r, t float64) float64 {
	return v * math.Exp(-t / (r * c))
}

