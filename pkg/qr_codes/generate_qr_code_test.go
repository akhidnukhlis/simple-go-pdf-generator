package qr_codes_test

import (
	"fmt"
	"simple-go-pdf-generator/pkg/qr_codes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateQRCode(t *testing.T) {
	key := "simple-go-pdf-generator"
	text := fmt.Sprintf("https://github.com/akhidnukhlis/simple-go-pdf-generator/%s", key)

	err := qr_codes.QrCodeGen(text, "qrcode-without-icon.png")

	assert.NoError(t, err)
}
