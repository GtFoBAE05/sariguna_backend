package images

import (
	"context"
	"errors"
	"path/filepath"
	"sariguna_backend/sariguna/pkg/configs"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func Upload(file interface{}, path string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cloudConfig := configs.LoadCloudinaryConfig()

	cld, err := cloudinary.NewFromParams(cloudConfig.CLOUDNAME, cloudConfig.CLOUDAPIKEY, cloudConfig.CLOUDAPISECRET)
	if err != nil {
		return "", err
	}

	filename := uuid.NewString()

	uploadParam, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: path + "/" + filename,
		Eager:    "q_auto:low",
	})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}

func Remove(path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cloudConfig := configs.LoadCloudinaryConfig()

	cld, err := cloudinary.NewFromParams(cloudConfig.CLOUDNAME, cloudConfig.CLOUDAPIKEY, cloudConfig.CLOUDAPISECRET)
	if err != nil {
		return err
	}

	pubId := getPublicID(path)
	res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: pubId,
	})
	if err != nil {
		return err
	}
	if strings.Contains(res.Result, "not found") {
		return errors.New("image not found")
	}
	return nil
}

func getPublicID(url string) string {
	// Membagi URL berdasarkan "/"
	parts := strings.Split(url, "/")
	// Mengambil bagian terakhir dari pembagian tersebut
	fileNameWithExt := parts[len(parts)-1]
	// Menghapus ekstensi file
	fileName := strings.TrimSuffix(fileNameWithExt, filepath.Ext(fileNameWithExt))
	// Menggabungkan bagian nama direktori dan nama file tanpa ekstensi
	imageID := parts[len(parts)-2] + "/" + fileName
	return imageID
}
