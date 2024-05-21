package main

import (
	"fmt"
)

func main() {
	intArr := [...]int32{1, 2, 3} //array
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6} //slice
	fmt.Printf("The length is %v with capacity of %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity of %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8) //(type, len, cap:optional)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8) //map[keyType]valueType
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45} //maps
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"]) //maps always return something
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}
	//delete(myMap2, "Adam")
	//fmt.Println(myMap2)

	for name, age := range myMap2 { //Maps
		fmt.Printf("Name: %v; Age: %v\n", name, age)
	}

	for i, v := range intArr { //Arrays and Slices
		fmt.Printf("Index: %v; Value: %v\n", i, v)
	}

	/*
		var i int = 0
		for i < 10 {
			fmt.Println(i)
			i = i + 1
		}
		//or
		i = 0
		for {
			if i >= 10 {
				break
			}
			fmt.Println(i)
			i = i + 1
		}
	*/
	for i := 0; i < 10; i++ { //dont have to declare i like above
		fmt.Println(i)
	}
}
