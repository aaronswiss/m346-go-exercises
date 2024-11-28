package main

import "fmt"

func main() {
	var fibs = []int{1, 1}

	for i := 2; i < 9; i++ {
		nextFib := fibs[i-1] + fibs[i-2]
		fibs = append(fibs, nextFib)
	}

	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
