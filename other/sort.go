package mysort

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left := mergeSort(arr[0 : len(arr)/2])
	right := mergeSort(arr[len(arr)/2 : len(arr)])
	result := make([]int, len(arr))
	l, r, k := 0, 0, 0

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result[k] = left[l]
			l++
		} else {
			result[k] = right[r]
			r++
		}
		k++
	}
	for l < len(left) {
		result[k] = left[l]
		l++
		k++
	}
	for r < len(right) {
		result[k] = right[r]
		r++
		k++
	}

	return result

}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less, greater []int
	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)

	return result
}

func main() {
	arr := []int{4, 2, 34, 56, 7, 21, 9, 6, 87}

	fmt.Println(mergeSort(arr))
	fmt.Println(quickSort(arr))
}
