package qr_codes_test

import (
	"fmt"
	"simple-go-pdf-generator/pkg/qr_codes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateQRCodeWithIcon(t *testing.T) {
	key := "simple-go-pdf-generator"
	text := fmt.Sprintf("https://github.com/akhidnukhlis/simple-go-pdf-generator/%s", key)

	qrCode, err := qr_codes.GenerateQRCodeWithIcon(text, "favicon.png", "qrcode-with-icon.png")

	assert.NoError(t, err)
	assert.NotEmpty(t, qrCode)
}
