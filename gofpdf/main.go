package main

import "github.com/jung-kurt/gofpdf"

func main() {
	// orientation, unit, size, fontPath
	pdf := gofpdf.New("P", "mm", "A4", "")
	// actionFlag, user pw, owner pw
	pdf.SetProtection(gofpdf.CnProtectPrint, "123", "abc")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Write(10, "Hello world!")
	err := pdf.OutputFileAndClose("example.pdf")
	if err != nil {
		panic(err)
	}
}
