package qr_codes_test

import (
	"go-pdf-generator/qr_codes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateQRCodeWithIcon(t *testing.T) {
	qrCode, err := qr_codes.GenerateQRCodeWithIcon("https://google.com", "favicon.png", "qrcode-generate.png")

	assert.NoError(t, err)
	assert.NotEmpty(t, qrCode)
}
