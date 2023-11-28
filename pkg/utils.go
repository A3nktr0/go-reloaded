package go_reloaded

import (
	"strconv"
	"strings"
)

//************************************************************************************************
//************************************************************************************************
// Package of tools used by the program
//************************************************************************************************
//************************************************************************************************

func IsAlpha(s string) bool {
	minMaj := 65
	maxMaj := 90
	minMin := 97
	maxMin := 122
	minNum := 48
	maxNum := 57
	count := 0
	for _, v := range s {
		if (v >= rune(minMin)) && (v <= rune(maxMin)) || (v >= rune(minMaj)) && (v <= rune(maxMaj)) || (v >= rune(minNum)) && (v <= rune(maxNum)) {
			count++
		}
	}
	return count == len(s)
}

func IsLower(s string) bool {
	count := 0
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			count++
		}
	}
	return count == len(s)
}

func IsUpper(s string) bool {
	count := 0
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			count++
		}
	}
	return count == len(s)
}

func Capitalize(s string) string {
	str := []rune(s)
	if IsLower(string(str[0])) {
		str[0] -= 32
	}
	for i := 1; i < len(str); i++ {
		if IsAlpha(string(str[i-1])) {
			if IsUpper(string(str[i])) {
				str[i] += 32
			}
		} else if IsLower(string(str[i])) {
			str[i] -= 32
		}
	}
	return string(str)
}

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Lowercase(s string) string {
	return strings.ToLower(s)
}

// Convert hexadecimal to decimal
func HexToDecimal(h string) string {
	decode, _ := strconv.ParseInt(h, 16, 64)
	return strconv.Itoa(int(decode))
}

// Convert binary to decimal
func BinToDecimal(binary string) string {
	decode, _ := strconv.ParseInt(binary, 2, 64)
	return strconv.Itoa(int(decode))
}

func RemoveIndex(arr []string, i int) []string {
	return append(arr[:i], arr[i+1:]...)
}

// Check if Numeric
func IsNumeric(s string) bool {
	minNum := 48
	maxNum := 57
	count := 0
	for _, v := range s {
		if v >= rune(minNum) && v <= rune(maxNum) {
			count++
		}
	}
	return count == len(s)
}

// Check if Vowel
func IsVowel(c string) bool {
	if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" || c == "A" || c == "E" || c == "I" || c == "O" || c == "U" {
		return true
	} else {
		return false
	}
}

// Check if Special Character
func IsPunct(c rune) bool {
	if c >= 33 && c <= 63 {
		return true
	} else {
		return false
	}
}

func IsQuote(c rune) bool {
	if c == '"' || c == '`' || c == '\'' {
		return true
	} else {
		return false
	}
}
