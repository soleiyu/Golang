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
		//fmt.Println(qt(3.0, 0.000001, 10.0, float64(i) * 0.001))
		//fmt.Println(it(3.0, 0.00001, 1.0, float64(i) * 0.0001))
		fmt.Println(float64(i) * 0.01, vc(e, c, r, float64(i) * 0.00001))
	}

	fmt.Println(bsvct(e, c, r, 2.0))

}

func bsvct(e, c, r, v float64) float64 {
	t1 := 0.0
	t2 := 9.9

	for ;; {
		tc := (t1 + t2) / 2.0
		vtc := vc(e, c, r, tc)

		if math.Abs(v - vtc) < 1e-16 {
			return tc
		}

		if vtc < v {
			t1 = tc
		} else {
			t2 = tc
		}
	}
}

func qt (e, c, r, t float64) float64 {
	return c * e * (1.0 - math.Exp(-t / (r * c)))
}

func it (e, c, r, t float64) float64 {
	return e * math.Exp(-t / (r * c)) / r
}

func vc (e, c, r, t float64) float64 {
	return e - r * it(e, c, r, t)
}
