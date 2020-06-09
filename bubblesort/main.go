package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(arr *[80000]int) {
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
	var arrtest [80000]int
	for i := 0; i < 80000; i++ {
		arrtest[i] = rand.Intn(900000)
	}

	fmt.Println("開始排序")
	start := time.Now().Unix()
	BubbleSort(&arrtest)
	end := time.Now().Unix()
	fmt.Printf("選擇排序耗時=%d秒\n", end-start)

	// arr := [5]int{10, 34, 19, 100, 80}
	// fmt.Println("arr = ", arr)
	// BubbleSort(&arr)
	// fmt.Println("sort arr = ", arr)
}
