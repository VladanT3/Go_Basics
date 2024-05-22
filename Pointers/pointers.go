package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32 = 5
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)
	*p = 10
	fmt.Printf("The value p points to is: %v\n", *p)
	p = &i
	fmt.Printf("The value p points to is: %v\n", *p)
	*p = 15
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	fmt.Println("")

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice, sliceCopy)

	fmt.Println("")

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of thing1 is: %p\n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("The result is: %v\n", result)
	fmt.Printf("The values of thing1 are: %v\n", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of thing2 is: %p\n", thing2)
	for i := range *thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
