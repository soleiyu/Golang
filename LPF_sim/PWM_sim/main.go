package main

import (
	"fmt"
	"math"
)

func main() {

	ins := NewPwmc(5)
	fmt.Println("#", ins.Vmax)

	num_plot := 500
	val_inp := 16

	for i := 0; i < num_plot; i++ {
		if ins.ClockUp(val_inp) {
			fmt.Println (i, 1)
		} else {
			fmt.Println (i, 0)
		}
	}
}

type Pwmc struct {
	bitWidth	int
	Cnt		int
	qc		int
	Vmax		int
}

func NewPwmc (bw int) *Pwmc {
	ins := new(Pwmc)
	ins.bitWidth = bw
	ins.Cnt = 0
	ins.Vmax = int(math.Pow(2.0, float64(ins.bitWidth))) - 1

	return ins
}

func (ins *Pwmc) ClockUp (vq int) bool {
	qc := true

	if vq < ins.Cnt {
		qc = false
	}

	ins.Cnt += 1
	if ins.Vmax < ins.Cnt {
		ins.Cnt = 0
	}

	return qc
}
