package main

import (
    "fmt"
    "os"
		"./pictFunc"
)

func main(){
	fmt.Println(os.Args[1])
	hogera := pictFunc.Pict{}
	hogera.Load(os.Args[1])

//	ans := pictFunc.KCol(hogera, 1, 4)

	a1 := pictFunc.Xhsample(hogera)
	a2 := pictFunc.X2sample(a1)

//	ans.Save("res.png")

	a1.Save("a1.png")
	a2.Save("a2.png")
