package mergesort

import "cmp"

func Sort[T cmp.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	arr1 := Sort(arr[:len(arr)/2])
	arr2 := Sort(arr[len(arr)/2:])

	return merge(arr1, arr2)
}

func merge[T cmp.Ordered](arr1 []T, arr2 []T) []T {
	result := make([]T, len(arr1)+len(arr2))

	i := 0
	j := 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result[i+j] = arr1[i]
			i++
		} else {
			result[i+j] = arr2[j]
			j++
		}
	}

	for i < len(arr1) {
		result[i+j] = arr1[i]
		i++
	}

	for j < len(arr2) {
		result[i+j] = arr2[j]
		j++
	}

	return result
}
