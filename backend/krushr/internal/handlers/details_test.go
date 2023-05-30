package handlers_test

import (
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
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestDetailsRoutes(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.Authorization())
	env.InitializeEnvironment("../../.env")
	handlers.RegisterDetailRoutes(r)
	database.InitializeDatabase("test.db", "test/")
	populateDatabaseWithDummyDetailData()

	t.Run("details", func(t *testing.T) {
		t.Run("testDeleteDetail", func(t *testing.T) {
			testDeleteDetail(t, r)
		})
		t.Run("testGetAllDetails", func(t *testing.T) {
			testGetAllDetails(t, r)
		})
	})

	os.Remove("test/test.db")
	os.Remove("test")
}

func testGetAllDetails(t *testing.T, r *gin.Engine) {

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

func testDeleteDetail(t *testing.T, r *gin.Engine) {
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

func populateDatabaseWithDummyDetailData() {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(utils.Sha256(env.AdminPassword)), bcrypt.DefaultCost)
	database.Db.Save(&models.User{ID: 2, Email: "creator@creator.com", Password: string(passwordBytes), RoleID: 2})
	if err != nil {
		log.Fatal(err)
	}

	addDetailToDatabase(models.Detail{ID: 1, Text: "Very interesting"})
	addDetailToDatabase(models.Detail{ID: 2, Text: "Hmmmmm"})

}

func addDetailToDatabase(detail models.Detail) {
	result := database.Db.Save(&detail)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

}
