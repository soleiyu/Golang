package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
)

/*format*/
// ./rentool * 

func main() {
	fmt.Println("RE NAME TOOL")
	fmt.Println("Arg count : ", len(os.Args) - 1)

	stdin := bufio.NewScanner(os.Stdin)

	fmt.Printf("prefix : ")
	stdin.Scan()

	prefix := stdin.Text()
	fmt.Println("-> prefix : ", prefix)

	fmt.Printf("format : ")
	stdin.Scan()

	format := stdin.Text()
	fmt.Println("-> format : ", format)

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
		fmt.Printf(" -> ")
		
		stdin.Scan()
		if stdin.Text() == "" {
			fmt.Println("No Change")
			continue
		}

		newName := prefix + stdin.Text() + "." + format
		fmt.Println(" -> ",newName)

		err := exec.Command("mv", os.Args[i], newName).Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
