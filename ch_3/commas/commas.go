// implementations of commas program
package commas

import (
	"bytes"
	"strings"
)

func Comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma1(s[:n-3]) + "," + s[n-3:]
}

// reverse a string
func reverseString(s string) string {
	result := make([]byte, len(s))
	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		result[j] = s[i]
		j++
	}
	return string(result)
}

func Comma2(s string) string {
	n := len(s)
	var buf bytes.Buffer
	if n <= 3 {
		return s
	}
	for len(s) > 3 {
		buf.WriteString(reverseString(s[n-3:]))
		buf.WriteString(",")
		s = s[:n-3]
		n -= 3
	}
	buf.WriteString(reverseString(s))

	// reverse the string
	return reverseString(buf.String())
}

func Comma3(s string) string {
	var sign byte
	if s[0] == '-' || s[0] == '+' {
		sign = s[0]
		s = s[1:]
	}

	var decimalFragment string
	decimalIndex := strings.LastIndex(s, ".")
	if decimalIndex != -1 {
		decimalFragment = s[decimalIndex:]
		s = s[:decimalIndex]
	}

	result := Comma2(s)
	if sign == '-' || sign == '+' {
		result = string(sign) + result
	}

	if decimalFragment != "" {
		result = result + decimalFragment
	}

	return result
}
