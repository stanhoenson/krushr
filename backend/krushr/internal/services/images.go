package services

import (
	"github.com/stanhoenson/krushr/internal/filemanager"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"fmt"
	"gorm.io/gorm"
	"mime/multipart"
)

func CreateImage(fileHeader *multipart.FileHeader, tx *gorm.DB) (*models.Image, error) {

	filePath, err := filemanager.StoreMulitpartImage(fileHeader)
	if err != nil {
		return nil, err
	}

	newImage := &models.Image{Path: filePath}

	createdImage, err := repositories.CreateEntity(newImage, tx)
	if err != nil {
		err := filemanager.DeleteFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to rollback file creation")
		}
		return nil, err
	}

	return createdImage, nil
}

func DeleteImage(ID uint, tx *gorm.DB) (uint, error) {
	image, err := repositories.GetEntityByID[models.Image](ID, tx)
	if err != nil {
		return 0, err
	}

	err = filemanager.DeleteFile(image.Path)
	if err != nil {
		return 0, err
	}

	_, err = repositories.DeleteEntity(image, tx)
	if err != nil {
		return 0, err
	}

	return image.ID, nil
}
