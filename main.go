package main

import (
	"fmt"

	utils "app_sol/utils"
)

func main() {
	//<----Assignment 1------>

	var t = utils.Tree{}

	t.Insert(10)
	t.Insert(20)
	t.Insert(30)
	t.Insert(40)
	t.Insert(50)

	fmt.Print("In Order ")
	utils.InOrder(t.Root)
	fmt.Println()

	fmt.Print("Pre Order ")
	utils.PreOrder(t.Root)
	fmt.Println()

	fmt.Print("Post Order ")
	utils.PostOrder(t.Root)
	fmt.Println()

	// <---Assignment 2 ---->

	data := []int{6, 7, 1, 3, 8, 2, 4}
	length := len(data)
	output := utils.FindMaxlootInCircle(data, length)
	fmt.Print("<----Max loot---->", output)

}
