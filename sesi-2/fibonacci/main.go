package main

import (
    "fmt"
)

func main() {
    fmt.Println(fib(0)) 
	fmt.Println(fib(1)) 
	fmt.Println(fib(2)) 
	fmt.Println(fib(3)) 
	fmt.Println(fib(4)) 

	// fmt.Println(fib(0)) --> expected cetak angka 0
	// fmt.Println(fib(1)) --> expected cetak angka 1
	// fmt.Println(fib(2)) --> expected cetak angka 1
	// fmt.Println(fib(3)) --> expected cetak angka 2
	// fmt.Println(fib(4)) 
	
	// --> expected cetak angka 3

}

// 

// Fib adalah function untuk menghitung angka fibonacci ke n
// 1, 1, 2, 3, 5, 8
// fib(0) = 0
// fib(1) = 1
// fib(2) = fib(1) + fib(0) = 1 + 0 = 1
// fib(3) = fib(2) + fib(1) = 1 + 1 = 2
// fib(4) = fib(3) + fib(2) = 2 + 1 = 3
func fib(n int) int {
	if n < 2{
		return n
	}
	
	return fib(n-1)+fib(n-2)
}