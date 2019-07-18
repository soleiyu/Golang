package PulseLib

import "math"

type Dsmc struct {
	bitWidth	int
	sgm				int
	qc				int
	Vmax			int
}

func NewDsmc (bw int) *Dsmc {
	ins := new(Dsmc)
	ins.bitWidth = bw
	ins.sgm = 0
	ins.Vmax = int(math.Pow(2.0, float64(ins.bitWidth))) - 1

	return ins
}

func (ins *Dsmc) ClockUp (vadd int) bool {
	qc := false
	delta := ins.sgm + vadd

	if ins.Vmax < delta {
		qc = true
		ins.sgm = delta - ins.Vmax
	} else {
		ins.sgm = delta
	}

	return qc
}
