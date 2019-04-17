package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Fprintf(os.Stdout, "%f", 0.1)
	fmt.Fprintf(os.Stdout, "%s", "hoge")
}
