package main

import (
	"fmt"

	"main/app"
)

var truthFind bool
	

func main() {
	part_I()
	fmt.Println(truthFind)
	c:=app.Adder(1,2)
	d:=app.Substracter(3,2)
	e:=app.HighOrder(2,3)
	fmt.Println(c," and ", d, " and ", e)
	fmt.Scanln()
}

func part_I()  {	
	myInt := arrayCreater(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(myInt)

	myInt2 :=arrayCreaterFirstSix(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(myInt2)

	truthFind = true	
}

func arrayCreater(intArray ...int)[]int{
	return intArray
}

func arrayCreaterFirstSix(intArray ...int)[6]int{
	myArray := [6]int{}
	copy(myArray[:], intArray[:6])
	return myArray
}