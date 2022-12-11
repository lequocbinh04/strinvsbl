package strinvsbl

import (
	"math"
	"strings"
)

var PADDING = "\u200c"

var CHARS []string = []string{
	"󠁡", "󠁢", "󠁣", "󠁤", "󠁥", "󠁦", "󠁧", "󠁨", "󠁩", "󠁪", "󠁫", "󠁬", "󠁭", "󠁮", "󠁯", "󠁰", "󠁱", "󠁲", "󠁳", "󠁴", "󠁵", "󠁶", "󠁷", "󠁸", "󠁹", "󠁺", "󠁿",
}

var CHARS_MAP = make(map[string]int)

var UNICODE_CHARS = 1114112
var BASE = len(CHARS)
var LEN = lenCalc(BASE, UNICODE_CHARS)

func Init() {
	for i, val := range CHARS {
		CHARS_MAP[val] = i
	}
}

func lenCalc(base int, chars int) int {
	var len int
	var curr = 1
	for curr < chars {
		curr *= base
		len++
	}
	return len
}

func charConvert(char string) []int {
	charCode := []rune(char)[0]
	arr := make([]int, 0)

	for charCode > 0 {
		arr = append(arr, int(charCode)%BASE)
		charCode = rune(math.Floor((float64(charCode)) / float64(BASE)))
	}

	for len(arr) < LEN {
		arr = append(arr, 0)
	}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}

	return arr
}

func charEncode(convertedChar []int) string {
	res := ""
	for _, digit := range convertedChar {
		res += CHARS[digit]
	}

	return res
}

func Encode(s string) string {
	converted := make([][]int, 0)

	for _, c := range s {
		converted = append(converted, charConvert(string(c)))
	}

	res := make([]string, 0)
	for _, char := range converted {
		res = append(res, charEncode(char))
	}

	return PADDING + strings.Join(res, "")
}

func decodeChar(encodedChar []int) string {
	// Reverse the string
	for i := 0; i < len(encodedChar)/2; i++ {
		encodedChar[i], encodedChar[len(encodedChar)-i-1] = encodedChar[len(encodedChar)-i-1], encodedChar[i]
	}

	curr := 1
	charCode := 0

	for _, digit := range encodedChar {
		charCode += int(rune(digit)) * curr
		curr *= BASE
	}

	return string(charCode)
}

func Decode(s string) string {
	var curr []int
	res := ""

	for _, c := range []rune(s) {
		if c < 917601 {
			continue
		}
		curr = append(curr, CHARS_MAP[string(c)])
		//fmt.Println(CHARS_MAP[string(c)])
		if len(curr) >= LEN {
			//fmt.Println(curr)
			res += decodeChar(curr)
			curr = []int{}
		}
	}

	return res
}
