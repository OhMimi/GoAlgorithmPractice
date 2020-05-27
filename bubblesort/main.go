package main

import "fmt"

func BubbleSort(arr *[5]int) {
	arrLength := len(arr)
	for i := 0; i < arrLength-1; i++ {
		for j := 0; j < arrLength-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := [5]int{10, 34, 19, 100, 80}
	fmt.Println("arr = ", arr)
	BubbleSort(&arr)
	fmt.Println("sort arr = ", arr)
}
