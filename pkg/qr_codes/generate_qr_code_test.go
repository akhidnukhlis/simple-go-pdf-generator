package qr_codes_test

import (
	"fmt"
	"simple-pdf-generator-golang/pkg/qr_codes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateQRCode(t *testing.T) {
	key := "simple-pdf-generator-golang"
	text := fmt.Sprintf("https://github.com/akhidnukhlis/simple-pdf-generator-golang/%s", key)

	err := qr_codes.QrCodeGen(text, "qrcode-without-icon.png")

	assert.NoError(t, err)
}
