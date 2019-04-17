package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json 化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	// ここにコードを書く
	gz, err := file.Create("q3.json.gz")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(os.Stdout, gz)
	writer.Write(source)
	writer.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
