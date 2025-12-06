package main

import (
	"fmt"
	"time"
)

func main() {

	i := 5

	switch i { 
	case 1:
		println("One")
	case 2:
		println("Two")
	case 3:
		println("Three")
	case 4:
		println("Four")
	case 5:
		println("Five")
	default:
		println("Unknown Number")
	}


	// multiple cases

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}
	
}