package main

import (
	"fmt"
	"log"
	"simple-go-pdf-generator/pkg/pdf"
	"simple-go-pdf-generator/pkg/qr_codes"
)

func main() {
	key := "simple-go-pdf-generator"
	text := fmt.Sprintf("https://github.com/akhidnukhlis/simple-go-pdf-generator/%s", key)

	if err := qr_codes.QrCodeGen(text, "qrcode-without-icon.png"); err != nil {
		log.Fatalf("gagal membuat qr-codes: %s", err)
	}

	qrCode, err := qr_codes.GenerateQRCodeWithIcon(text, "favicon.png", "qrcode-with-icon.png")
	if err != nil {
		log.Fatalf("gagal menempelkan icon ke qr-codes: %s", err)
	}
	fmt.Sprintf("sukses membuat qr-codes: %s", qrCode)

	doc, err := pdf.ConvertImageToPDF("doc.jpg")
	if err != nil {
		log.Fatalf("gagal convert image ke pdf: %s", err)
	}
	fmt.Sprintf("sukses convert image ke pdf: %s", doc)

	pdfProcess := pdf.NewPDFGopher("samples/pdf/origin/doc.pdf",
		pdf.WithOptionMetadataPDF(pdf.OptionMetadataPDF{Title: "Hero life in You", Author: "Me as Author", Subject: "You as Subject", Keywords: "Kopi Luwak"}),
		pdf.WithOptionFilePDF(pdf.OptionFilePDF{QRCodePath: "qrcode-with-icon.png", StampPosition: "c"}),
	)

	// Process file
	err = pdfProcess.ProcessFile()
	if err != nil {
		log.Fatalf("gagal memproses meta pdf: %s", err)
	}
	fmt.Sprintf("sukses memproses meta pdf: %s", pdfProcess.Base64Output)
}
