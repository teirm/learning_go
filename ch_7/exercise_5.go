package main

import (
    "fmt"
    "io"
    "strings"
)

type LimitedReader struct {
    reader  io.Reader
    curr    int64
    limit   int64
}

// give the LimitedReader the io.Reader interface by
// defining the Read method
func (l *LimitedReader) Read(p []byte) (int, error) {
    read := 0 
    var err error
    
    // check to see if entire buffer can be read
    if int64(len(p)) < (l.limit - l.curr) { 
        read, err = l.reader.Read(p)
    } else {
    // read up to limit into buffer
        end := l.limit - l.curr
        read, err = l.reader.Read(p[:end])
    }
    l.curr += int64(read)
    // return io.EOF if there is no error already
    // and limit has not been reached
    if l.curr >= l.limit && err == nil {
        err = io.EOF
    }
    return read, err 
}

// create an io.Reader from an io.Reader with a 
// limit on how many bytes it will read before returning
// io.EOF
func LimitReader(r io.Reader, n int64) io.Reader {
    var l LimitedReader
    l.reader = r
    l.curr = 0
    l.limit = n
    return &l
}

func main() {
    reader := strings.NewReader("fishsticks and custard")
    limitReader := LimitReader(reader, 6)
    buf := make([]byte, 5)
    count, err := limitReader.Read(buf[:])
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    fmt.Printf("Read %d bytes\n", count)
}
