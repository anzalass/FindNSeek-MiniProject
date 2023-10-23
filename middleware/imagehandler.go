package middleware

import (
	"context"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/sirupsen/logrus"
)

func ImageUploader(input interface{}) (string, error) {
	ctx := context.Background()

	url := "cloudinary://527419765713161:zid0lkkiTmJz8HpEGaMYb6xo3-Q@djtcxk7fs"

	cloudService, err := cloudinary.NewFromURL(url)
	if err != nil {
		logrus.Error("cloudService error", err.Error())
	}

	upload, err := cloudService.Upload.Upload(ctx, input, uploader.UploadParams{})
	if err != nil {
		logrus.Error("upload error", err.Error())
	}
	return upload.SecureURL, nil
}
