package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var (
	adminUserID         uint = 1
	creatorUserID       uint = 2
	publishedStatusID   uint = 1
	unpublishedStatusID uint = 2
)

func TestRoutesRoutes(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.Authorization())
	handlers.RegisterRouteRoutes(r)
	database.InitializeDatabase("test.db", "test/", true)
	validators.InitializeValidators()
	populateDatabaseWithDummyRoutesData()

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
		t.Run("testAdminGetRoute", func(t *testing.T) {
			testAdminGetRoute(t, r)
		})
		t.Run("testCreatorGetOwnUnpublishedRoute", func(t *testing.T) {
			testCreatorGetOwnUnpublishedRoute(t, r)
		})
		t.Run("testCreatorGetOthersUnpublishedRoute", func(t *testing.T) {
			testCreatorGetOthersUnpublishedRoute(t, r)
		})
		t.Run("testCreatorDeleteOwnRoute", func(t *testing.T) {
			testCreatorDeleteOwnRoute(t, r)
		})
		t.Run("testCreatorDeleteOthersRoute", func(t *testing.T) {
			testCreatorDeleteOthersRoute(t, r)
		})
		t.Run("testCreatorPostRoute", func(t *testing.T) {
			testCreatorPostRoute(t, r)
		})
		t.Run("testVisitorPostRoute", func(t *testing.T) {
			testVisitorPostRoute(t, r)
		})
		t.Run("testCreatorPutOwnRoute", func(t *testing.T) {
			testCreatorPutOwnRoute(t, r)
		})
	})

	os.Remove("test/test.db")
	os.Remove("test")
}

func testGetAllRoutes(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/routes", nil)
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
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/routes", nil)
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
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/routes", nil)
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
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/routes/1", nil)
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
	assert.Equal(t, "Example Route 1", route.Name)
}

func testCreatorGetOwnUnpublishedRoute(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("creator@creator.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/routes/4", nil)
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
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/routes/2", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func testCreatorDeleteOwnRoute(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("creator@creator.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	route := initializeRoute(5, creatorUserID, publishedStatusID, "Boris")
	addRouteToDatabase(route)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/routes/5", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var count int
	database.Db.Raw("SELECT COUNT(*) FROM routes WHERE id = 5").Scan(&count)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "5", w.Body.String())
	assert.Equal(t, 0, count)
}

