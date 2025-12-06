package main

import "fmt"

func main() { 

	sample_map := make (map[string]int)

	sample_map["one"] = 1
	sample_map["two"] = 2
	sample_map["three"] = 3

	for key := range sample_map {
		fmt.Println("Key:", key, "Value:", sample_map[key])
	}
}
