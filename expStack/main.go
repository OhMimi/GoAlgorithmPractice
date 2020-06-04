package main

import (
	"errors"
	"fmt"
	"strconv"
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

// 判斷一個字元是否為運算符[+, -, *, /]
func (this *Stack) isOper(val int) bool {
	// ASCII編碼 * => 42, + => 43, - => 45, / => 47
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 運算的方法
func (this *Stack) Cal(num1, num2, oper int) int {
	res := 0
	switch oper {
	case 42:
		// *
		res = num2 * num1
	case 43:
		// +
		res = num2 + num1
	case 45:
		// -
		res = num2 - num1
	case 47:
		// /
		res = num2 / num1
	default:
		fmt.Println("運算符錯誤")
	}
	return res
}

// 運算符優先級方法
// [* / => 1,  + - => 0]
func (this *Stack) Priorty(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	} else {
		fmt.Println("運算符先優級錯誤")
	}
	return res
}

func main() {
	// stack := &Stack{
	// 	MaxTop: 5,  // 表示棧最大可以存放5個數
	// 	Top:    -1, // 當棧頂為-1時,表示棧為空
	// }
	// 數字棧
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	// 符號棧
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "30+30*6-4"
	// 定義一個index,幫助我們掃描exp
	index := 0
	// 為了配合運算,定義需要的變量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {
		ch := exp[index : index+1]  // 字符串 type => string
		temp := int([]byte(ch)[0])  // 字符對應的ASCII編碼
		if operStack.isOper(temp) { // 說明是符號
			// 如果operStack 是空棧,直接入棧
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				// 棧頂運算符優先權大於當前要入棧的運算符
				// 將符號Pop出, 也從數字棧Pop倆的數,進行運算,運算後的結過再重新放回數字棧,符號再放入符號棧
				if operStack.Priorty(operStack.arr[operStack.Top]) >= operStack.Priorty(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					// 將計算結果重新入數字棧
					numStack.Push(result)
					// 將當前運算符入符號棧
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else { // 說明是數字

			// 處理多位元的思路
			// 1. 定義一個變量 keepNum string 做拼接
			keepNum += ch
			// 2. 每次要向index的後面字符測試一下,看看是否為運算符,再做處理
			// 如果已經到表達式最後,直接將keepNum
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(ch, 10, 64)
				numStack.Push(int(val))
			} else {
				// 向index後面字符測試一下是否為運算符
				// 後一位需擷取index+1~index+2
				if operStack.isOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}

		// 繼續掃描
		// 先判斷index是否已經掃描到計算表達式的最後
		if index+1 == len(exp) {
			break
		}
		index++
	}

	// 掃描表達式完畢,依次從符號棧取出符號,然後再從數字棧取出兩個數
	// 運算後結果直接入數字棧,直到符號棧為空
	for {
		if operStack.Top == -1 {
			fmt.Println("符號棧已空")
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		// 將計算結果重新入棧
		numStack.Push(result)
	}

	// 結果就是numStack最後數
	res, _ := numStack.Pop()
	fmt.Printf("表達式 %s = %v\n", exp, res)
}
