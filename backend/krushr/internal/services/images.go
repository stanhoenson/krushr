package services

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/filemanager"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func CreateImage(fileHeader *multipart.FileHeader) (*models.Image, error) {
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

func DeleteImage(ID uint) (uint, error) {
	tx := database.Db.Begin()
	image, err := repositories.GetEntityByID[models.Image](ID, tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = filemanager.DeleteFile(image.Path)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = repositories.DeleteEntity(image, tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return image.ID, nil
}

func GetImageFile(ID uint) (*os.File, error) {
	image, err := repositories.GetEntityByID[models.Image](ID, database.Db)
	if err != nil {
		return nil, err
	}

	file, err := filemanager.RetrieveFile(image.Path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
