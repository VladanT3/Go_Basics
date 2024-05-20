package main

import "fmt"

func main(){
	var intNum int16 = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 7
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString1 string = "Hello \nworld"
	var myString2 string = `Hello
	world`
	fmt.Println(myString1 + "\n" + myString2)
	fmt.Println(len(myString1))

	var myRune rune = 'a'
	fmt.Println(myRune)

	myVar := "text"
	fmt.Println(myVar)

	var var1, var2 = 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Println(myConst)
}