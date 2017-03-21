package main

import (
	"fmt"
	"runtime"
)

func merge(a []int, b []int) []int {

	i, j, k := 0, 0, 0
	temp := make([]int, len(a)+len(b))
	runtime.Breakpoint()
	for (i < len(a)) && (j < len(b)) {
		if a[i] < b[j] {
			temp[k] = a[i]
			i = i + 1
		} else {
			temp[k] = b[j]
			j = j + 1
		}
		k = k + 1
	}
	for i < len(a) {
		temp[k] = a[i]
		i = i + 1
		k = k + 1
	}
	for j < len(b) {
		temp[k] = b[j]
		j = j + 1
		k = k + 1
	}

	fmt.Println("merge result", temp)
	return temp
}

func divide(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	var a1, a2 []int

	middle := int(len(a) / 2)
	fmt.Println("a=", a[:middle], "b=", a[middle:])
	a1 = divide(a[:middle])
	a2 = divide(a[middle:])
	fmt.Println("c=", a1, "d=", a2)
	result := merge(a1, a2)
	return result

}
func main() {
	fmt.Println("MergeSort program")
	a := []int{3, 1, 2, 3, 8, 1, 4, 5}
	result := divide(a)
	fmt.Println(result)
}
