package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Println(indexed)
	fmt.Printf("%v, %T\n", indexed, indexed) //%T is for type

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("The length of 'myString' is %v\n", len(myString))

	var myRune = 'a' //runes can be declared with ''
	fmt.Printf("myRune = %v\n", myRune)

	var strSlice = []string{"i", "d", "e", " ", "g", "a", "s", "!"}
	/*
		    var concatStr string
			for i := range strSlice {
				concatStr += strSlice[i] //this is inneficient because everytime a letter is added a new arrays is created
			}
			fmt.Printf("Concatenated string: %v\n", concatStr)
		    //concatStr[0] = "a" like arrays strings cant be changed
	*/
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i]) //this solves the efficiency problem above
	}
	var concatStr = strBuilder.String()
	fmt.Printf("Concatenated string: %v\n", concatStr)
}
