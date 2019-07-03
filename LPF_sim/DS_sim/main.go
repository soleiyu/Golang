package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"./Lib"
)

func main() {

//	ins := PulseLib.NewPwmc(10)
	ins := PulseLib.NewDsmc(10)
	fmt.Println("#", ins.Vmax)

	vcycle := 1000
	vnum := 50
	pan_num := 10
	val_inp := 0
	wav_num := 1

	//ANALOG PARAM
	conv := 0.0
	e := 1.0
	rc, _ := strconv.ParseFloat(os.Args[1], 64)
	dt := 1.0 / float64(vnum * vcycle * pan_num * 1000)

	for j := 0; j < vnum * wav_num; j++ {
		val_inp = int(float64(ins.Vmax) * (0.5 + 0.5 * math.Sin(2.0 * math.Pi * float64(j) / float64(vnum))))

		for i := 0; i < vcycle; i++ {
			if ins.ClockUp(val_inp) {
				for k := 0; k < pan_num; k++ {
					conv = conv + vcu(e - conv, rc, dt)
					fmt.Println ((pan_num * (vcycle * j + i) + k) / pan_num, val_inp, 1, conv)
				}
			} else {
				for k := 0; k < pan_num; k++ {
					conv = vcd(conv, rc, dt)
					fmt.Println ((pan_num * (vcycle * j + i) + k) / pan_num, val_inp, 0, conv)
				}
			}
		}
	}
}

func vcu (v, rc, t float64) float64 {
	return v * (1.0 - math.Exp(-t / rc))
}

func vcd (v, rc, t float64) float64 {
	return v * math.Exp(-t / rc)
}

