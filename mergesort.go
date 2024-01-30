package mergesort

import "cmp"

func Sort[T cmp.Ordered](arr []T) []T {
	return SortCustom(arr, cmp.Compare)
}

func SortReversed[T cmp.Ordered](arr []T) []T {
	return SortCustom(arr, func(a, b T) int { return -cmp.Compare(a, b) })
}

func SortCustom[T cmp.Ordered](arr []T, comp func(a, b T) int) []T {
	if len(arr) <= 1 {
		return arr
	}

	arr1 := SortCustom(arr[:len(arr)/2], comp)
	arr2 := SortCustom(arr[len(arr)/2:], comp)

	return merge(arr1, arr2, comp)
}

func merge[T cmp.Ordered](arr1 []T, arr2 []T, comp func(a, b T) int) []T {
	result := make([]T, len(arr1)+len(arr2))

	i := 0
	j := 0

	for i < len(arr1) && j < len(arr2) {
		switch c := comp(arr1[i], arr2[j]); {
		case c == 0:
			result[i+j] = arr1[i]
			i++
			result[i+j] = arr2[j]
			j++
		case c < 0:
			result[i+j] = arr1[i]
			i++
		case c > 0:
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
