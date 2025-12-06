package main 

import "fmt"

func add( a int , b int) int {
	return a + b
}

func getLang() (string, string) {
	return "Golang", "Java"
}

func main () {

	res := add(3, 4)
	fmt.Println("Result:", res)

	fmt.Println(getLang())

}