package main

import "fmt"

// 定義貓的結構體
type CatNode struct {
	no   int // 貓貓的編號
	name string
	next *CatNode
}

func InsertCatNode(head, newCatNode *CatNode) {
	// 先判斷是不是添加第一隻貓
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Printf("%v這隻貓已加入到環形鍊表\n", newCatNode)
		return
	}

	// 定義一個臨時的變量
	temp := head
	for {
		if temp.next == head { // 找到最後一個貓貓
			break
		}
		temp = temp.next
	}
	// 加入到環形鍊表中
	temp.next = newCatNode
	newCatNode.next = head
}

func DeleteNode(head *CatNode, id int) *CatNode {
	temp := head   // 用來比較,指向鍊表的一開始
	helper := head // 用來刪除,指向鍊表的最後

	// 空鍊表
	if temp.next == nil {
		fmt.Println("這是一個空的環形鍊表,無法刪除")
		return head
	}

	// 如果只有1個節點
	if temp.next == head { // 只有1個節點
		temp.next = nil
		return head
	}

	// 將helper訂位到鍊表最後
	for {
		if helper.next == head { // 說明已找到最後一個
			break
		}
		helper = helper.next
	}

	// 如果是2個以上節點
	flag := true // 有到最後一個節點
	for {
		if temp.next == head { // 說明已找到最後一個,最後一個尚未比較(繞了一圈)
			break
		}

		if temp.no == id { // 找到id,直接刪除
			if temp == head { // 說明刪除的是頭節點
				head = head.next
			}
			helper.next = temp.next
			fmt.Printf("貓貓=%d\n", id)
			flag = false
			break
		}
		temp = temp.next     // 用來比較
		helper = helper.next // 用來刪除
	}

	// 這裡還要比較後一次
	if flag { // 如果flag為true,則上面沒有刪除
		if temp.no == id { // 找到id,直接刪除
			helper.next = temp.next
			fmt.Printf("貓貓=%d\n", id)
		} else {
			fmt.Printf("找不到id=%d貓貓\n", id)
		}
	}

	return head
}

func ListCircleLink(head *CatNode) {
	fmt.Println("環形鍊表的情況如下:")
	temp := head

	if temp.next == nil {
		fmt.Println("這是一個空的環形鍊表...")
		return
	}

	for {
		fmt.Printf("貓的訊息為 = [id:%d name=%s] -> \n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {
	// 這裡我們初始化一個環形鍊表的頭節點
	head := &CatNode{}

	// 創建一隻新貓
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "hulk",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tony",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)

	head = DeleteNode(head, 2)

	ListCircleLink(head)
}
