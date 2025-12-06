package main

func main() {

	i:=1
	for i<=5 {
		println(i)
		i++
	}

	for j :=range 5 {
		if j == 2 {
			continue
		}
		println(j)
	}
}