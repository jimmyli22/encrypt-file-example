package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {
	p, err := pdfcpu.CreateResourceDictInheritanceDemoXRef()
	if err != nil {
		panic(err)
	}

	path := "example.pdf"

	conf := api.LoadConfiguration()
	conf.UserPW = "123"
	conf.OwnerPW = "123"
	err = api.CreatePDFFile(p, path, conf)
	if err != nil {
		panic(err)
	}

	err = api.EncryptFile(path, path, conf)
	if err != nil {
		panic(err)
	}
}
