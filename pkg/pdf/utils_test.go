package pdf_test

import (
	"github.com/stretchr/testify/assert"
	util "simple-pdf-generator-golang/pkg/pdf"
	"testing"
)

func TestIsStructEmpty(t *testing.T) {
	result := util.IsStructEmpty(&util.OptionMetadataPDF{
		Title:   "",
		Author:  "",
		Subject: "",
	})

	assert.Equal(t, true, result)
	assert.NotEmpty(t, result)
}

func TestGetFileType(t *testing.T) {
	result := util.GetFileType("doc.jpg")

	expectImage := util.Image

	assert.Equal(t, expectImage, result)
	assert.NotEmpty(t, result)
}

func TestChangeFileExtension(t *testing.T) {
	result := util.ChangeFileExtension("doc.jpg", "pdf")

	expectFile := "doc.pdf"

	assert.Equal(t, expectFile, result)
	assert.NotEmpty(t, result)
}

func TestAddedMetadata(t *testing.T) {
	err := util.AddedMetadata("../samples/pdf/origin/doc.pdf", &util.OptionMetadataPDF{
		Title:   "Test Tittle",
		Author:  "Test Author",
		Subject: "Test Subject",
	})

	assert.NoError(t, err)
}

func TestAddedKeywords(t *testing.T) {
	err := util.AddKeywords("../samples/pdf/origin/doc.pdf", &util.OptionMetadataPDF{
		Keywords: "Test Keywords",
	})

	assert.NoError(t, err)
}

func TestEncrypted(t *testing.T) {
	err := util.Encrypted("../samples/pdf/origin/doc.pdf", "012345")

	assert.NoError(t, err)
}

func TestDecrypted(t *testing.T) {
	err := util.Decrypted("../samples/pdf/origin/doc.pdf", "012345")

	assert.NoError(t, err)
}

func TestHasPDFPassword(t *testing.T) {
	result, err := util.HasPDFPassword("../samples/pdf/origin/doc.pdf", "012345")

	assert.NotEmpty(t, result)
	assert.Equal(t, true, result)
	assert.NoError(t, err)
}
