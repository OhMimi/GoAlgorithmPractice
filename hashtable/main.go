package main

import (
	"fmt"
	"os"
)

// 定義Emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 輸出雇員訊息
func (this *Emp) ShowEmp() {
	fmt.Printf("鍊表%d 找到該雇員id=%d 名稱=%s\n", this.Id%7, this.Id, this.Name)
}

// 定義EmpLink
// 這邊的不帶表頭,及第一個節點就存放雇員
type EmpLink struct {
	Head *Emp
}

// 1.添加員工的方法,保證添加時,編號從小到大
func (this *EmpLink) Insert(emp *Emp) {
	// pre 在 cur 前面
	cur := this.Head   // 輔助指針
	var pre *Emp = nil // 這也是輔助指針

	// 如果當前的EmpLink就是一個空鍊表
	if cur == nil {
		this.Head = emp // 完成
		return
	} else {
		// 如果不是一個空鍊表,給emp找到對應的位置並插入
		// 思路是 讓cur跟emp比較, 讓pre始終保持在cur前方
		for {
			if cur != nil {
				// 比較
				if cur.Id >= emp.Id {
					// 找到位置了
					break
				}
				pre = cur // 保證同步
				cur = cur.Next
			} else {
				// 代表沒有比emp大的Id,需放在最後
				break
			}
		}

		// 退出時,看是否將emp添加到鍊表的最後
		// if cur == nil {
		pre.Next = emp
		emp.Next = cur
		// }
	}
}

// 顯示鍊表的信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("鍊表%d為空\n", no)
		return
	}

	//變量當前的鍊表,並顯示數據
	cur := this.Head // 輔助的指針
	for {
		if cur != nil {
			fmt.Printf("鍊表%d  雇員id=%d  名字=%s  ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else { // cur == nil 代表已遍歷至鍊表尾部
			break
		}
	}
	fmt.Println() // 換行處理
}

// 根據id查找對應的雇員,如果沒有就返回nil
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		// 已找到符合id目標 或是 已找到鍊表最後但沒找到符合 就中斷遍歷
		if (cur != nil && cur.Id == id) || cur == nil {
			break
		} else {
			cur = cur.Next
		}
	}

	return cur
}

// 定義hashtable, 含有一個鍊表陣列
type HashTable struct {
	LinkArr [7]EmpLink
}

// HashTable Insert雇員的方法
func (this *HashTable) Insert(emp *Emp) {
	// 使用散列函數,確定將該雇員添加到哪個鍊表
	linkNo := this.HashFun(emp.Id)
	// 使用對應的鍊表添加
	this.LinkArr[linkNo].Insert(emp)
}

// 顯示hashtable的所有雇員
func (this *HashTable) ShowAll() {
	var arrLen = len(this.LinkArr)
	for i := 0; i < arrLen; i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

// 編寫一個散列函數
func (this *HashTable) HashFun(id int) int {
	return id % 7 // 得到一個值,就是對應的鍊表下標
}

// 編寫一個方法完成查找
func (this *HashTable) FindById(id int) *Emp {
	// 使用散列函數,確定將該雇員在哪個鍊表
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

func main() {
	var hashTable HashTable
	key := ""
	id := 0
	name := ""
	for {
		fmt.Println("====雇員系統菜單====")
		fmt.Println("input 添加雇員")
		fmt.Println("show  顯示雇員")
		fmt.Println("find  查找雇員")
		fmt.Println("exit  退出系統")
		fmt.Println("請輸入你的選擇:")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("輸入雇員的ID:")
			fmt.Scanln(&id)
			fmt.Println("輸入雇員的name:")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.ShowAll()
		case "find":
			fmt.Println("輸入雇員的ID:")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇員不存在\n", id)
			} else {
				emp.ShowEmp()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("輸入錯誤")
		}
	}
}
