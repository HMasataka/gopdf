package main

import (
	"log"

	"github.com/signintech/gopdf"
)

const defaultX = 10
const a4MaxWith = 2360

// Reference gopdf.PageSizeA4
var PageSizeA4 = gopdf.Rect{W: 595 - 20, H: 842 - 10}

func main() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("BIZ", "./BIZUDMincho-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("BIZ", "", 12)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.MultiCell(&PageSizeA4, "HHEEEEEEEEEEEEEEEEEEELLLLLLLLLLLLLLLLLLWOOOOOOOOOOOOOOOOOLDHEEEEEEEEEEEEEEEEEEELLLLLLLLLLLLLLLLLLWOOOOOOOOOOOOOOOOOLDEEEEEEEEEEEEEEEEEEELLLLLLLLLLLLLLLLLLWOOOOOOOOOOOOOOOOOLD")

	pdf.WritePdf("hello.pdf")
}

func lineBreak(pdf *gopdf.GoPdf) {
	pdf.SetX(defaultX)
	pdf.SetY(pdf.GetY() + pdf.GetY() + 10)
}
