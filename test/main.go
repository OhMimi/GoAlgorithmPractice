package main

import (
	"fmt"
	"strconv"
)

func main() {
	test := "3+2*6-2"
	char1 := test[0:1]
	fmt.Printf("char1 type=%T char1=%v\n", char1, char1)
	char2 := []byte(test[0:1])[0]
	fmt.Printf("char2 type=%T char2=%v\n", char2, char2)
	char3, _ := strconv.Atoi(test[0:1])
	fmt.Printf("char3 type=%T char3=%v\n", char3, char3)
	char4 := int([]byte(test[0:1])[0])
	fmt.Printf("char4 type=%T char4=%v\n", char4, char4)
}
