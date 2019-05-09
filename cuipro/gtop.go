package main

import (
	"fmt"
	"os/exec"
	"strings"
	"strconv"
	"math"
)

func main() {
	cnum := coreNum()
	fmt.Println("CPUs : ", cnum)

	out, _ := exec.Command("top", "-b", "-n", "1", "-i").Output()
	lines := strings.Split(string(out), "\n")

	i := 0
	tcpu := float32(0)
	tmem := float32(0)

	// GOMI
	for ; i < len(lines); i++ {
		if -1 < strings.Index(lines[i], "PID") {
			i++
			break
		}
	}

	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}

		ccpu, cmem := gvs(lines[i])
		tcpu += ccpu
		tmem += cmem
	}

	fmt.Printf("CPU : %3.1f%%\n", tcpu / float32(cnum))
	fmt.Println(mkbar(100, tcpu / float32(cnum)))

	fmt.Printf("MEM : %3.1f%%\n", tmem)
	fmt.Println(mkbar(100, tmem))
}

func coreNum() int {
	out, _ := exec.Command("lscpu").Output()
	lines := strings.Split(string(out), "\n")
	vals := strings.Split(lines[3], " ")
	cnum, _ := strconv.Atoi(vals[len(vals) - 1])

	return cnum
}

func gvs(inp string) (float32, float32) {
	strs := strings.Split(inp, " ")
	vals := make([]string, 32)
	j := 0
	
	for i := 0; i < len(strs); i++ {
		if strs[i] != "" {
			vals[j] = strs[i]
			j++
		}
	}

	vcpu, _ := strconv.ParseFloat(vals[8], 32)
	vmem, _ := strconv.ParseFloat(vals[9], 32)
	
	return float32(vcpu), float32(vmem)
}

func fcfu(width, ratio int) int {
	val := float64(width) * float64(ratio) / 100.0
	return int(math.Floor(val + .5))
}

func mkbar(width int, uratio float32) string {
	val := float32(width) * uratio / float32(100)
	uv := int(math.Floor(float64(val + .5)))
	iv := width - uv

	//fmt.Println(uv, iv)

	us := "\x1b[47m \x1b[41m"
	is := "\x1b[40m"
	ls := "\x1b[47m \x1b[0m"

	for i := 0; i < uv; i++ {
		us += " "
	}
	for i := 0; i < iv; i++ {
		is += " "
	}

	return us + is + ls
}
