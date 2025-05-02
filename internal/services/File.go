package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func SaveImage(file multipart.File, originalFilename string) (string, error) {
	uploadDir := "uploads"

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), originalFilename)
	path := filepath.Join(uploadDir, filename)

	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filename, nil
}


func GetImage(url string) (string, error) {

	path := filepath.Join("uploads", url)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist")
	}

	return path, nil
}
