package main

import "fmt"

// 小孩的結構體
type Boy struct {
	No   int  // 編號
	Next *Boy // 指向下一個小孩的指針
}

// 編寫一個函數,構成單向的環形鍊表
// params num : 小孩的個數
// return *Boy : 返回該環形鍊表第一個小孩的指針
func AddBoy(num int) *Boy {

	first := &Boy{}  // 空節點
	curBoy := &Boy{} // 空節點

	// 判斷
	if num < 1 {
		fmt.Println("num的值無法小於1")
		return first
	}
	// 循環的構建這個環形鍊表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}

		// 分析出構成一個循環鍊表,需要一個輔助指針
		// 1. 因為第一個小孩比較特殊
		if i == 1 { // 第一個小孩
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first // 構成環形鍊表
		}
	}

	return first
}

func ShowBoy(first *Boy) {
	// 處理一下如果環形鍊表為空
	if first.Next == nil {
		fmt.Println("鍊表為空,沒有小孩")
		return
	}

	// 先定義創建指針,幫助遍歷
	// 至少有一個小孩
	curBoy := first
	for {
		fmt.Printf("小孩編號=%d ->\n", curBoy.No)
		// 退出的條件
		if curBoy.Next == first {
			break
		}

		curBoy = curBoy.Next
	}
}

func main() {
	first := AddBoy(5)
	// 顯示
	ShowBoy(first)
}
