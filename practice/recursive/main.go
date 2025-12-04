package main

import "fmt"

func iterator(index int) int {

	result := 0

	for i := 0; i < index+1; i++ {
		// result = result + i
		result += i
	}

	return result
}

func recursor(index int) int {

	if index == 1 {
		return 1
	}

	if index > 1 {
		return index + recursor(index-1)
	}
	return 0
}

func main() {
	result1 := 0 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10
	fmt.Println(result1)

	result2 := iterator(10)

	fmt.Printf("the result of the ITERATIVE function is %v\n", result2)

	result3 := recursor(10)

	fmt.Printf("the result of the RECURSIVE function is %v\n", result3)

}
