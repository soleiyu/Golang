package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"strconv"
)

/*format*/
// ./amp3 * 

func main() {
	fmt.Println("AUTO FFMPEG TO MP3 v1.1")
	fmt.Println("Arg count : ", len(os.Args) - 1)

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
		fmt.Println(strconv.Itoa(i))

		err := exec.Command("ffmpeg", "-i", os.Args[i],
			strings.Split(os.Args[i], ".")[0] + ".mp3").Run()

		if err != nil {
			fmt.Println(err)
		}
	}
}
