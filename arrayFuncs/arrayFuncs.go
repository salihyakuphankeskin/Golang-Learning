package arrayFuncs

func ArrayCreater(intArray ...int)[]int{
	return intArray
}

func ArrayCreaterFirstSix(intArray ...int)[6]int{
	myArray := [6]int{}
	copy(myArray[:], intArray[:6])
	return myArray
}