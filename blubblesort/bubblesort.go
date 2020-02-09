//Package bubblesort contains a BubbleSort implementation
package bubblesort

//Sort sorts an int slice using bubble sort
func Sort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := 0; i < len(arr); i++ {
		somethingSorted := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				continue
			}

			somethingSorted = true
			swap(arr, j)
		}
		if !somethingSorted {
			break
		}
	}
}
func swap(arr []int, idx int) {
	arr[idx], arr[idx+1] = arr[idx+1], arr[idx]
}
