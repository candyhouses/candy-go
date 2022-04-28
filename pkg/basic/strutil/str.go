package strutil

import (
	"strings"

	"github.com/candyhouses/candy-go/pkg/basic/sliceutil"
)

// ------ Blank

//IsBlank Check if a string is blank.
//blank string means
//1.Empty string
//2.Consists of invisible characters e.g (" ","\n","\t","\r")

//IsBlankIfStr Check if a interface{} is string and if this string is blank.
func IsBlank(str string) bool {

	if isEmpty := IsEmpty(str); isEmpty {
		return isEmpty
	}
	for i := 0; i < len(str); i++ {
		if !isBlankASCII(str[i]) {
			return false
		}
	}

	return true

}

//IsNotBlank Check if a string is not blank
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

//HasBlank Check if has blank string in the array of string.
//is equivalent to : IsBlank(...) || IsBlank(...) || ...
func HasBlank(strs ...string) bool {
	if sliceutil.IsEmpty(strs) {
		return true
	}

	for _, v := range strs {
		if IsBlank(v) {
			return true
		}
	}

	return false

}

//IsAllBlank Check if all string is blank
//IsAllBlank is equivalent to : IsBlank(...) && IsBlank(...) && ...
func IsAllBlank(strs ...string) bool {

	if sliceutil.IsEmpty(strs) {
		return true
	}

	for _, v := range strs {
		if IsNotBlank(v) {
			return false
		}
	}

	return true
}

//IsAllNotEmpty Check if all of strs is Blank.
func IsAllNotBlank(strs ...string) bool {
	return !HasBlank(strs...)
}

//BlankToDefault  Check if a string is blank ,if it's blank ,set a default value.
func BlankToDefault(str, defaultStr string) string {

	if IsBlank(str) {
		return defaultStr
	}
	return str
}

// ------ Empty

//IsEmpty Check if a string is empty.
//Empty means
//1. ""
//2. len(str) == 0
//E.g:
//1. IsEmpty("") == true
//2. IsEmpty(" ") == false
//3. IsEmpty(" \n \t \r") = false
//
//IsEmpty can't check blank string.Check blank string should use IsBlank

func IsEmpty(str string) bool {

	return str == "" || len(str) == 0

}

//IsNotEmpty Check if a string is not empty
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

//EmptyToDefault Check if a string is empty ,if it's empty ,set a default value.
func EmptyToDefault(str, defaultStr string) string {

	if IsEmpty(str) {
		return defaultStr
	}
	return str
}

//HasEmpty  Check if has empty string in the array of string.
func HasEmpty(strs ...string) bool {
	if sliceutil.IsEmpty(strs) {
		return true
	}

	for _, v := range strs {
		if IsEmpty(v) {
			return true
		}
	}
	return false
}

//IsAllEmpty Check if all string is empty
func IsAllEmpty(strs ...string) bool {

	if sliceutil.IsEmpty(strs) {
		return true
	}
	for _, v := range strs {
		if IsNotEmpty(v) {
			return false
		}
	}
	return true
}

//IsAllNotEmpty Check if all of strs is Empty.
func IsAllNotEmpty(strs ...string) bool {
	return !HasEmpty(strs...)
}

//------- check underfined

//isNullOrUnderfined Check if str is null or undefined
func IsNullOrUnderfined(str string) bool {
	str = TrimBlankChar(str)
	return NULL == str || UNDEFINED == str
}

func IsBlankOrUnderfined(str string) bool {
	if IsBlank(str) {
		return true
	}
	return IsNullOrUnderfined(str)
}

func IsEmptyOrUnderfined(str string) bool {
	if IsEmpty(str) {
		return true
	}
	return IsNullOrUnderfined(str)
}

//-------- trim

//trimBlankCharByMode Removes blank char at the beginning and end of the string by mode.
//mode :
//	-1 : left trim
//	 0 : left and right
//	 1 : right trim
func trimBlankCharByMode(str string, mode int) string {
	if IsEmpty(str) {
		return ""
	}

	strSize := len(str)

	headPoint, endPoint := 0, strSize

	if mode <= 0 {
		for headPoint < endPoint && isBlankASCII(str[headPoint]) {
			headPoint++
		}
	}

	if mode >= 0 {
		for headPoint < endPoint && isBlankASCII(str[endPoint-1]) {
			endPoint--
		}
	}

	return str[headPoint:endPoint]

}

func TrimBlankCharRight(str string) string {
	return trimBlankCharByMode(str, 1)
}

func TrimBlankCharLeft(str string) string {
	return trimBlankCharByMode(str, -1)
}

func TrimBlankChar(str string) string {
	return trimBlankCharByMode(str, 0)
}

//-----star check

func IsStartWith(str, prefix string, ignoreCase bool) bool {
	if IsEmpty(str) || IsEmpty(prefix) {

		if !ignoreCase {
			return false
		}

		return IsEmpty(str) && IsEmpty(prefix)
	}

	if ignoreCase {
		str, prefix = strings.ToLower(str), strings.ToLower(prefix)

	}

	return strings.HasPrefix(str, prefix)
}

func IsStartWithAny(str string, ignoreCase bool, prefixes ...string) bool {

	if IsEmpty(str) || sliceutil.IsEmpty(prefixes) {
		return false
	}

	for _, v := range prefixes {
		if IsStartWith(str, v, ignoreCase) {
			return true
		}
	}

	return false
}

//todo: add char utils and move this func to it.
func isBlankASCII(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}
