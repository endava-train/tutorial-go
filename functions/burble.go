package main

import "math/rand"

func burbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	randomNumber := make([]int, 10)
	for i := 0; i < len(randomNumber); i++ {
		randomNumber[i] = rand.Intn(100)
	}
	for _, element := range randomNumber {
		print(element, " ")
	}
	println()
	arr := burbleSort(randomNumber)
	for _, element := range arr {
		print(element, " ")
	}
	println()

}
