package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (engine gasEngine) milesLeft() uint8 {
	return engine.gallons * engine.mpg
}

func (engine electricEngine) milesLeft() uint8 {
	return engine.kwh * engine.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(engine engine, miles uint8) {
	if miles < engine.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You need to refill the gas tank!")
	}
}

type owner struct {
	name string
}

func main() {
	//var myEngine gasEngine //takes default values for both
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)

	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15} //this is called an anonymous struct and its not reusable like gasEngine above
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	fmt.Printf("Total miles left in the tank: %v\n", myEngine.milesLeft())

	canMakeIt(myEngine, 30)

	var myEngine3 electricEngine = electricEngine{30, 20}
	canMakeIt(myEngine3, 40)
}
