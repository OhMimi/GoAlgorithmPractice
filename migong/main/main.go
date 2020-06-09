package main

import "fmt"

// 編寫一個函數完成老鼠找路
// myMap: 地圖,須保證是同一個地圖,因此用引用
// i,j 表示為地圖上的座標點, i: row j: col
func SetWay(myMap *[8][7]int, i, j int) bool {
	// 分析出什麼情況下會找到出路
	if myMap[6][5] == 2 {
		return true
	} else {
		// 說明要繼續找
		if myMap[i][j] == 0 { // 說明可被探測,為0

			// 假設這個點是可以通的,但是需要探測 下右上左
			myMap[i][j] = 2
			if SetWay(myMap, i+1, j) { // 下
				return true
			} else if SetWay(myMap, i, j+1) { // 右
				return true
			} else if SetWay(myMap, i-1, j) { // 上
				return true
			} else if SetWay(myMap, i, j-1) { // 左
				return true
			} else { // 死路
				myMap[i][j] = 3
				return false
			}
		} else {
			// 說明無法被探測,為1 是牆
			return false
		}
	}
}

func PrintMap(myMap *[8][7]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf("%d ", myMap[i][j])
			// fmt.Printf("myMap[%d][%d] = %d ", i, j, myMap[i][j])
		}
		fmt.Println("")
	}
}

func main() {
	// 先創建一個二維陣列,模擬迷宮
	// 規則
	// 1. 如果元素的值為1,就是牆
	// 2. 如果元素的值為0,就是沒有走過的點
	// 3. 如果元素的值為2,就是通路
	// 4. 如果元素的值為3,就是已經走過的點,但是走不通
	var myMap [8][7]int

	// 先把地圖上的最上以及最下設置為1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	// 先把地圖上的最左以及最右設置為1
	for i := 0; i < 7; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}

	myMap[3][1] = 1
	myMap[3][2] = 1
	myMap[1][2] = 1
	myMap[2][2] = 1

	// 輸出地圖
	PrintMap(&myMap)

	// 使用測試
	SetWay(&myMap, 1, 1)
	fmt.Println("探測完畢...")
	PrintMap(&myMap)

}
