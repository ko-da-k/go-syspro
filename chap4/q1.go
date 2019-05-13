package main

import (
	"fmt"
	"time"
)

func main() {
	after := time.After(5 * time.Second)
	fmt.Println("job started")
	<-after
	fmt.Println("job finished")
}
