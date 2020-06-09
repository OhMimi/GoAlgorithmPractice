package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertSort(arr *[80000]int) {
	arrLength := len(arr)
	for i := 1; i < arrLength; i++ {
		// 完成第i次,給第i+1個元素找到合適的位置並插入
		interVal := arr[i]
		insertIndex := i - 1 // 欲插入的位置下標

		// 從大到小
		for insertIndex >= 0 && arr[insertIndex] < interVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex-- // 到最前面的位子後會等於-1
		}

		// 插入
		// fmt.Printf("第%d次，insertIndex+1= %d\n", i, insertIndex+1)
		if insertIndex+1 != i { // 執行完上面的迴圈會多一次--需要+1做補正
			arr[insertIndex+1] = interVal
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
	InsertSort(&arrtest)
	end := time.Now().Unix()
	fmt.Printf("選擇排序耗時=%d秒\n", end-start)

	// arr := [5]int{10, 34, 19, 100, 80}
	// arr := [10]int{89, 98, 42, 71, 95, 40, 66, 26, 6, 14}
	// fmt.Println("arr = ", arr)
	// InsertSort(&arr)
	// fmt.Println("sort arr = ", arr)
}
