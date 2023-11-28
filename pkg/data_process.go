package go_reloaded

import (
	"strconv"
	"strings"
)

//************************************************************************************************
//************************************************************************************************
// Functions for data manipulation
//************************************************************************************************
//************************************************************************************************

// Put Modifications into separate index in array
func SplitParenthesis(s string) []string {
	s = strings.ReplaceAll(s, "‘", "'")
	delim := func(char rune) bool {
		return char == '(' || char == ')' || char == ' '
	}
	return strings.FieldsFunc(s, delim)
}

// Apply Text Transformation
func Transform(str_arr []string) []string {
	count := 0
	for i := 0; i < len(str_arr); i++ {
		str_arr[i] = strings.TrimSpace(str_arr[i])
		switch {
		case str_arr[i] == "cap" || str_arr[i] == "cap,":
			if IsNumeric(str_arr[i+1]) {
				n, _ := strconv.Atoi(str_arr[i+1])
				for n != 0 {
					str_arr[i-n] = Capitalize(str_arr[i-n])
					n--
				}
				RemoveIndex(str_arr, i+1)
				count++
			}
			str_arr[i-1] = Capitalize(str_arr[i-1])
			RemoveIndex(str_arr, i)
			count++
		case str_arr[i] == "up" || str_arr[i] == "up,":
			if IsNumeric(str_arr[i+1]) {
				n, _ := strconv.Atoi(str_arr[i+1])
				for n != 0 {
					str_arr[i-n] = Uppercase(str_arr[i-n])
					n--
				}
				RemoveIndex(str_arr, i+1)
				count++
			}
			str_arr[i-1] = Uppercase(str_arr[i-1])
			RemoveIndex(str_arr, i)
			count++
		case str_arr[i] == "low" || str_arr[i] == "low,":
			if IsNumeric(str_arr[i+1]) {
				n, _ := strconv.Atoi(str_arr[i+1])
				for n != 0 {
					str_arr[i-n] = Lowercase(str_arr[i-n])
					n--
				}
				RemoveIndex(str_arr, i+1)
				count++
			}
			str_arr[i-1] = Lowercase(str_arr[i-1])
			RemoveIndex(str_arr, i)
			count++
		case str_arr[i] == "hex":
			str_arr[i-1] = HexToDecimal(str_arr[i-1])
			RemoveIndex(str_arr, i)
			count++
		case str_arr[i] == "bin":
			str_arr[i-1] = BinToDecimal(str_arr[i-1])
			RemoveIndex(str_arr, i)
			count++
		}
	}
	str_arr = str_arr[:len(str_arr)-count]
	return str_arr
}

func ReplaceA(str_arr []string) []string {
	for i := 0; i < len(str_arr); i++ {
		switch str_arr[i] {
		case "a":
			if IsVowel(string(str_arr[i+1][0])) {
				str_arr[i] = "an"
			}
		case "A":
			if IsVowel(string(str_arr[i+1][0])) {
				str_arr[i] = "An"
			}
		}
	}
	return str_arr
}

// Split comma from word and put it after preceding
func CheckComma(str_arr []string) []string {
	count := 0
	for i := 0; i < len(str_arr); i++ {
		if str_arr[i][0] == ',' {
			if len(str_arr[i]) == 1 {
				str_arr[i-1] = str_arr[i-1] + ","
				RemoveIndex(str_arr, i)
				count++
			} else if len(str_arr[i]) > 1 {
				str_arr[i-1] = str_arr[i-1] + ","
				str_arr[i] = str_arr[i][1:]
			}
		}
	}
	str_arr = str_arr[:len(str_arr)-count]
	return str_arr
}

// Parse a string and remove whitespace before special characters
func Punctuation(str string) string {
	str_out := ""
	tmp := ""

	for i := 1; i <= len(str)-1; i++ {
		if i == 1 {
			tmp = tmp + string(str[0])
		}
		if IsQuote(rune(str[i-1])) && str[i] == ' ' {
			tmp = tmp + ""
		} else {
			tmp = tmp + string(str[i])
		}
	}
	for i, c := range tmp {
		if c == ' ' && IsPunct(rune(tmp[i+1])) && !IsNumeric(string(tmp[i+1])) {
			if IsQuote(rune(tmp[i+1])) && tmp[i-1] == ':' {
				str_out = str_out + " "
			}
			str_out = str_out + ""
		} else if IsQuote(c) && IsPunct(rune(tmp[i-1])) &&
			!IsNumeric(string(tmp[i-1])) && !IsQuote(rune(tmp[i-1])) {
			str_out = str_out + " " + string(c)
		} else {
			str_out = str_out + string(c)
		}
	}
	return str_out
}
