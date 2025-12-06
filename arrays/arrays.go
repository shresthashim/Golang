package main


func main () {

	// var nums [5]int = [5]int{1,2,3,4,5}

	nums := [5]int{1,2,3,4,5}

	for i:= range nums {
		println(nums[i])
	}

	nums1 := [2][2]int{{1,2}, {3,4}}

	for i:= range nums1 {
		for j:= range nums1[i] {
			println(nums1[i][j])
		}
	}
}
