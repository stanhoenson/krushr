package handlers_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"mime/multipart"
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
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestImagesRoutes(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.Authorization())
	env.InitializeEnvironment("../../.env")
	handlers.RegisterImageRoutes(r)
	database.InitializeDatabase("test.db", "test/")
	populateDatabaseWithDummyDetailData()

	t.Run("images", func(t *testing.T) {
		t.Run("testPostImage", func(t *testing.T) {
			testPostImage(t, r)
		})
		// t.Run("testDeleteImage", func(t *testing.T) {
		// 	testDeleteImage(t, r)
		// })
		// t.Run("testGetImage", func(t *testing.T) {
		// 	testGetImage(t, r)
		// })
		// t.Run("testGetImageData", func(t *testing.T) {
		// 	testGetImageData(t, r)
		// })
	})

	os.Remove("test/test.db")
	os.Remove("test")
}

func testPostImage(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("creator@creator.com")

	signInBody := models.SignInBody{
		Email:    user.Email,
		Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	pngData := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg=="

	imageBytes, err := base64.StdEncoding.DecodeString(pngData)
	if err != nil {
		t.Fatal(err)
	}

	imageFileName := "testImage.png"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", imageFileName)
	if err != nil {
		t.Fatal(err)
	}
	_, err = part.Write(imageBytes)
	if err != nil {
		t.Fatal(err)
	}
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/images", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var image models.Image
	err = json.Unmarshal(w.Body.Bytes(), &image)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, image.Path, imageFileName)
}

func testGetImage(t *testing.T, r *gin.Engine) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/details", nil)
	r.ServeHTTP(w, req)

	var details []models.Detail
	err := json.Unmarshal(w.Body.Bytes(), &details)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 2, len(details))

}

func testGetImageData(t *testing.T, r *gin.Engine) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/details", nil)
	r.ServeHTTP(w, req)

	var details []models.Detail
	err := json.Unmarshal(w.Body.Bytes(), &details)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 2, len(details))

}

func testDeleteImage(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	addDetailToDatabase(models.Detail{ID: 3, Text: "some text"})

	var countBefore int
	database.Db.Raw("SELECT COUNT(*) FROM details WHERE id = 3").Scan(&countBefore)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/details/3", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var count int
	database.Db.Raw("SELECT COUNT(*) FROM categories WHERE id = 3").Scan(&count)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "3", w.Body.String())
	assert.Equal(t, 0, count)
	assert.NotEqual(t, countBefore, count)
}

func populateDatabaseWithDummyImageData() {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(utils.Sha256(env.AdminPassword)), bcrypt.DefaultCost)
	database.Db.Save(&models.User{ID: 2, Email: "creator@creator.com", Password: string(passwordBytes), RoleID: 2})
	if err != nil {
		log.Fatal(err)
	}

	addImageToDatabase(models.Detail{ID: 1, Text: "Very interesting"})
	addImageToDatabase(models.Detail{ID: 2, Text: "Hmmmmm"})

}

func addImageToDatabase(detail models.Detail) {
	result := database.Db.Save(&detail)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

}
