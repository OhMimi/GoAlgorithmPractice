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

/*
設編號1, 2, ... n 的n個人坐一圈,約定編號為k (1<=startNo<=n)
的人從1開始報數,數到countNum的那個人出列,他的下一個又從1開始報數,
數到m的那個人又出列，依次類推,直到所有人出列為止,由此生產出一個出隊編號的序列
*/

// 分析思路
// 1. 編寫一個函數,PlayGame(first *Boy, startNo int, countNum int)
// 2. 最後使用一個算法,按照要求,在環形鍊表中留下最後一個人即完成
func PlayGame(first *Boy, startNo, countNum int) {

	// 1. 空的鍊表我們單獨處理
	if first.Next == nil {
		fmt.Println("鍊表為空,沒有小孩")
		return
	}
	// 留一個,判斷 startNo <= 小孩的總數
	// 2. 需要定義一個輔助節點,幫助我們刪除小孩
	tail := first
	// 3. 讓tail執行環形鍊表的最後一個小孩,非常重要
	// 因為tail在刪除小孩時會需要用到
	for {
		if tail.Next == first { // 代表已經到最後一個小孩了
			break
		}
		tail = tail.Next
	}
	// 4. 讓first移動到startNo [後面刪除小孩,就以first為準]
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	// 5. 開始數 countNum, 然後刪除first指向的小孩
	for {
		// 開始數countNum-1次
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩編號為 %d 出列 ->\n", first.No)
		// 刪除first指向的節點
		first = first.Next
		tail.Next = first

		// 判斷如果tail == first 圈中只剩下一個
		if tail == first {
			break
		}
	}
	fmt.Printf("最後小孩編號為 %d 出列 ->\n", first.No)
}

func main() {
	first := AddBoy(5)
	// 顯示
	ShowBoy(first)

	PlayGame(first, 2, 3)
}
