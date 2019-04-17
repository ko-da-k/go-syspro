package main

import (
	"os"
	"encoding/csv"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{"Alice", "20"})
	writer.Flush()
}
