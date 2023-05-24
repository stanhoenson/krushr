package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stanhoenson/krushr/internal/constants"
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

// TODO prob some issues when .env variables are used somewhere?
func TestRoutesRoutes(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.Authorization())
	handlers.RegisterRouteRoutes(r)
	database.InitializeDatabase("test.db", "test/")
	populateDatabaseWithDummyData()

	t.Run("routes", func(t *testing.T) {
		t.Run("testGetAllRoutes", func(t *testing.T) {
			testGetAllRoutes(t, r)
		})
		t.Run("testCreatorGetAllRoutes", func(t *testing.T) {
			testCreatorGetAllRoutes(t, r)
		})
		t.Run("testAdminGetAllRoutes", func(t *testing.T) {
			testAdminGetAllRoutes(t, r)
		})
		t.Run("testCreatorGetOwnUnpublishedRoute", func(t *testing.T) {
			testCreatorGetOwnUnpublishedRoute(t, r)
		})
		t.Run("testCreatorGetOthersUnpublishedRoute", func(t *testing.T) {
			testCreatorGetOthersUnpublishedRoute(t, r)
		})
	})

	os.Remove("test/test.db")
	os.Remove("test")
}

func testGetAllRoutes(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/routes", nil)
	r.ServeHTTP(w, req)

	var routes []models.Route
	err := json.Unmarshal(w.Body.Bytes(), &routes)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 2, len(routes))
}

func testCreatorGetAllRoutes(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("creator@creator.com")
	signInBody := models.SignInBody{Email: user.Email, Password: utils.Sha256(env.AdminPassword)}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/routes", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var routes []models.Route
	err = json.Unmarshal(w.Body.Bytes(), &routes)
	if err != nil {
		t.Fatal(err)
	}
	for _, route := range routes {
		if route.Status.Name == constants.UnpublishedStatusName && route.UserID != user.ID {
			t.Errorf("found unpublished route by someone else")
		}
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(routes))
}

func testAdminGetAllRoutes(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{Email: user.Email, Password: utils.Sha256(env.AdminPassword)}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/routes", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var routes []models.Route
	err = json.Unmarshal(w.Body.Bytes(), &routes)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 4, len(routes))
}

func testAdminGetRoute(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("admin@admin.com")
	signInBody := models.SignInBody{Email: user.Email, Password: utils.Sha256(env.AdminPassword)}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/route/1", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var route models.Route
	err = json.Unmarshal(w.Body.Bytes(), &route)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Example", route.Name)
}

func testCreatorGetOwnUnpublishedRoute(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("creator@creator.com")
	signInBody := models.SignInBody{Email: user.Email, Password: utils.Sha256(env.AdminPassword)}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/routes/4", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var route models.Route
	err = json.Unmarshal(w.Body.Bytes(), &route)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, user.Email, route.User.Email)
	assert.Equal(t, constants.UnpublishedStatusName, route.Status.Name)
}