func testCreatorDeleteOthersRoute(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("creator@creator.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/routes/1", nil)
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var count int
	database.Db.Raw("SELECT COUNT(*) FROM routes WHERE id = 1").Scan(&count)

	assert.Equal(t, 1, count)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func testCreatorPostRoute(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("creator@creator.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	route := initializeRoute(5, creatorUserID, publishedStatusID, "Jantje")
	postRouteBody := toPostRouteBody(route)
	postRouteBodyJSON, err := json.Marshal(postRouteBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/routes", bytes.NewBuffer(postRouteBodyJSON))
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	var count int
	database.Db.Raw("SELECT COUNT(*) FROM routes WHERE id = 5").Scan(&count)

	assert.Equal(t, 1, count)
	assert.Equal(t, http.StatusOK, w.Code)
}

func testVisitorPostRoute(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/routes", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func testCreatorPutOwnRoute(t *testing.T, r *gin.Engine) {
	user, _ := repositories.GetUserByEmail("creator@creator.com")
	signInBody := models.SignInBody{
		Email: user.Email, Password: utils.Sha256(env.AdminPassword),
	}
	token, err := services.Authenticate(&signInBody)
	if err != nil {
		t.Fatal(err)
	}

	route := initializeRoute(4, creatorUserID, publishedStatusID, "Brego")
	postRouteBody := toPostRouteBody(route)
	postRouteBodyJSON, err := json.Marshal(postRouteBody)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/routes/4", bytes.NewBuffer(postRouteBodyJSON))
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: token,
	}
	req.AddCookie(cookie)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// Helper functions
func initializeRoute(ID, userID, statusID uint, name string) (route models.Route) {
	route = models.Route{
		ID:       ID,
		Name:     name,
		Distance: 10.5,
		UserID:   userID,
		StatusID: statusID,
	}

	image := models.Image{ID: 1, Path: "example/image.jpg"}
	route.Images = []*models.Image{&image}

	detail := models.Detail{ID: 1, Text: "Example Detail"}
	route.Details = []*models.Detail{&detail}

	link := models.Link{ID: 1, Text: "Example Link", URL: "https://example.com"}
	route.Links = []*models.Link{&link}

	category := models.Category{ID: 1, Name: "Example Category", Position: 1}
	route.Categories = []*models.Category{&category}

	poi1 := models.PointOfInterest{
		ID:        1,
		Name:      "Example POI 1",
		Longitude: 123.456,
		Latitude:  78.90,
		UserID:    userID,
		Support:   false,
	}
	poi1.Images = []*models.Image{&image}
	poi1.Details = []*models.Detail{&detail}
	route.PointsOfInterest = []*models.PointOfInterest{&poi1}
	poi2 := models.PointOfInterest{
		ID:        2,
		Name:      "Example POI 2",
		Longitude: 123.456,
		Latitude:  78.90,
		UserID:    userID,
		Support:   false,
	}
	poi2.Images = []*models.Image{&image}
	poi2.Details = []*models.Detail{&detail}
	route.PointsOfInterest = append(route.PointsOfInterest, &poi2)

	return route
}

func toPostRouteBody(route models.Route) models.PostRouteBody {
	var postRouteBody models.PostRouteBody

	postRouteBody.Name = route.Name

	var postPOIBodies []models.PostPointOfInterestBody
	for _, poi := range route.PointsOfInterest {
		var postPOIBody models.PostPointOfInterestBody
		postPOIBody.Name = poi.Name
		postPOIBody.Latitude = poi.Latitude
		postPOIBody.Longitude = poi.Longitude
		for _, image := range poi.Images {
			postPOIBody.ImageIDs = append(postPOIBody.ImageIDs, image.ID)
		}
		for _, detail := range poi.Details {
			postDetailBody := models.PostDetailBody{
				Text: detail.Text,
			}
			postPOIBody.Details = append(postPOIBody.Details, postDetailBody)
		}
		postPOIBodies = append(postPOIBodies, postPOIBody)
	}
	postRouteBody.PointsOfInterest = postPOIBodies

	var imageIDs []uint
	for _, image := range route.Images {
		imageIDs = append(imageIDs, image.ID)
	}
	postRouteBody.ImageIDs = imageIDs

	var postDetailBodies []models.PostDetailBody
	for _, detail := range route.Details {
		var postDetailBody models.PostDetailBody
		postDetailBody.Text = detail.Text
		postDetailBodies = append(postDetailBodies, postDetailBody)
	}
	postRouteBody.Details = postDetailBodies

	var postLinkBodies []models.PostLinkBody
	for _, link := range route.Links {
		var postLinkBody models.PostLinkBody
		postLinkBody.Text = link.Text
		postLinkBody.URL = link.URL
		postLinkBodies = append(postLinkBodies, postLinkBody)
	}
	postRouteBody.Links = postLinkBodies

	var postCategoryBodies []models.PostCategoryBody
	for _, category := range route.Categories {
		var postCategoryBody models.PostCategoryBody
		postCategoryBody.Name = category.Name
		postCategoryBody.Position = category.ID
		postCategoryBodies = append(postCategoryBodies, postCategoryBody)
	}
	postRouteBody.Categories = postCategoryBodies

	postRouteBody.StatusID = route.StatusID

	return postRouteBody
}

func addRouteToDatabase(route models.Route) {
	// TODO handle errors
	result := database.Db.Save(&route)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	image := route.Images[0]
	detail := route.Details[0]
	link := route.Links[0]
	category := route.Categories[0]
	poi := route.PointsOfInterest[0]

	err := database.Db.Model(&route).Association("Images").Append(&image)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Db.Model(&route).Association("Details").Append(&detail)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Db.Model(&route).Association("Links").Append(&link)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Db.Model(&route).Association("Categories").Append(&category)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Db.Model(&route).Association("PointsOfInterest").Append(&poi)
	if err != nil {
		log.Fatal(err)
	}

	// Create many-to-many associations for points of interest
	err = database.Db.Model(&poi).Association("Images").Append(&image)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Db.Model(&poi).Association("Details").Append(&detail)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Db.Model(&poi).Association("Links").Append(&link)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Db.Model(&poi).Association("Categories").Append(&category)
	if err != nil {
		log.Fatal(err)
	}
}

func populateDatabaseWithDummyRoutesData() {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(utils.Sha256(env.AdminPassword)), bcrypt.DefaultCost)
	database.Db.Save(&models.User{ID: 2, Email: "creator@creator.com", Password: string(passwordBytes), RoleID: 2})
	if err != nil {
		log.Fatal(err)
	}

	route := initializeRoute(1, adminUserID, publishedStatusID, "Example Route 1")
	addRouteToDatabase(route)
	route = initializeRoute(2, adminUserID, unpublishedStatusID, "Example Route 2")
	addRouteToDatabase(route)
	route = initializeRoute(3, creatorUserID, publishedStatusID, "Example Route 3")
	addRouteToDatabase(route)
	route = initializeRoute(4, creatorUserID, unpublishedStatusID, "Example Route 4")
	addRouteToDatabase(route)
}
