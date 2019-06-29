// implementations of commas program
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma1("12345"))
	fmt.Println(comma2("1234568910"))
	fmt.Println(comma3("-123456.8910"))
}

func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma1(s[:n-3]) + "," + s[n-3:]
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

func comma2(s string) string {
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

func comma3(s string) string {
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

	return string(sign) + comma2(s) + decimalFragment
}
