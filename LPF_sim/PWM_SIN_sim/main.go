package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	ins := NewPwmc(10)
	fmt.Println("#", ins.Vmax)

	orthant, err := strconv.Atoi(os.Args[1])
	fmt.Println("# ORTHANT :", orthant, err)

	vcycle := 1000
	vnum := 50
	val_inp := 0

	for j := 0; j < vnum; j++ {
//		val_inp = int(float64(ins.Vmax) * (0.5 + 0.5 * math.Sin(0.5 * math.Pi * (float64(orthant) + float64(j) / float64(vnum)))))
		val_inp = int(float64(ins.Vmax) * (0.5 + 0.5 * math.Sin(2.0 * math.Pi * float64(j) / float64(vnum))))

		for i := 0; i < vcycle; i++ {
			if ins.ClockUp(val_inp) {
				fmt.Println (vcycle * j + i, val_inp, 1)
			} else {
				fmt.Println (vcycle * j + i, val_inp, 0)
			}
		}
	}
}

type Pwmc struct {
	bitWidth	int
	Cnt				int
	qc				int
	Vmax			int
}

func NewPwmc (bw int) *Pwmc {
	ins := new(Pwmc)
	ins.bitWidth = bw
	ins.Cnt = 0
	ins.Vmax = int(math.Pow(2.0, float64(ins.bitWidth))) - 1

	return ins
}

func (ins *Pwmc) ClockUp (vq int) bool {
	qc := false

	if ins.Cnt < vq {
		qc = true
	}

	ins.Cnt += 1
	if ins.Vmax < ins.Cnt {
		ins.Cnt = 0
	}

	return qc
}

