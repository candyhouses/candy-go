package strutil

import (
	"strings"

	"github.com/candyhouses/candy-go/pkg/basic/sliceutil"
)

// ------ Space

// IsSpace Check if a string is space.
// IsSpace string means
// 1.Empty string
// 2.Consists of invisible characters e.g (" ","\n","\t","\r")

// IsBlankIfStr Check if a interface{} is string and if this string is blank.
func IsSpace(str string) bool {
	if isEmpty := IsEmpty(str); isEmpty {
		return isEmpty
	}
	for i := 0; i < len(str); i++ {
		if !isSpaceASCII(str[i]) {
			return false
		}
	}

	return true
}

// IsNotSpace Check if a string is not Space
func IsNotSpace(str string) bool {
	return !IsSpace(str)
}

// HasSpace Check if has Space string in the array of string.
// is equivalent to : IsSpace(...) || IsSpace(...) || ...
func HasSpace(strs ...string) bool {
	if sliceutil.IsEmpty(strs) {
		return true
	}

	for _, v := range strs {
		if IsSpace(v) {
			return true
		}
	}

	return false
}

// IsAllSpace Check if all string is Space
// IsAllSpace is equivalent to : IsSpace(...) && IsSpace(...) && ...
func IsAllSpace(strs ...string) bool {
	if sliceutil.IsEmpty(strs) {
		return true
	}

	for _, v := range strs {
		if IsNotSpace(v) {
			return false
		}
	}

	return true
}

// IsAllNotSpace Check if all of strs is Space.
func IsAllNotSpace(strs ...string) bool {
	return !HasSpace(strs...)
}

// SpaceToDefault  Check if a string is Space ,if it's Space ,set a default value.
func SpaceToDefault(str, defaultStr string) string {
	if IsSpace(str) {
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

// IsNotEmpty Check if a string is not empty
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// EmptyToDefault Check if a string is empty ,if it's empty ,set a default value.
func EmptyToDefault(str, defaultStr string) string {
	if IsEmpty(str) {
		return defaultStr
	}
	return str
}

// HasEmpty  Check if has empty string in the array of string.
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

// IsAllEmpty Check if all string is empty
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

// IsAllNotEmpty Check if all of strs is Empty.
func IsAllNotEmpty(strs ...string) bool {
	return !HasEmpty(strs...)
}

//------- check underfined

// isNullOrUnderfined Check if str is null or undefined
func IsNullOrUnderfined(str string) bool {
	str = TrimSpace(str)
	return NULL == str || Undefind == str
}

func IsSpaceOrUnderfined(str string) bool {
	if IsSpace(str) {
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

//-----star check and end check

// IsStartWith Checks if the str starts with a prefix
// ignoreCase ，Indicates whether case is ignored
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

func IsEndWith(str, suffix string, ignoreCase bool) bool {
	if IsEmpty(str) || IsEmpty(suffix) {
		return IsEmpty(str) && IsEmpty(suffix)
	}

	if ignoreCase {
		str, suffix = strings.ToLower(str), strings.ToLower(suffix)
	}
	return strings.HasSuffix(str, suffix)
}

func IsEndWithAny(str string, ignoreCase bool, suffixs ...string) bool {
	if IsEmpty(str) || sliceutil.IsEmpty(suffixs) {
		return false
	}

	for _, v := range suffixs {
		if IsEndWith(str, v, ignoreCase) {
			return true
		}
	}

	return false
}

//------ contains
func Contains(str, substr string, ignoreCase bool) bool {
	if ignoreCase {
		str = strings.ToLower(str)
		substr = strings.ToLower(substr)
	}

	return strings.Contains(str, substr)
}

//---- string operator

func CutByMax(str string, max int) string {
	if max < 0 {
		panic("max mast bigger than 0")
	}

	if IsEmpty(str) {
		return ""
	}

	if len(str) < max {
		return str
	}

	res := strings.Builder{}

	for _, v := range str {
		if max > 0 {
			res.WriteRune(v)
			max--
		} else {
			break
		}
	}
	res.WriteString(MoreInfoTarget)
	return res.String()
}

func ToSymbolCase(str string, symbol byte) string {
	if IsEmpty(str) {
		return ""
	}

	builder := strings.Builder{}

	for _, v := range str {
		if isUpperCase(v) {
			builder.WriteByte(symbol)
			builder.WriteRune(toLower(v))
		} else {
			builder.WriteRune(v)
		}
	}

	res := builder.String()
	if res[0] == symbol {
		return res[1:]
	}
	return res
}

func ToUnderLineCase(str string) string {
	return ToSymbolCase(str, '-')
}

//-------- trim

// trimSpaceByMode Removes space char at the beginning and end of the string by mode.
// mode :
//	-1 : left trim
//	 0 : left and right
//	 1 : right trim
func trimSpaceByMode(str string, mode int) string {
	if IsEmpty(str) {
		return ""
	}

	strSize := len(str)

	headPoint, endPoint := 0, strSize

	if mode <= 0 {
		for headPoint < endPoint && isSpaceASCII(str[headPoint]) {
			headPoint++
		}
	}

	if mode >= 0 {
		for headPoint < endPoint && isSpaceASCII(str[endPoint-1]) {
			endPoint--
		}
	}

	return str[headPoint:endPoint]
}

// TrimSpaceRight Removes space char at the  end of the string
func TrimSpaceRight(str string) string {
	return trimSpaceByMode(str, 1)
}

// TrimSpaceLeft Removes space char at the beginning of the string
func TrimSpaceLeft(str string) string {
	return trimSpaceByMode(str, -1)
}

// TrimSpace Removes space char at the beginning and end of the string
func TrimSpace(str string) string {
	return trimSpaceByMode(str, 0)
}

// todo: add char utils and move this func to it.
func isSpaceASCII(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

// todo: add char utils and move this func to it.
func isUpperCase(char rune) bool {
	if char >= 'A' && char <= 'Z' {
		return true
	}
	return false
}

// todo: add char utils and move this func to it.
func toLower(char rune) rune {
	if isUpperCase(char) {
		char += 'a' - 'A'
	}

	return char
}
