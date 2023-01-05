package main

import "fmt"

type arr [5]int

func printArray(ar arr, n int) int {
	var ukuran int = n
	panjang := len(ar)
	if ukuran == panjang {
		return 0
	}
	fmt.Printf("%d ", ar[ukuran])
	return printArray(ar, ukuran+1)

}

func main() {
	var arr arr
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	printArray(arr, 0)
}