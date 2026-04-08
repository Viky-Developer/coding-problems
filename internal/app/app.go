package app

import (
	"fmt"

	targetvalue "coding-problems/internal/targetValue"
)

func StartApp() {

	// arr := []int{5, 3, 1, 4, 9, 1}
	arr := []int{8, 1, 7, 3}

	target := 10
	result := targetvalue.BruteForceTargetValue(arr, target)

	result2 := targetvalue.SortTargetValue(arr, target)

	result3 := targetvalue.HashMapTargetValue(arr, target)

	fmt.Printf("result: %v\n", result)
	fmt.Printf("result2: %v\n", result2)
	fmt.Printf("result3: %v\n", result3)
}
