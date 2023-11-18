package main

import (
	"fmt"
	"newProject/arrayFuncs"
	"newProject/flowControl"
	"newProject/mathPlus"
	"newProject/moreTypes"
)

var truthFind bool
	

func main() {
	part_I()
	fmt.Println(truthFind)
	mathPlus.Runner()
	flowControl.Runner()
	moreTypes.Runner()
	

	// fmt.Scanln()
}


func part_I()  {	
	myInt := arrayFuncs.ArrayCreater(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(myInt)

	myInt2 := arrayFuncs.ArrayCreaterFirstSix(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(myInt2)

	truthFind = true	
}

