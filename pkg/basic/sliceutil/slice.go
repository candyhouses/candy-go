package sliceutil

//IsEmpty Check if a slice is empty
func IsEmpty[T any](s []T) bool {
	if nil == s {
		return true
	}
	return len(s) == 0
}

//IsNotEmpty Check if a slice is not empty
func IsNotEmpty[T any](s []T) bool {
	return !IsEmpty(s)
}

//HasNull Check if has null in slices
func HasNull[T any](s ...[]T) bool {
	if IsNotEmpty(s) {
		for _, v := range s {
			if nil == v {
				return true
			}
		}
	}
	return false
}

//Insert Insert a element or a slience at the specified position of the sli
func Insert[T any](sli []T, i int, element ...T) []T {

	if IsEmpty(element) {
		return sli
	}

	if IsEmpty(sli) {
		return element
	}

	if i < 0 || i > len(sli) {
		panic("The range of index is [0 - len(oldSlice)] ")
	}

	sli = append(sli[:i], append(element, sli[i:]...)...)

	return sli
}

func Delete[T any](sli []T, start, end int) (res []T) {

	if IsEmpty(sli) || end >= len(sli) || start > end || start < 0 {
		return sli
	}

	res = append(res, sli[:start]...)
	res = append(res, sli[end+1:]...)

	return
}

func DeleteByIndex[T any](sli []T, i int) []T {

	return Delete(sli, i, i)
}

//------ Range

func RangeByStep(start, end, step int) []int {
	if start > end {
		start, end = end, start
	}

	if step <= 0 {
		step = 1
	}

	deviation := end - start
	length := deviation / step

	if deviation%step != 0 {
		length += 1
	}

	sli := make([]int, length)

	for i := 0; i < length; i++ {
		sli[i] = start
		start += step
	}
	return sli
}

func Range(end int) []int {
	return RangeByStep(0, end, 1)
}

//-------filter

func Filter[T any](sli []T, filter func(element T) bool) (res []T) {

	for _, v := range sli {

		if !filter(v) {
			res = append(res, v)
		}
	}

	return
}
