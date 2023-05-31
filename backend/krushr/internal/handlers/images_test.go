package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"image"
	"image/color"
	"image/png"
	"io/ioutil"

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
		//this order matters
		t.Run("testPostImage", func(t *testing.T) {
			testPostImage(t, r)
		})
		t.Run("testGetImage", func(t *testing.T) {
			testGetImage(t, r)
		})
		t.Run("testGetImageData", func(t *testing.T) {
			testGetImageData(t, r)
		})
		t.Run("testGetImageDataFaultyID", func(t *testing.T) {
			testGetImageDataFaultyID(t, r)
		})
		t.Run("testGetImageDataNoRecord", func(t *testing.T) {
			testGetImageDataNoRecord(t, r)
		})
		t.Run("testDeleteImage", func(t *testing.T) {
			testDeleteImage(t, r)
		})
		t.Run("testPostImageFaultyImageFile", func(t *testing.T) {
			testPostImageFaultyImageFile(t, r)
		})
		t.Run("testPostImageNoMultiparForm", func(t *testing.T) {
			testPostImageNoMultiparForm(t, r)
		})
	})

	os.RemoveAll("test")
	os.RemoveAll("data")
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

	filePath, err := generateTempPNGWithPixel(color.Black, 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	imageFile, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	imageBytes, err := ioutil.ReadAll(imageFile)
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
	assert.Contains(t, image.Path, strings.Split(imageFileName, ".png")[0])
}
func testPostImageFaultyImageFile(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("creator@creator.com")

	signInBody := models.SignInBody{
		Email:    user.Email,
		Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	imageFileName := "testImage.png"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", imageFileName)
	_, err = part.Write([]byte{})
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

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Error creating image", response["error"])
}
func testPostImageNoMultiparForm(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("creator@creator.com")

	signInBody := models.SignInBody{
		Email:    user.Email,
		Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/images", bytes.NewBufferString("not a MultipartForm"))

	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Error parsing MultipartForm", response["error"])
}

func testGetImage(t *testing.T, r *gin.Engine) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/images/1", nil)
	r.ServeHTTP(w, req)

	var image models.Image
	err := json.Unmarshal(w.Body.Bytes(), &image)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, uint(1), image.ID)

}
func testGetImageDataFaultyID(t *testing.T, r *gin.Engine) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/imagedata/-10", nil)
	r.ServeHTTP(w, req)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid ID parameter", response["error"])

}

func testGetImageDataNoRecord(t *testing.T, r *gin.Engine) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/imagedata/10", nil)
	r.ServeHTTP(w, req)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Error retrieving image", response["error"])

}
func testGetImageData(t *testing.T, r *gin.Engine) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/imagedata/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.Bytes())

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

	var countBefore int
	database.Db.Raw("SELECT COUNT(*) FROM images ").Scan(&countBefore)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/images/1", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var count int
	database.Db.Raw("SELECT COUNT(*) FROM images").Scan(&count)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "1", w.Body.String())
	assert.Equal(t, 0, count)
	assert.NotEqual(t, countBefore, count)
}

func populateDatabaseWithDummyImageData() {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(utils.Sha256(env.AdminPassword)), bcrypt.DefaultCost)
	database.Db.Save(&models.User{ID: 2, Email: "creator@creator.com", Password: string(passwordBytes), RoleID: 2})
	if err != nil {
		log.Fatal(err)
	}

}

func addImageToDatabase(detail models.Detail) {
	result := database.Db.Save(&detail)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

}

func generateTempPNGWithPixel(color color.Color, width, height int) (string, error) {
	// Create a new RGBA image with the specified dimensions
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Set the pixel color for the entire image
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color)
		}
	}

	// Create a temporary file with a .png extension
	file, err := ioutil.TempFile("", "image-*.png")
	if err != nil {
		return "", err
	}

	// Save the image as PNG
	err = png.Encode(file, img)
	if err != nil {
		return "", err
	}

	// Get the temporary file path
	filePath := file.Name()

	return filePath, nil
}
