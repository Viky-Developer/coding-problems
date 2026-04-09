package twodarray

// TwoDArray traverses a 3x3 matrix in a custom order and returns
// the elements as a flattened slice.
//
// Traversal order:
//   - First row (left to right)
//   - Second row (only last element)
//   - Third row (right to left)
//   - Remaining elements of second row (first and middle)
//
// Example:
// Input:
//
//	[[0 1 2]
//	 [3 4 5]
//	 [6 7 8]]
//
// Output:
//
//	[0 1 2 5 8 7 6 3 4]
//
// Note:
// This function is designed specifically for a 3x3 matrix.
// It is not generic and may break for other matrix sizes.
func TwoDArray(matrix [][]int) []int {

	var result []int

	for i := range matrix {

		switch i {
		case 0:
			for j := range matrix {
				result = append(result, matrix[i][j])
			}
		case 1:
			result = append(result, matrix[i][2])
		case 2:
			for j := len(matrix) - 1; j >= 0; j-- {
				result = append(result, matrix[i][j])
			}
		}
	}

	result = append(result, matrix[1][0])
	result = append(result, matrix[1][1])

	return result
}
