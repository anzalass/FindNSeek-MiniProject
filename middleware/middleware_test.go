package middleware

import (
	"bytes"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	res, _ := CreateToken("abcd123", "anzalas", "anz@gmail.com")
	assert.Equal(t, res, res)

}
func TestExtraxtToken(t *testing.T) {
	res, _ := ExtractToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuemFsYXMubXVoYW1tYWRAZ21haWwuY29tIiwiZXhwIjoxNjk4NjQ3NzA4LCJpZCI6IjY3NmY4YzE1LWNjNGYtNDYwZC04NjZkLTFjOWNkMTEwYTAyYSIsIm5hbWUiOiJhbnphbGFzIn0.oTwKpUg2dolsYspuPROtVbwPiFFo1bDnjbA79Ttyz6Q")
	assert.Equal(t, res, res)
}

func TestImageHandler(t *testing.T) {
	var fakeFileContents = []byte("file.jpg")
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "file.jpg")
	part.Write(fakeFileContents)
	writer.Close()
	res, _ := ImageUploader(body)
	assert.Equal(t, res, res)
}
