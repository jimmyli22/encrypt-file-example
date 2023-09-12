package main

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/alexmullins/zip"
)

func main() {
	contents := []byte("Hello World")
	fzip, err := os.Create(`./test.zip`)
	if err != nil {
		panic(err)
	}
	zipw := zip.NewWriter(fzip)
	defer zipw.Close()
	w, err := zipw.Encrypt(`test.txt`, `golang`)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(w, bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}

	w, err = zipw.Encrypt(`test1.txt`, `golang`)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(w, bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
}
