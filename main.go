package main

import "fmt"

func main() {
	arr := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	fmt.Println(selectionSort(arr))
}

func selectionSort(array []int) []int {
	currentIndex := 0
	smallestIndex := 0
	for currentIndex < len(array)-1 {
		smallestIndex = currentIndex
		for i := currentIndex + 1; i < len(array); i++ {
			if array[smallestIndex] > array[i] {
				smallestIndex = i
			}
		}
		swap(currentIndex, smallestIndex, array)
		currentIndex++
	}
	return array
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}
