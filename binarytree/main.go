package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

// 前序遍歷 (先輸出根節點,再輸出左子樹,再輸出右子樹)
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍歷 (先輸出左子樹,再輸出根節點,再輸出右子樹)
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

// 後序遍歷 (先輸出左子樹,再輸出右子樹,再輸出根節點)
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func main() {
	// 構建一個二叉樹
	root := &Hero{
		No:   1,
		Name: "hulk",
	}

	left1 := &Hero{
		No:   2,
		Name: "tommy",
	}

	right1 := &Hero{
		No:   3,
		Name: "danny",
	}

	root.Left = left1
	root.Right = right1

	left2 := &Hero{
		No:   10,
		Name: "addff",
	}

	left3 := &Hero{
		No:   12,
		Name: "sssddqq",
	}

	left1.Left = left2
	left1.Right = left3

	right2 := &Hero{
		No:   4,
		Name: "ban",
	}

	// right3 := &Hero{
	// 	No:   7,
	// 	Name: "basdasd",
	// }

	// right1.Left = right2
	right1.Right = right2

	// PreOrder(root)
	// InfixOrder(root)
	PostOrder(root)
}
