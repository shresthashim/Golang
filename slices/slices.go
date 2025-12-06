package main 

import "fmt"

func main() {

	var nums = make([]int, 2, 5)

	fmt.Println("Length:", len(nums))
	fmt.Println("Capacity:", cap(nums))

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)

	fmt.Println("After appending elements:")
	fmt.Println("Length:", len(nums))
	fmt.Println("Capacity:", cap(nums))

	for i := range nums {
		fmt.Println(nums[i])
	}	
}
