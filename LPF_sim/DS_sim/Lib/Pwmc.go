package PulseLib

import "math"

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


