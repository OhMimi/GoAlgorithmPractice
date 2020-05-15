package main

import "fmt"

type ValNode struct {
	Row int
	Col int
	Val int
}

func main() {
	// 1. 先創建一個原始2維陣列
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 藍子

	// 2. 輸出看看原始陣列
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 3. 轉換成稀疏陣列
	// 方式
	// (1) 遍歷 chesssMap, 如果發現有一個元素值不為0, 創建一個node結構體
	// (2) 將其放入到對應的切片即可
	var sparseArr []*ValNode

	// 標準的稀疏陣列需要含有一個紀錄2維陣列的規模(行、列、默認值)
	// 創建一個 ValNode 值節點
	valNode := &ValNode{
		Row: 11,
		Col: 11,
		Val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v1 := range chessMap {
		for j, v2 := range v1 {
			if v2 != 0 {
				// 創建一個 ValNode 值節點
				valNode = &ValNode{
					Row: i,
					Col: j,
					Val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 輸出稀疏陣列
	fmt.Println("當前的稀疏陣列:")
	for i, valNode := range sparseArr {
		fmt.Printf("%d:\t%d\t%d\t%d\n", i, valNode.Row, valNode.Col, valNode.Val)
	}

	// 將此稀疏陣列儲存

	// 如何恢復回原本的陣列
	// 1. 打開儲存文件 => 恢復原始數組

	// 2. 這裡使用稀疏陣列恢復

	// 先創建一個原始陣列
	var chessMap2 [11][11]int
	// 遍歷 sparseArr [遍歷文件每一行]
	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.Row][valNode.Col] = valNode.Val
		}
	}

	fmt.Println("回復出的原始陣列:")
	// 確認chessMap2是不是恢復了
	for _, v1 := range chessMap2 {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
