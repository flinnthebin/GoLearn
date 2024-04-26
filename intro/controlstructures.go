package main

import "fmt"

func forLoop() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(numbers); i++ {
		fmt.Println("iterator: ", i, " | ", "number: ", numbers[i])
	}
}

func forRange() {
	names := []string{"a", "b", "c", "d"}

	for _, name := range names {
		fmt.Println("iterator: ", " | ", "name: ", name)
	}
}

func rangeMap() {
	users := map[string]int{
		"foo":   1,
		"bar":   2,
		"baz":   3,
		"Alice": 33,
		"Bob":   88,
	}

	for key, value := range users {
		fmt.Printf("key: %s | value %d\n", key, value)
	}
}

func main() {

	fmt.Println("For Loop")
	forLoop()
	fmt.Println("\nFor Range")
	forRange()
	fmt.Println("\nFor Range Map")
	rangeMap()
}
