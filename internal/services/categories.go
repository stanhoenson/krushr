package services

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"gorm.io/gorm"
)

func FirstOrCreateCategory(postCategoryBody *models.PostCategoryBody, tx *gorm.DB) (*models.Category, error) {
	category := models.Category{
		Name:     postCategoryBody.Name,
		Position: postCategoryBody.Position,
	}

	createdCategory, err := repositories.FirstOrCreateEntity(&category, &models.Category{Name: postCategoryBody.Name}, tx)
	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}

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

func UpdateCategory(ID uint, putCategoryBody *models.PutCategoryBody) (*models.Category, error) {
	category, err := repositories.GetEntityByID[models.Category](ID, database.Db)
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
