package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 選擇排序(select sort)大到小
func SelectSort(arr *[80000]int) {
	// (*arr)[1] 等價於 arr[1]
	// arr[1] = 600

	arrLength := len(arr)
	// 最後一個不需比較，因為前面都比完了
	for j := 0; j < arrLength-1; j++ {
		// 1. 先完成將第一個最大值 和 arr[j] 交換=> 先易後難
		// 假設 arr[0] 是最大值
		max := arr[j]
		maxIndex := j

		// 2. 遍歷後面的 1 ~ len(arr)-1 比較
		for i := j; i < arrLength; i++ {
			if max < arr[i] { // 找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		// 交換
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}
	// fmt.Println("第1次", *arr)
}

func main() {
	var arrtest [80000]int
	for i := 0; i < 80000; i++ {
		arrtest[i] = rand.Intn(900000)
	}

	fmt.Println("開始排序")
	start := time.Now().Unix()
	SelectSort(&arrtest)
	end := time.Now().Unix()
	fmt.Printf("選擇排序耗時=%d秒\n", end-start)

	// 定義一個數組
	// arr := [5]int{10, 34, 19, 100, 80}
	// fmt.Println("arr = ", arr)
	// SelectSort(&arr)
	// fmt.Println("sort arr = ", arr)
}
