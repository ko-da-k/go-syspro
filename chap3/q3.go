package main

import (
	"archive/zip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("q3.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("q3.txt")
	if err != nil {
		panic(err)
	}

	src, err := os.Open("q3.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	_, err = io.Copy(writer, src)
	if err != nil {
		panic(err)
	}
}
