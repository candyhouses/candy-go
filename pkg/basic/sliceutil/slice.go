package sliceutil

//IsEmpty Check if a slice is empty
func IsEmpty[T int8 | int | int16 | int32 | int64 | string | float32 | float64 | bool | byte](s []T) bool {
	if nil == s {
		return true
	}
	return len(s) == 0
}

//Insert Insert a element or a slience at the specified position of the sli
func Insert[T int8 | int | int16 | int32 | int64 | string | float32 | float64 | bool | byte](sli []T, i int, element ...T) []T {

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

// func AddAll[T int8 | int | int16 | int32 | int64 | string | float32 | float64 | bool | byte](slis ...[]T) []T {

// }
