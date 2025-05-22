package main

func CalculateSum(slice []int, ch chan int) {
	var sum int
	for _, value := range slice {
		sum += value
	}
	ch <- sum
}

func SplitSlice(slice []int) ([]int, []int) {
	mid := len(slice) / 2
	return slice[:mid], slice[mid:]
}
