package main

import (
	"io"
	"os"
)

func main() {
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		panic(err)
		return
	}

	var r io.Reader
	r = tty

	var w io.Writer
	w = r.(io.Writer)

	_, err = w.Write([]byte("hello world\n"))
	if err != nil {
		panic(err)
		return
	}
}
