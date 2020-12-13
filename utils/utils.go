package utils

import (
	"fmt"
	"math"
)

type node struct {
	key   int
	left  *node
	right *node
}

type Tree struct {
	Root *node
}

/* <------Assignment 1------> */
func (t *Tree) Insert(data int) {
	if t.Root == nil {
		t.Root = &node{key: data}
	} else {
		t.Root.Insert(data)
	}
}

func (n *node) Insert(data int) {
	if data < n.key {
		if n.left == nil {
			n.left = &node{key: data}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &node{key: data}
		} else {
			n.right.Insert(data)
		}
	}
}

func InOrder(n *node) {
	if n == nil {
		return
	} else {
		InOrder(n.left)
		fmt.Print(n.key)
		InOrder(n.right)
	}
}

func PreOrder(n *node) {
	if n == nil {
		return
	} else {
		fmt.Print(n.key)
		PreOrder(n.left)
		PreOrder(n.right)
	}
}

func PostOrder(n *node) {
	if n == nil {
		return
	} else {
		PostOrder(n.left)
		PostOrder(n.right)
		fmt.Print(n.key)
	}
}

/* <------Assignment 2------> */

func FindMaxlootInCircle(data []int, length int) int {
	if length == 0 {
		return 0
	}
	if length == 1 {
		return data[0]
	}

	if length == 2 {
		return int(math.Max(float64(data[0]), float64(data[1])))
	}
	firtIncluded := findMaxlootCalculation(data, 0, length-1)
	firstNotIncluded := findMaxlootCalculation(data, 1, length)
	return int(math.Max(float64(firtIncluded), float64(firstNotIncluded)))
}

func findMaxlootCalculation(data []int, start int, end int) int {
	preMax := data[start]
	curMax := int(math.Max(float64(preMax), float64(data[start+1])))
	for i := start + 2; i < end; i++ {
		temp := curMax
		curMax = int(math.Max(float64(data[i]+preMax), float64(curMax)))
		preMax = temp
	}
	return curMax
}
