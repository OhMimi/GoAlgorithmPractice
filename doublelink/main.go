package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode // 這裡表示指向的前一個節點
	next     *HeroNode // 這裡表示指向的下一個節點
}

// 給練表插入一個節點
// 編寫第一種插入方式,在單鍊表的最後加入
func InsertHeroNode(head, newHeroNode *HeroNode) {
	// 思路
	// 1. 先找到該鍊表的最後的節點
	// 2. 創建一個輔助節點
	temp := head
	for {
		if temp.next == nil { // 表示找到最後
			break
		} else {
			temp = temp.next // 讓temp不斷地指向下一個節點
		}
	}

	// 3. 將newHeroNode加入到鍊表的最後
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

// 給練表插入一個節點
// 編寫第二種插入方式,在單鍊表的最後加入
func InsertHeroNode2(head, newHeroNode *HeroNode) {
	// 思路
	// 1. 先找到該鍊表的最後的節點
	// 2. 創建一個輔助節點
	temp := head
	flag := false
	for {
		if temp.next == nil { // 表示找到最後
			break
		} else if temp.next.no > newHeroNode.no {
			// 說明newHeroNode此時需插入到temp之後
			break
		} else if temp.next.no == newHeroNode.no {
			// 說明鍊表中已有no號存在，無法插入
			flag = true
			break
		} else {
			temp = temp.next
		}
	}

	if flag {
		fmt.Printf("no號已存在不允許插入, no = %d\n", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}

// 刪除鍊表中一個節點
func DelHeroNode(head *HeroNode, id int) {
	// 1. 先找到該鍊表的最後的節點
	temp := head
	flag := false
	// 找到要刪除的節點
	for {
		if temp.next == nil { // 表示找到最後
			break
		} else if temp.next.no == id {
			// 說明已經找到要刪除的節點
			flag = true
			break
		} else {
			temp = temp.next
		}
	}

	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("要刪除的ID不存在")
	}
}

// 顯示鍊表的所有節點信息
func ListHeroNode(head *HeroNode) {
	// 1. 創建一個輔助節點
	temp := head

	// 先判斷該鍊表是否為空鍊表
	if temp.next == nil {
		fmt.Println("空鍊表無法顯示")
		return
	}

	// 2. 遍歷這個鍊表
	for {
		fmt.Printf("[%d , %s , %s] ==> ", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		// 判斷是否為鍊表最後
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

// 倒序顯示鍊表的所有節點信息
func ListHeroNode2(head *HeroNode) {
	// 1. 創建一個輔助節點
	temp := head

	// 先判斷該鍊表是否為空鍊表
	if temp.next == nil {
		fmt.Println("空鍊表無法顯示")
		return
	}

	// 2. 讓temp定位到雙向鍊表的最後一節點
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// 2. 遍歷這個鍊表
	for {
		fmt.Printf("[%d , %s , %s] ==> ", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		// 判斷是否為鍊表頭
		if temp.pre == nil {
			break
		}
	}
	fmt.Println()
}

func main() {

	// 1. 先創建一個頭節點,
	head := &HeroNode{}

	// 2. 創建一個新的節點
	hero1 := &HeroNode{
		no:       1,
		name:     "hulk",
		nickname: "popopo",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "tom",
		nickname: "tommy",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "tony",
		nickname: "tony123",
	}

	hero4 := &HeroNode{
		no:       4,
		name:     "appaa",
		nickname: "appaa123",
	}

	// 3. 加入
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	InsertHeroNode(head, hero4)

	// 4. 顯示
	fmt.Println("正序打印:")
	ListHeroNode(head)
	fmt.Println("逆序打印:")
	ListHeroNode2(head)

	// // 5. 刪除
	// DelHeroNode(head, 2)

	// // 6. 顯示
	// ListHeroNode(head)
}
