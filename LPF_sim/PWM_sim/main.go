package main

import "fmt"

func main() {

	num_plot := 50
	pwmins := NewPwmc(5)
	val_inp := lb ("01000")

	for i := 0; i < num_plot; i++ {
		if pwmins.ClockUp(val_inp) {
			fmt.Println (i, 1)
		} else {
			fmt.Println (i, 0)
		}
	}

}

type Pwmc struct {
	count []bool
	bitWidth int
}

func bitPrint(inp []bool) {
	for i := 0; i < len(inp); i++ {
		if inp[len(inp) - 1 - i] {
			fmt.Printf("1 ")
		} else {
			fmt.Printf("0 ")
		}
	}
	fmt.Println()
}

func NewPwmc(bw int) *Pwmc {
	ins := new(Pwmc)
	ins.bitWidth = bw
	ins.count = make([]bool, bw)
	return ins
}

func (ins *Pwmc) ClockUp (vq []bool) bool {
	qc := fullSubber(
		append(vq, lb("1")...), append(ins.count, lb("0")...))

	ins.count = countUp(ins.count)

	return qc[ins.bitWidth]
}

func lb(inp string) []bool {
	ans := make([]bool, len(inp))

	for i := 0; i < len(ans); i++ {
		if inp[len(ans) - 1 - i] == '1' {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans
}

func allFlip(inp []bool) []bool {
	ans := make([]bool, len(inp))

	for i := 0; i < len(ans); i++ {
		ans[i] = !inp[i]
	}

	return ans
}

func countUp(inp []bool) []bool {
	ans := make([]bool, len(inp))
	c := true

	for i := 0; i < len(ans); i++ {
		ans[i] = xor3(inp[i], c, false)
		c = inp[i] && c
	}

	return ans
}

func fullAdder(a, b []bool) []bool {
	ans := make([]bool, len(a))
	c := false

	for i := 0; i < len(ans); i++ {
		ans[i] = xor3(a[i], b[i], c)
		c = (a[i] && b[i]) || (b[i] && c) || (c && a[i])
	}

	return ans
}

func fullSubber(a, b []bool) []bool {
	return fullAdder(a, allFlip(b))
}

func xor3(a, b, c bool) bool {
	return a && b && c || a && !b && !c || 
		!a && b && !c || !a && !b && c
}
