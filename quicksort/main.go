package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 快速排序
// left: 表示陣列最左邊的下標
// right: 表示陣列最右邊的下標
// arr: 表示要排序的陣列
func QuickSort(left, right int, arr *[80000]int) {
	l := left
	r := right
	// pivot 陣列的中軸，從陣列中間開始當中軸通常較快排序完成
	pivot := arr[(left+right)/2]

	// for循環的目標是將比pivot小的數放到左邊
	// 比pivot大的數放到右邊
	for l < r {
		// 先從pivot的左邊找到大於等於piovt的值
		for arr[l] < pivot {
			l++
		}
		// 從pivot的右邊找到小於等於piovt的值
		for arr[r] > pivot {
			r--
		}
		// l >= r 說明本次分解任務完成, break
		if l >= r {
			break
		}
		// 交換
		arr[l], arr[r] = arr[r], arr[l]
		// 優化
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}

	// 如果l == r, 在移動到下一位比較
	if l == r {
		l++
		r--
	}

	// 向左遞歸
	if left < r {
		QuickSort(left, r, arr)
	}
	// 向右遞歸
	if right > l {
		QuickSort(l, right, arr)
	}
}

func main() {
	var arrtest [80000]int
	for i := 0; i < 80000; i++ {
		arrtest[i] = rand.Intn(900000)
	}

	fmt.Println("開始排序")
	start := time.Now().Unix()
	QuickSort(0, len(arrtest)-1, &arrtest)
	end := time.Now().Unix()
	fmt.Printf("選擇排序耗時=%d秒\n", end-start)

	// arr := [6]int{-9, 78, 0, 23, -567, 70}
	// fmt.Println("arr = ", arr)
	// QuickSort(0, len(arr)-1, &arr)
	// fmt.Println("sort arr = ", arr)
}
