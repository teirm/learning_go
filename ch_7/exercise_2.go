package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	BytesWritten int64
	Writer       io.Writer
}

func (cw *CountWriter) Write(p []byte) (int, error) {
	written, err := cw.Writer.Write(p)
	cw.BytesWritten += int64(written)
	return written, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cw CountWriter
	cw.BytesWritten = 0
	cw.Writer = w
	return &cw, &cw.BytesWritten
}

func main() {
	writer, count := CountingWriter(os.Stdout)

	writer.Write([]byte("Hello\n"))
	fmt.Println(*count)

	writer.Write([]byte("Fish Sticks\n"))
	fmt.Println(*count)

}
