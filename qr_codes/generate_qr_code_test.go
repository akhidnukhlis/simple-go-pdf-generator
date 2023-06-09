package qr_codes_test

import (
	"go-pdf-generator/qr_codes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateQRCode(t *testing.T) {
	err := qr_codes.QrCodeGen("Example text to create QR Codes", "qrcode-generate.png")

	assert.NoError(t, err)
}
