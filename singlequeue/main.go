package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一個結構體管理隊列
type Queue struct {
	MaxSize int
	Array   [5]int // 陣列 => 模擬隊列
	Front   int    // 表示指向隊列隊首  -1
	Rear    int    // 表示指向隊列隊尾  -1
}

// 添加數據到對列
func (q *Queue) AddQueue(val int) (err error) {

	// 先判斷隊列是否已滿
	if q.Rear == q.MaxSize-1 { // 重要的提示; Rear是隊列尾部(包含最後的元素)
		return errors.New("Queue is full")
	}
	q.Rear++ // Rear後移
	q.Array[q.Rear] = val

	return
}

// 從隊列中取出數據
func (q *Queue) GetQueue() (val int, err error) {
	// 先判斷隊列是否為空
	if q.Front == q.Rear { // 隊列已空
		return -1, errors.New("Queue is empty")
	}
	q.Front++
	val = q.Array[q.Front]
	return
}

// 顯示隊列，找到隊首，然後遍歷到隊尾
func (q *Queue) ShowQueue() {
	// Front不包含隊首的元素
	fmt.Println("Queue的內容:")
	for i := q.Front + 1; i <= q.Rear; i++ {
		fmt.Printf("array[%d] = %d\n", i, q.Array[i])
	}
}

// 主函數測試
func main() {
	queue := &Queue{
		MaxSize: 5,
		Front:   -1,
		Rear:    -1,
	}

	var key string // 選項
	var val int    // 欲加入的數值

	for {
		fmt.Println("1. 輸入add 表示添加數據到隊列")
		fmt.Println("2. 輸入get 表示隊列中獲取數據")
		fmt.Println("3. 輸入show 表示顯示隊列")
		fmt.Println("4. 輸入exit 表示退出")
		fmt.Scanf("%s\n", &key)

		switch key {
		case "add":
			fmt.Println()
			fmt.Println("請輸入要加入的值:")
			fmt.Scanf("%d\n", &val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Printf("queue.AddQueue err = %v\n", err.Error())
			} else {
				fmt.Println("加入隊列成功")
			}
		case "get":
			fmt.Println()
			fmt.Println("從隊列中取出的值:")
			getValue, err := queue.GetQueue()
			if err != nil {
				fmt.Printf("queue.GetQueue err = %v\n", err.Error())
			} else {
				fmt.Printf("隊列取出的值為: %d\n", getValue)
			}
		case "show":
			fmt.Println()
			fmt.Println("目前隊列中的數據:")
			queue.ShowQueue()
		case "exit":
			os.Exit(0)

		default:
			fmt.Println("輸入錯誤，請重新輸入")
		}
		fmt.Println()
	}
}
