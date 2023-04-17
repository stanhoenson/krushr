package services

import (
	"fmt"
	"mime/multipart"

	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/filemanager"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func CreateImage(fileHeader *multipart.FileHeader) (image *models.Image, err error) {
	tx := database.Db.Begin()

	filePath, err := filemanager.StoreMulitpartImage(fileHeader)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	newImage := &models.Image{Path: filePath}

	createdImage, err := repositories.CreateEntity(newImage, tx)
	if err != nil {
		tx.Rollback()
		err := filemanager.DeleteFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to rollback file creation")
		}
		return nil, err
	}

	tx.Commit()
	return createdImage, nil

}
