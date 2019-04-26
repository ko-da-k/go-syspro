package main

import (
	"crypto/rand"
	"io"
	"fmt"
)

func main() {
	buffer := make([]byte, 16)
	io.ReadFull(rand.Reader, buffer)
	fmt.Printf("%#v", buffer)
}
