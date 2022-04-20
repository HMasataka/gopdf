package main

import (
	"log"

	"github.com/signintech/gopdf"
)

const defaultX = 10

func main() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("BIZ", "./BIZUDMincho-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("BIZ", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Cell(nil, "Hello, World")
	lineBreak(&pdf)
	pdf.Cell(nil, "Hello, World")

	pdf.WritePdf("hello.pdf")
}

func lineBreak(pdf *gopdf.GoPdf) {
	pdf.SetX(defaultX)
	pdf.SetY(pdf.GetY() + pdf.GetY() + 10)
}
