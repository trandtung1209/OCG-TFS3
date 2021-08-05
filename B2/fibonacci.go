package main

import "fmt"

func fib(i int) int {
	if i < 0 {
		return -1
	} else if i==0||i==1 {
		return i
	} else {
		return fib(i-1) + fib(i-2)
	}
}

func main() {
	fmt.Print("Nhap so luong can in:")
	var n int
	fmt.Scan(&n)
	for i:=0;i<n;i++ {
		fmt.Println(fib(i))
	}

}