package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	result := 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		result += 1
	}
	*w += WordCounter(result)
	return result, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	result := 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		result += 1
	}
	*l += LineCounter(result)
	return result, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	var w WordCounter
	w.Write([]byte("Hello fish sticks"))
	fmt.Println(w)

	w = 0
	fmt.Fprintf(&w, "hello, %s", name)
	fmt.Println(w)

	var l LineCounter
	l.Write([]byte("Hello\nFish\nSticks"))
	fmt.Println(l)
}
