package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

/*format*/
// ./amp3 * 

func main() {
	fmt.Println("AUTO FFMPEG TO MP3")
	fmt.Println("Arg count : ", len(os.Args) - 1)


	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
		fmt.Println(strconv.Itoa(i))
		err := exec.Command("ffmpeg", "-i", os.Args[i], strconv.Itoa(i) + ".mp3")
			if err != nil {
				fmt.Println("ERR")
			}
	}
}
