package filemanager

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/stanhoenson/krushr/internal/env"
)

func StoreMulitpartImage(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open fileHeader's file: %v", err)
	}
	defer file.Close()

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
	folderpath := filepath.Join(env.DataFolder, env.FileStorageFolder)
	filepath := filepath.Join(folderpath, filename)

	// Create the directory if it doesn't exist
	err = os.MkdirAll(folderpath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Create the file on disk
	dst, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer dst.Close()

	// TODO if something goes wrong we need to delete the file
	// Copy the uploaded file to the file on disk
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	return filepath, nil
}

func DeleteFile(filepath string) error {
	err := os.Remove(filepath)
	if err != nil {
		return fmt.Errorf("failed to remove file: %v", err)
	}

	return nil
}

func RetrieveFile(filename string) (*os.File, error) {
	filepath := filepath.Join(env.DataFolder, env.FileStorageFolder, filename)

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
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return false, err
	}

	// Detect the file type based on its content
	filetype := http.DetectContentType(fileContent)

	// reset thing
	file.Seek(0, io.SeekStart)
	switch filetype {
	case "image/jpeg", "image/jpg", "image/png", "image/gif":
		return true, nil
	default:
		return false, nil
	}
}
