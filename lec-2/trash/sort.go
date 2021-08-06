package main

import "fmt"

func bubble(arr []int) []int {
	n:= len(arr)
	for i:=0;i<n;i++ {
		haveSwap := true
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j] = arr[j] ^ arr[j+1]
				arr[j+1] = arr[j] ^ arr[j+1]
				arr[j] = arr[j] ^ arr[j+1]
				haveSwap = false
			}
		}
		if haveSwap {
			break
		}
	}
	return arr
}

func main() {
	numbers:=[]int{8,4,2,7,3,6,1,5}
	bubble(numbers)
	fmt.Println(numbers)
}