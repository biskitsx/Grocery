package service

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService interface {
	UploadImage(fileHeader *multipart.FileHeader) (string, error)
}

type cloudinaryService struct {
	cld *cloudinary.Cloudinary
	ctx context.Context
}

func NewCloudinaryService() CloudinaryService {
	CLOUDINARY_URL := os.Getenv("CLOUDINARY_URL")
	cld, _ := cloudinary.NewFromURL(CLOUDINARY_URL)
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return &cloudinaryService{
		cld: cld,
		ctx: ctx,
	}
}

func (service *cloudinaryService) UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	file, _ := fileHeader.Open()
	resp, err := service.cld.Upload.Upload(service.ctx, file, uploader.UploadParams{})
	if err != nil {
		return "error", err
	}
	return resp.SecureURL, nil
}
