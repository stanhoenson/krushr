package services

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func CreateCategory(postCategoryBody *models.PostCategoryBody) (*models.Category, error) {

	category := models.Category{
		Name:     postCategoryBody.Name,
		Position: postCategoryBody.Position,
	}

	createdCategory, err := repositories.CreateEntity(&category, database.Db)
	if err != nil {
		return nil, err
	}

	return createdCategory, nil

}

func UpdateCategory(putCategoryBody *models.PutCategoryBody) (*models.Category, error) {

	category, err := repositories.GetEntity[models.Category](putCategoryBody.ID, database.Db)

	if err != nil {
		return nil, err
	}

	category.Name = putCategoryBody.Name
	category.Position = putCategoryBody.Position

	updatedCategory, err := repositories.UpdateEntity(category, database.Db)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil

}
