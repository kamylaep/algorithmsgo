package quicksort

//Sort sorts int slices using quick sort algorithm
func Sort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	var left []int
	var right []int
	pivot := arr[0]

	for _, v := range arr[1:len(arr)] {
		if pivot <= v {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	result := append(Sort(right), pivot)
	result = append(result, Sort(left)...)

	return result
}
