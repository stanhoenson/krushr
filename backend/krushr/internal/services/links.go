package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"gorm.io/gorm"
)

func FirstOrCreateLink(postLinkBody *models.PostLinkBody, tx *gorm.DB) (*models.Link, error) {

	link := models.Link{
		URL: postLinkBody.URL,
	}

	createdLink, err := repositories.FirstOrCreateEntity(&link, &link, tx)
	if err != nil {
		return nil, err
	}

	return createdLink, nil

}
