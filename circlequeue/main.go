package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一個結構體管理環形隊列
type CircleQueue struct {
	MaxSize int    // 4
	Array   [4]int // 陣列
	Head    int    // 指向隊列的隊首 0
	Tail    int    // 指向隊列的隊尾 0
}

// 添加數據到對列
func (cq *CircleQueue) Push(val int) (err error) {
	if cq.IsFull() {
		return errors.New("Queue is Full")
	}

	// 分析出cq.Tail在隊列的尾部,但是包含最後的元素
	cq.Array[cq.Tail] = val // 把值給尾部
	cq.Tail = (cq.Tail + 1) % cq.MaxSize

	return
}

// 從隊列中取出數據
func (cq *CircleQueue) Pop() (val int, err error) {
	if cq.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	// 取出,Head是指向隊首,並且包含隊首元素
	val = cq.Array[cq.Head]
	cq.Head = (cq.Head + 1) % cq.MaxSize
	return
}

// 顯示隊列
func (cq *CircleQueue) ListQueue() {
	fmt.Println("環形隊列情況如下:")
	size := cq.Size()
	if size == 0 {
		fmt.Println("隊列為空")
	}

	// 設計一個輔助的變量,指向Head
	tempHead := cq.Head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", tempHead, cq.Array[tempHead])
		tempHead = (tempHead + 1) % cq.MaxSize
	}
	fmt.Println()
}

// 判斷環形隊列為滿
func (cq *CircleQueue) IsFull() bool {
	return (cq.Tail+1)%cq.MaxSize == cq.Head
}

// 判斷環形隊列為空
func (cq *CircleQueue) IsEmpty() bool {
	return cq.Tail == cq.Head
}

// 顯示出環形隊列有多少個元素
func (cq *CircleQueue) Size() int {
	// 這是一個關鍵的算法
	return (cq.Tail + cq.MaxSize - cq.Head) % cq.MaxSize
}

func main() {
	queue := &CircleQueue{
		MaxSize: 4,
		Head:    0,
		Tail:    0,
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
			err := queue.Push(val)
			if err != nil {
				fmt.Printf("queue.AddQueue err = %v\n", err.Error())
			} else {
				fmt.Println("加入隊列成功")
			}
		case "get":
			fmt.Println()
			fmt.Println("從隊列中取出的值:")
			getValue, err := queue.Pop()
			if err != nil {
				fmt.Printf("queue.GetQueue err = %v\n", err.Error())
			} else {
				fmt.Printf("隊列取出的值為: %d\n", getValue)
			}
		case "show":
			fmt.Println()
			fmt.Println("目前隊列中的數據:")
			queue.ListQueue()
		case "exit":
			os.Exit(0)

		default:
			fmt.Println("輸入錯誤，請重新輸入")
		}
		fmt.Println()
	}
}
