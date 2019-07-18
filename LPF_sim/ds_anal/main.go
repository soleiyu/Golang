package main

import (
	"fmt"
	"math"
//	"os"
//	"strconv"
	"./Lib"
)

func main() {

//	ins := Pwmc.NewPwmc(10)
	ins := PulseLib.NewDsmc(10)

	vcycle := 1000
	vnum := 50
	pan_num := 10
	val_inp := 0
	wave_num := 3
	rc_num := 100
	rc_ratio := 0.000001

	varr := make([]float64, vcycle * vnum * pan_num * wave_num)
	parr := make([]bool, vcycle * vnum * pan_num * wave_num)

	//ANALOG PARAM
	conv := 0.0
	e := 1.0
	//rc, _ := strconv.ParseFloat(os.Args[1], 64)
	dt := 1.0 / float64(vnum * vcycle * pan_num * 1000)

	for l := 1; l < rc_num + 1; l++ {

		for j := 0; j < vnum * wave_num; j++ {
			val_inp = int(float64(ins.Vmax) * (0.5 + 0.5 * math.Sin(2.0 * math.Pi * float64(j) / float64(vnum))))

			for i := 0; i < vcycle; i++ {
				if ins.ClockUp(val_inp) {
					for k := 0; k < pan_num; k++ {
						//conv = conv + vcu(e - conv, rc, dt)
						conv = conv + vcu(e - conv, float64(l) * rc_ratio, dt)
						varr[pan_num * (vcycle * j + i) + k] = conv
						parr[pan_num * (vcycle * j + i) + k] = true
					}
				} else {
					for k := 0; k < pan_num; k++ {
						//conv = vcd(conv, rc, dt)
						conv = vcd(conv, float64(l) * rc_ratio, dt)
						varr[pan_num * (vcycle * j + i) + k] = conv
						parr[pan_num * (vcycle * j + i) + k] = false
					}
				}
			}
		}

		vmax := anal_vmax(varr)
		updw := anal_updw(varr, parr)
	
		fmt.Println(float64(l) * rc_ratio, vmax, updw, vmax * (1.0 - updw))
	}
}

func anal_updw (varr []float64, parr []bool) float64 {
	vmax := 0.0

	ptv := make([]float64, 0)

	for i := 1; i < len(varr); i++ {
		if parr[i] != parr[i - 1] {
			ptv = append(ptv, varr[i])
		}
	}

	for i := 0; i < len(ptv) - 2; i++ {
		vline := 0.5 * (ptv[i] + ptv[i + 2])
		v := math.Abs(ptv[i + 1] - vline)

		if vmax < v {
			vmax = v
		}
	}

	return vmax
}
		
func anal_vmax (varr []float64) float64 {
	vmax := 0.0

	for i := 0; i < len(varr); i++ {
		if vmax < varr[i] {
			vmax = varr[i]
		}
	}

	return 2.0 * (vmax - 0.5)
}

func vcu (v, rc, t float64) float64 {
	return v * (1.0 - math.Exp(-t / rc))
}

func vcd (v, rc, t float64) float64 {
	return v * math.Exp(-t / rc)
}

