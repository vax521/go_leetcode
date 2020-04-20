package main

import "fmt"

func partition(aList []int, left, right int) int {
	baseNum := aList[left]
	for right > left {
		for aList[right] >= baseNum && right > left {
			right--
		}
		aList[left] = aList[right]
		for aList[left] <= baseNum && right > left {
			left++
		}
		aList[right] = aList[left]
	}
	aList[right] = baseNum
	fmt.Println(aList)
	return right
}

func quickSort(input []int, left, right int) {
	if left > right {
		return
	}
	index := partition(input, left, right)
	quickSort(input, left, index-1)
	quickSort(input, index+1, right)
}

func main() {
	var intArray = []int{9, 8, 7, 6, 5, 0, 4, 3, 2, 1}
	quickSort(intArray, 0, len(intArray)-1)
	fmt.Println(intArray)
}
