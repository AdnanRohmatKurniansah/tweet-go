package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var allowedImageExts = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true,
}

const MaxImageSize = 5 * 1024 * 1024

func randomString(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func SaveImage(c *gin.Context, fieldName string, folder string, required bool) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		if required {
			return "", fmt.Errorf("Image is required")
		}
		return "", nil
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))

	if !allowedImageExts[ext] {
		return "", fmt.Errorf("File type not allowed (jpg, jpeg, png, gif, webp)")
	}

	if file.Size > MaxImageSize {
		return "", fmt.Errorf("File too large (max 5MB)")
	}

	uploadDir := filepath.Join("uploads", folder)

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("Failed to create upload directory")
		}
	}

	fileName := randomString(16) + ext

	filePath := filepath.Join(uploadDir, fileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", fmt.Errorf("Failed to save image")
	}

	return "/" + filePath, nil
}

func DeleteImage(imageUrl string) {
	if imageUrl == "" {
		return
	}

	filePath := filepath.Join(".", imageUrl)

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return
	}
	uploadsAbs, err := filepath.Abs("uploads")
	if err != nil {
		return
	}
	if !strings.HasPrefix(absPath, uploadsAbs) {
		return
	}

	os.Remove(filePath)
}