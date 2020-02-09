package countingsort

//Sort sorts an int slice using counting sort
func Sort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	temp := make([]int, max(arr)+1)

	for _, v := range arr {
		temp[v]++
	}

	curr := 0

	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[curr] = i
			curr++
		}
	}

	return arr
}

func max(arr []int) int {
	max := arr[0]

	for _, v := range arr[1:len(arr)] {
		if v > max {
			max = v
		}
	}

	return max
}
