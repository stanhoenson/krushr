package handlers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/middleware"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/utils"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCategoriesRoutes(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.Authorization())
	env.InitializeEnvironment("../../.env")
	handlers.RegisterCategoryRoutes(r)
	database.InitializeDatabase("test.db", "test/", true)
	validators.InitializeValidators()
	populateDatabaseWithDummyCategoryData()

	t.Run("catgeories", func(t *testing.T) {
		t.Run("testGetAllCategories", func(t *testing.T) {
			testGetAllCategories(t, r)
		})
		t.Run("testDeleteCategory", func(t *testing.T) {
			testDeleteCategory(t, r)
		})
		t.Run("testDeleteCategoryInvalidID", func(t *testing.T) {
			testDeleteCategoryInvalidID(t, r)
		})
		t.Run("testPostCategory", func(t *testing.T) {
			testPostCategory(t, r)
		})
		t.Run("testPutCategory", func(t *testing.T) {
			testPutCategory(t, r)
		})
		t.Run("testPutCategoryNoRecord", func(t *testing.T) {
			testPutCategoryNoRecord(t, r)
		})
		t.Run("testPutCategoryDuplicateName", func(t *testing.T) {
			testPutCategoryDuplicateName(t, r)
		})
		t.Run("testPostCategoryDuplicateName", func(t *testing.T) {
			testPostCategoryDuplicateName(t, r)
		})
	})

	os.Remove("test/test.db")
	os.Remove("test")
}

func testGetAllCategories(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
	r.ServeHTTP(w, req)

	var categories []models.Category
	err := json.Unmarshal(w.Body.Bytes(), &categories)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	// 3 because the normal populateDatabase function also adds a category
	assert.Equal(t, 3, len(categories))
}

func testDeleteCategory(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	addCategoryToDatabase(models.Category{ID: 4, Name: "Building", Position: 4})

	var countBefore int
	database.Db.Raw("SELECT COUNT(*) FROM categories WHERE id = 4").Scan(&countBefore)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/categories/4", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var count int
	database.Db.Raw("SELECT COUNT(*) FROM categories WHERE id = 4").Scan(&count)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "4", w.Body.String())
	assert.Equal(t, 0, count)
	assert.NotEqual(t, countBefore, count)
}

func testDeleteCategoryInvalidID(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/categories/-10", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid ID parameter", response["error"])
}

func testPostCategory(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	postCategoryBody := models.PostCategoryBody{Name: "Hills", Position: 6}

	postCategoryBodyJson, _ := json.Marshal(postCategoryBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(postCategoryBodyJson))
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var createCategory models.Category
	err = json.Unmarshal(w.Body.Bytes(), &createCategory)

	if err != nil {
		t.Error(err)
	}

	var count int
	database.Db.Raw("SELECT COUNT(*) FROM categories WHERE id = ?", createCategory.ID).Scan(&count)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, count)
	assert.NotEmpty(t, createCategory.ID)
}

func testPutCategory(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	nameBefore := "Building"

	addCategoryToDatabase(models.Category{ID: 4, Name: nameBefore, Position: 4})

	putCategoryBody := models.PutCategoryBody{
		PostCategoryBody: models.PostCategoryBody{
			Name: "Hills", Position: 6,
		},
	}

	putCategoryBodyJson, _ := json.Marshal(putCategoryBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/categories/4", bytes.NewBuffer(putCategoryBodyJson))
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var updatedCategory models.Category
	err = json.Unmarshal(w.Body.Bytes(), &updatedCategory)

	if err != nil {
		t.Error(err)
	}

	var nameAfter string
	database.Db.Raw("SELECT name FROM categories WHERE id = ?", 4).Scan(&nameAfter)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, nameBefore, nameAfter)
	assert.NotEmpty(t, updatedCategory.ID)
	assert.Equal(t, uint(4), updatedCategory.ID)
	assert.Equal(t, nameAfter, updatedCategory.Name)
}

func testPutCategoryNoRecord(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	putCategoryBody := models.PutCategoryBody{
		PostCategoryBody: models.PostCategoryBody{
			Name: "Hills", Position: 6,
		},
	}

	putCategoryBodyJson, _ := json.Marshal(putCategoryBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/categories/99", bytes.NewBuffer(putCategoryBodyJson))
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Error updating Categoryrecord not found", response["error"])
}

func testPostCategoryDuplicateName(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	postCategoryBody := models.PostCategoryBody{Name: "Nature", Position: 6}

	postCategoryBodyJson, _ := json.Marshal(postCategoryBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(postCategoryBodyJson))
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Error creating CategoryUNIQUE constraint failed: categories.name", response["error"])
}

func testPutCategoryDuplicateName(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	nameBefore := "Street"

	addCategoryToDatabase(models.Category{ID: 5, Name: nameBefore, Position: 5})

	putCategoryBody := models.PutCategoryBody{
		PostCategoryBody: models.PostCategoryBody{
			Name: "Nature", Position: 7,
		},
	}

	putCategoryBodyJson, _ := json.Marshal(putCategoryBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/categories/5", bytes.NewBuffer(putCategoryBodyJson))
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
	}

	var nameAfter string
	database.Db.Raw("SELECT name FROM categories WHERE id = ?", 5).Scan(&nameAfter)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, nameBefore, nameAfter)
	assert.Equal(t, "Error updating CategoryUNIQUE constraint failed: categories.name", response["error"])
}

func populateDatabaseWithDummyCategoryData() {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(utils.Sha256(env.AdminPassword)), bcrypt.DefaultCost)
	database.Db.Save(&models.User{ID: 2, Email: "creator@creator.com", Password: string(passwordBytes), RoleID: 2})
	if err != nil {
		log.Fatal(err)
	}

	addCategoryToDatabase(models.Category{ID: 2, Name: "Nature", Position: 1})
	addCategoryToDatabase(models.Category{ID: 3, Name: "Food", Position: 2})
}

func addCategoryToDatabase(category models.Category) {
	result := database.Db.Save(&category)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
