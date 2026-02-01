package main

import "fmt"

func main() {
	fmt.Println(Sum(11,12.32))	
}

func Sum[T float64 | int](a T, b T) T{
	return a + b
}