func testCreatorGetOthersUnpublishedRoute(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("creator@creator.com")
	signInBody := models.SignInBody{Email: user.Email, Password: utils.Sha256(env.AdminPassword)}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/routes/2", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func populateDatabaseWithDummyData() {
	// Create roles
	adminRole := models.Role{ID: 1, Name: constants.AdminRoleName}
	result := database.Db.Save(&adminRole)
	if result.Error != nil {
		// Handle error
	}

	creatorRole := models.Role{ID: 2, Name: constants.CreatorRoleName}
	result = database.Db.Save(&creatorRole)
	if result.Error != nil {
		// Handle error
	}

	// Create creator user
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(utils.Sha256(env.AdminPassword)), bcrypt.DefaultCost)
	database.Db.Save(&models.User{ID: 2, Email: "creator@creator.com", Password: string(passwordBytes), RoleID: 2})
	if err != nil {
	}

	// CREATE ADMIN PUBLISHED ROUTE
	route := models.Route{
		ID:       1,
		Name:     "Example Route",
		Distance: 10.5,
		UserID:   1,
		StatusID: 1,
	}

	// Create associations
	image := models.Image{ID: 1, Path: "example/image.jpg"}
	route.Images = []*models.Image{&image}

	detail := models.Detail{ID: 1, Text: "Example Detail"}
	route.Details = []*models.Detail{&detail}

	link := models.Link{ID: 1, Text: "Example Link", URL: "https://example.com"}
	route.Links = []*models.Link{&link}

	category := models.Category{ID: 1, Name: "Example Category", Position: 1}
	route.Categories = []*models.Category{&category}

	poi := models.PointOfInterest{
		ID:        1,
		Name:      "Example POI",
		Longitude: 123.456,
		Latitude:  78.90,
		UserID:    1,
		Support:   false,
	}
	route.PointsOfInterest = []*models.PointOfInterest{&poi}

	// Save route and associations
	result = database.Db.Save(&route)
	if result.Error != nil {
		// Handle error
	}

	// Create many-to-many relationships
	err = database.Db.Model(&route).Association("Images").Append(&image)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Details").Append(&detail)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Links").Append(&link)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Categories").Append(&category)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("PointsOfInterest").Append(&poi)
	if err != nil {
		// Handle error
	}

	// Create many-to-many associations for points of interest
	err = database.Db.Model(&poi).Association("Images").Append(&image)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Details").Append(&detail)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Links").Append(&link)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Categories").Append(&category)
	if err != nil {
		// Handle error
	}

	// CREATE ADMIN UNPUBLISHED ROUTE
	route = models.Route{
		ID:       2,
		Name:     "Example Route 2",
		Distance: 10.5,
		UserID:   1,
		StatusID: 2,
	}

	// Create associations
	image = models.Image{ID: 2, Path: "example/image.jpg"}
	route.Images = []*models.Image{&image}

	detail = models.Detail{ID: 2, Text: "Example Detail"}
	route.Details = []*models.Detail{&detail}

	link = models.Link{ID: 2, Text: "Example Link", URL: "https://example.com"}
	route.Links = []*models.Link{&link}

	category = models.Category{ID: 2, Name: "Example Category", Position: 2}
	route.Categories = []*models.Category{&category}

	poi = models.PointOfInterest{
		ID:        2,
		Name:      "Example POI",
		Longitude: 123.456,
		Latitude:  78.90,
		UserID:    1,
		Support:   false,
	}
	route.PointsOfInterest = []*models.PointOfInterest{&poi}

	// Save route and associations
	result = database.Db.Save(&route)
	if result.Error != nil {
		// Handle error
	}

	// Create many-to-many relationships
	err = database.Db.Model(&route).Association("Images").Append(&image)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Details").Append(&detail)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Links").Append(&link)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Categories").Append(&category)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("PointsOfInterest").Append(&poi)
	if err != nil {
		// Handle error
	}

	// Create many-to-many associations for points of interest
	err = database.Db.Model(&poi).Association("Images").Append(&image)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Details").Append(&detail)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Links").Append(&link)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Categories").Append(&category)
	if err != nil {
		// Handle error
	}

	// CREATE CREATOR PUBLISHED ROUTE
	route = models.Route{
		ID:       3,
		Name:     "Example Route 3",
		Distance: 10.5,
		UserID:   2,
		StatusID: 1,
	}

	// Create associations
	image = models.Image{ID: 3, Path: "example/image.jpg"}
	route.Images = []*models.Image{&image}

	detail = models.Detail{ID: 3, Text: "Example Detail"}
	route.Details = []*models.Detail{&detail}

	link = models.Link{ID: 3, Text: "Example Link", URL: "https://example.com"}
	route.Links = []*models.Link{&link}

	category = models.Category{ID: 3, Name: "Example Category", Position: 3}
	route.Categories = []*models.Category{&category}

	poi = models.PointOfInterest{
		ID:        3,
		Name:      "Example POI",
		Longitude: 123.456,
		Latitude:  78.90,
		UserID:    2,
		Support:   false,
	}
	route.PointsOfInterest = []*models.PointOfInterest{&poi}

	// Save route and associations
	result = database.Db.Save(&route)
	if result.Error != nil {
		// Handle error
	}

	// Create many-to-many relationships
	err = database.Db.Model(&route).Association("Images").Append(&image)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Details").Append(&detail)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Links").Append(&link)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Categories").Append(&category)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("PointsOfInterest").Append(&poi)
	if err != nil {
		// Handle error
	}

	// Create many-to-many associations for points of interest
	err = database.Db.Model(&poi).Association("Images").Append(&image)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Details").Append(&detail)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Links").Append(&link)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Categories").Append(&category)
	if err != nil {
		// Handle error
	}

	// CREATE CREATOR UNPUBLISHED ROUTE
	route = models.Route{
		ID:       4,
		Name:     "Example Route 4",
		Distance: 10.5,
		UserID:   2,
		StatusID: 2,
	}

	// Create associations
	image = models.Image{ID: 4, Path: "example/image.jpg"}
	route.Images = []*models.Image{&image}

	detail = models.Detail{ID: 4, Text: "Example Detail"}
	route.Details = []*models.Detail{&detail}

	link = models.Link{ID: 4, Text: "Example Link", URL: "https://example.com"}
	route.Links = []*models.Link{&link}

	category = models.Category{ID: 4, Name: "Example Category", Position: 4}
	route.Categories = []*models.Category{&category}

	poi = models.PointOfInterest{
		ID:        4,
		Name:      "Example POI",
		Longitude: 123.456,
		Latitude:  78.90,
		UserID:    2,
		Support:   false,
	}
	route.PointsOfInterest = []*models.PointOfInterest{&poi}

	// Save route and associations
	result = database.Db.Save(&route)
	if result.Error != nil {
		// Handle error
	}

	// Create many-to-many relationships
	err = database.Db.Model(&route).Association("Images").Append(&image)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Details").Append(&detail)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Links").Append(&link)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("Categories").Append(&category)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&route).Association("PointsOfInterest").Append(&poi)
	if err != nil {
		// Handle error
	}

	// Create many-to-many associations for points of interest
	err = database.Db.Model(&poi).Association("Images").Append(&image)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Details").Append(&detail)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Links").Append(&link)
	if err != nil {
		// Handle error
	}

	err = database.Db.Model(&poi).Association("Categories").Append(&category)
	if err != nil {
		// Handle error
	}
}
