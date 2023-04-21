package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"gorm.io/gorm"
)

func FirstOrCreateDetail(postDetailBody *models.PostDetailBody, tx *gorm.DB) (*models.Detail, error) {

	detail := models.Detail{
		Text: postDetailBody.Text,
	}

	createdDetail, err := repositories.FirstOrCreateEntity(&detail, &detail, tx)
	if err != nil {
		return nil, err
	}

	return createdDetail, nil

}
