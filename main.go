package main

import (
	"fmt"
	"newProject/arrayFuncs"
	"newProject/mathPlus"
	"newProject/flowControl"
)

var truthFind bool
	

func main() {
	part_I()
	fmt.Println(truthFind)
	c:= mathPlus.Adder(1,2)
	d:= mathPlus.Substracter(3,2)
	e:= mathPlus.HighOrder(2,3)
	fmt.Println(c," and ", d, " and ", e)
	flowControl.MyForLoop(5)


	fmt.Scanln()
}


func part_I()  {	
	myInt := arrayFuncs.ArrayCreater(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(myInt)

	myInt2 := arrayFuncs.ArrayCreaterFirstSix(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(myInt2)

	truthFind = true	
}

