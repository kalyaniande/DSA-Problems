package main

import (
	"fmt"
)

// const abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// func toCharStrConst(i int) string {
// 	return abc[i-1 : i]
// }

func findColumnNumber(column_title string) int {
	response := 0
	for i := 0; i < len(column_title); i++ {

		char_num := column_title[i] - 64
		response *= 26
		response += int(char_num)
		fmt.Println(i, ":", char_num, ":", response)
	}

	return response
}

func main() {
	//fmt.Println(toCharStrConst(1))
	column_title := "CDA"
	fmt.Println(findColumnNumber(column_title))

}
