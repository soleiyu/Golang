package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"strings"
)

/*format*/
// ./aydtool list 

func main() {
	fmt.Println("Auto Youtube-dl TOOL v1.0")
	fmt.Println("List :", os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	urls := strings.Split(string(b), "\n")

	for i := 0; i < len(urls); i++ {
		if len(urls[i]) < 3 {
			continue
		}

		fmt.Println(i + 1, "/", len(urls), urls[i])
		err = exec.Command("youtube-dl", urls[i]).Run()

		if err != nil {
			fmt.Println("ERR youtube-dl")
		}
	}
}	
