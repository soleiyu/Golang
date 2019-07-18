package main

import (
    "fmt"
//    "image"
//    _ "image/jpeg"
    "os"
		"./pictFunc"
)

func main(){
	fmt.Println(os.Args[1])
	hogera := pictFunc.Pict{}
	hogera.Load(os.Args[1])

	//ans := pictFunc.MyGaus(hogera, 101)

//	ans := pictFunc.Xhsample(hogera)
//	ans = pictFunc.X2sample(ans)
	ans := pictFunc.KCol(hogera, 1, 4)

	ans.Save("res.png")
}
