package mergesort

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Sort[T Number](arr []T) []T {
	return SortCustom(arr, func(a, b T) T { return a - b })
}

func SortReversed[T Number](arr []T) []T {
	return SortCustom(arr, func(a, b T) T { return b - a })
}

func SortCustom[T Number](arr []T, comp func(a, b T) T) []T {
	if len(arr) <= 1 {
		return arr
	}

	arr1 := SortCustom(arr[:len(arr)/2], comp)
	arr2 := SortCustom(arr[len(arr)/2:], comp)

	return merge(arr1, arr2, comp)
}

func merge[T Number](arr1 []T, arr2 []T, comp func(a, b T) T) []T {
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
