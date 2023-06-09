package pdf_test

import (
	"fmt"
	entity "go-pdf-generator/entities"
	"go-pdf-generator/pdf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertImageToPDF(t *testing.T) {
	doc, err := pdf.ConvertImageToPDF("doc.jpg")

	assert.NoError(t, err)
	assert.NotEmpty(t, doc)
}

func TestAddQRCodeToPDF(t *testing.T) {
	err := pdf.AddQRCodeToPDF("../samples/pdf/origin/doc.pdf", "qrcode-generate.png", "br")

	assert.NoError(t, err)
}

func TestProcessPDF(t *testing.T) {
	pdfProcess := pdf.NewPDFGopher("doc.pdf",
		pdf.WithOptionMetadataPDF(entity.OptionMetadataPDF{Title: "Hero life in You", Author: "Me as Author", Subject: "You as Subject"}),
		pdf.WithOptionFilePDF(entity.OptionFilePDF{QRCodePath: "qrcode-generate.png", StampPosition: "c"}),
	)

	// Process file
	err := pdfProcess.ProcessFile()

	fmt.Printf("pdfProcess.Base64Output: %v\n", pdfProcess.Base64Output)

	assert.NoError(t, err)
	assert.NotEmpty(t, pdfProcess.Base64Output)
}
