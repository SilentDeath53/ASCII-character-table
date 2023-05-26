package main

import "fmt"

func main() {
	fmt.Println("ASCII Character Table")
	fmt.Println("---------------------")
	fmt.Println("Decimal\tASCII")

	for i := 0; i <= 127; i++ {
		fmt.Printf("%d\t%c\n", i, i)
	}
}
