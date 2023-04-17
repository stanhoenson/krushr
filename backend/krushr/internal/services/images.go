package services

import (
	"fmt"
	"mime/multipart"

	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/filemanager"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"gorm.io/gorm"
)

func CreateImage(fileHeader *multipart.FileHeader) (image *models.Image, err error) {
	err = database.Db.Transaction(func(tx *gorm.DB) error {

		filePath, err := filemanager.StoreMulitpartImage(fileHeader)
		if err != nil {
			return err
		}
		return err

		newImage := &models.Image{Path: filePath}

		image, err = repositories.CreateEntity(newImage)
		if err != nil {
			return err
		}

		return nil
	})

	return

}
