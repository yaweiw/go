package main

import (
	"fmt"

	funcs "github.com/yaweiw/go/funcs"
	mets "github.com/yaweiw/go/methods"
)

func main() {
	N := 0
	defer funcs.Timeinday()
	for i := 0; i <= N; i++ {
		fmt.Println("Sqrtfunc(2, 3):" + funcs.Sqrtfunc(2, 3))
		fmt.Println("funcs.GetSqrtfunc()(2, 3):" + funcs.GetSqrtfunc()(2, 3))
		fmt.Println("SqrtfuncInvoker(-4, -5, Sqrtfunc):" + funcs.SqrtfuncInvoker(-4, -5, funcs.Sqrtfunc))
		fmt.Println("SqrtfuncInvoker(-4, -3, Sqrtfunc):" + funcs.SqrtfuncInvoker(-4, -3, funcs.Sqrtfunc))
		c := mets.Coordinate{X: "32.534", Y: "67.246"}
		fmt.Println(c.ToString())
		var i mets.ICoordinate = mets.Coordinate{X: "133.53683", Y: "6537.24476"}
		fmt.Println(i.ToString())
	}

	//funcs.F()
	fmt.Println("Returned normally from f.")
}
