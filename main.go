package main

import (
	"fmt"
	"golang101/gotype"
	"golang101/mysyntax"
)

func main() {
	//START Function
	mysyntax.Greeting("Jame", "Inw")
	fmt.Println(mysyntax.GreetingWithAge("Jame", 25))
	fmt.Println(mysyntax.GreetingWithAgeAndDrink("Jame", 25, "Cola"))
	// END Function

	// START IF ELSE
	fmt.Println(mysyntax.IsOdd(3))
	// END IF ELSE

	// START Loop
	fmt.Println(mysyntax.SumOffFirst(10))
	// END Loop

	// START Pointer
	n := 4
	gotype.Double(&n)

	s := "Jame"
	gotype.AppendGreeting(&s)
	// END pointer

	// Array and Slice
	gotype.Abc()
	gotype.Efg()
	gotype.Cde()
}
