package main

import (
	"errors"
	"fmt"
)

// 使用陣列來模擬一個棧的使用
type Stack struct {
	MaxTop int    // 表示棧最大可以存放的個數
	Top    int    // 表示棧頂,因為棧頂固定,因此直接使用Top
	arr    [5]int // 陣列模擬棧
}

func (this *Stack) Push(val int) (err error) {
	// 先判斷棧是否已滿
	if this.Top == this.MaxTop-1 {
		fmt.Println("棧已滿(stack full)")
		return errors.New("stack full")
	}

	this.Top++
	// 放入數據
	this.arr[this.Top] = val
	return
}

// 遍歷棧,注意需要從棧頂開始遍歷
func (this *Stack) List() {
	// 先判斷棧是否為空
	if this.Top == -1 {
		fmt.Println("棧是空的(stack empty)")
		return
	}
	// curTop := this.Top
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d\n", i, this.arr[i])
	}
}

// 出棧
func (this *Stack) Pop() (val int, err error) {
	// 先判斷棧是否為空
	if this.Top == -1 {
		fmt.Println("棧是空的(stack empty)")
		return 0, errors.New("stack empty")
	}
	// 先取值,再 this.Top--
	val = this.arr[this.Top]
	this.Top--

	return val, nil
}

func main() {
	stack := &Stack{
		MaxTop: 5,  // 表示棧最大可以存放5個數
		Top:    -1, // 當棧頂為-1時,表示棧為空
	}

	// 入棧
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)

	// 顯示
	stack.List()

	// 出棧
	val, _ := stack.Pop()
	fmt.Println("出棧的值 = ", val)
	val, _ = stack.Pop()
	fmt.Println("出棧的值 = ", val)
	val, _ = stack.Pop()
	fmt.Println("出棧的值 = ", val)
	// val, _ = stack.Pop()
	// fmt.Println("出棧的值 = ", val)
	// val, _ = stack.Pop()
	// fmt.Println("出棧的值 = ", val)
	// val, _ = stack.Pop()
	// fmt.Println("出棧的值 = ", val)

	// 顯示
	stack.List()
}
