package filemanager

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/stanhoenson/krushr/internal/env"
)

func StoreMulitpartImage(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	defer file.Close()

	if err != nil {
		return "", fmt.Errorf("failed to open fileHeader's file: %v", err)
	}

	// Check if the uploaded file is an image
	isImage, err := IsImage(file)
	if err != nil {
		return "", fmt.Errorf("failed to check if the file is an image: %v", err)
	}
	if !isImage {
		return "", fmt.Errorf("the uploaded file is not an image")
	}

	// Generate a unique filename for the uploaded file
	filename := GenerateFilename(fileHeader.Filename)

	// Create the directory if it doesn't exist
	err = os.MkdirAll(env.FileStoragePath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Create the file on disk
	dst, err := os.Create(filepath.Join(env.FileStoragePath, filename))
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer dst.Close()

	// Copy the uploaded file to the file on disk
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	return env.FileStoragePath + filename, nil
}

func DeleteFile(filepath string) error {

	err := os.Remove(filepath)
	if err != nil {
		return fmt.Errorf("failed to remove file: %v", err)
	}

	return nil
}

func RetrieveFile(filename string) (*os.File, error) {
	filepath := filepath.Join(env.FileStoragePath, filename)

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	return file, nil
}

// GenerateFilename generates a unique filename for a file
func GenerateFilename(originalFilename string) string {
	extension := filepath.Ext(originalFilename)
	filename := originalFilename[:len(originalFilename)-len(extension)]
	return fmt.Sprintf("%s-%d%s", filename, time.Now().UnixNano(), extension)
}

// Helper function to check if the uploaded file is an image
func IsImage(file multipart.File) (bool, error) {
	// Read the first 512 bytes of the file
	buffer := make([]byte, 512)
	_, err := io.ReadAtLeast(file, buffer, 512)
	if err != nil {
		return false, err
	}

	// Detect the file type based on its content
	filetype := http.DetectContentType(buffer)

	//reset thing
	file.Seek(0, io.SeekStart)
	switch filetype {
	case "image/jpeg", "image/jpg", "image/png", "image/gif":
		return true, nil
	default:
		return false, nil
	}
}
