package targetvalue

/*
BruteForceTargetValue finds two indices such that their values sum to the target.

This approach checks all possible pairs using a nested loop.

Time Complexity:

	O(n^2) — checks every pair

Space Complexity:

	O(1) — no extra space used

Note:
  - Works for both sorted and unsorted arrays
  - Not efficient for large datasets
*/
func BruteForceTargetValue(arr []int, target int) []int {

	temp := 0

	for i := range arr {
		for j := i + 1; j < len(arr); j++ {

			temp = arr[i] + arr[j]

			if target == temp {
				return []int{i, j}
			}
		}
	}

	return nil
}

/*
SortTargetValue finds two indices using the two-pointer technique.

This approach assumes the array is sorted.

Time Complexity:

	O(n) — single pass with two pointers

Space Complexity:

	O(1) — no extra space used

Important:
  - Only works correctly on sorted arrays
  - If array is unsorted, results may be incorrect
  - Sorting before using this changes original indices

Example (works):

	arr := []int{2, 4, 6, 8}, target := 10

Example (fails):

	arr := []int{8, 1, 7, 3}, target := 10
*/
func SortTargetValue(arr []int, target int) []int {

	left := 0

	right := len(arr) - 1

	for left < right {

		sum := arr[left] + arr[right]

		if target == sum {
			return []int{left, right}

		} else if sum < target {
			left++

		} else {
			right--
		}
	}

	return nil
}

/*
HashMapTargetValue finds two indices such that their values sum to the target.

This approach uses a hash map to store previously seen values.

Time Complexity:

	O(n) — single pass through array

Space Complexity:

	O(n) — extra space for hash map

Why it works:
  - For each element, compute diff = target - value
  - Check if diff already exists in map
  - If yes → solution found

Advantages:
  - Works on unsorted arrays
  - Maintains original indices
  - Most efficient approach

Note:
  - Map stores value → index
*/
func HashMapTargetValue(arr []int, target int) []int {

	m := make(map[int]int)

	for i, val := range arr {

		diff := target - val

		if idx, ok := m[diff]; ok {

			return []int{idx, i}
		}

		m[val] = i
	}

	return nil
}
