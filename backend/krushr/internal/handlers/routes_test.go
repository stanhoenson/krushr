package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/middleware"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

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

// TODO prob some issues when .env variables are used somewhere?
func TestRoutesRoutes(t *testing.T) {
	// common setup
	r := gin.Default()
	r.Use(middleware.Authorization())
	handlers.RegisterRouteRoutes(r)
	database.InitializeDatabase("test.db", "test/")
	populateDatabaseWithDummyData()

	// parallel
	t.Run("routes", func(t *testing.T) {
		t.Run("testGetAllRoutes", func(t *testing.T) {
			testGetAllRoutes(t, r)
		})
		// t.run("testdeleteroutebyid", func(t *testing.t) {
		// 	testdeleteroutebyid(t, r)
		// })
		// t.Run("testDeleteRouteUnauthorized", func(t *testing.T) {
		// 	testDeleteRouteUnauthorized(t, r)
		// })
		// t.Run("testDeleteRouteByIDWithInvalidID", func(t *testing.T) {
		// 	testDeleteRouteByIDWithInvalidID(t, r)
		// })
		// t.Run("testGetRouteByIDWithInvalidID", func(t *testing.T) {
		// 	testDeleteRouteByIDWithInvalidID(t, r)
		// })
		// t.Run("testGetRouteByID", func(t *testing.T) {
		// 	testGetRouteByID(t, r)
		// })
	})

	// teardown
	// os.Remove("test/test.db")
	// os.Remove("test")
}

func testGetAllRoutes(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/routes", nil)
	r.ServeHTTP(w, req)

	var createdRoutes []models.Route
	err := json.Unmarshal(w.Body.Bytes(), &createdRoutes)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(createdRoutes))
}

func testCreatorGetAllRoutes(t *testing.T, r *gin.Engine) {
}

// TODO hashed password
func testAdminGetAllRoutes(t *testing.T, r *gin.Engine) {
	// user, err := services.CreateUser(&models.PostUserBody{"test@test.com", "test", 2}, database.database.Db)
	// if err != nil {
	// 	t.Error(err)
	// }
	// dummyRoute := createDummyRoute()
	// route, err := services.CreateRoute(&dummyRoute, user, database.database.Db)
}

func testAdminGetOwnUnpublishedRoute(t *testing.T, r *gin.Engine) {
}

func testAdminGetOthersUnpublishedRoute(t *testing.T, r *gin.Engine) {
}

func testAdminGetRoute(t *testing.T, r *gin.Engine) {
}

// func testDeleteRouteByID(t *testing.T, r *gin.Engine) {
// 	createdRoute := createDummyRoute()
// 	createdRouteJSON, _ := json.Marshal(createdRoute)

// 	createdUser := createDummyUser()
// 	expirationTime := time.Now().Add(constants.TokenValidityPeriod)
// 	jwt, _ := services.GenerateJWTWithUser(createdUser, expirationTime)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/routes/"+strconv.Itoa(int(createdRoute.ID)), nil)
// 	req.Header.Add("Authorization", jwt)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, string(createdRouteJSON), w.Body.String())
// }

// func testDeleteRouteUnauthorized(t *testing.T, r *gin.Engine) {
// 	createdRoute := createDummyRoute()
// 	createdRouteJSON, _ := json.Marshal(createdRoute)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/routes/"+strconv.Itoa(int(createdRoute.ID)), nil)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, string(createdRouteJSON), w.Body.String())
// }

// func testDeleteRouteByIDWithInvalidID(t *testing.T, r *gin.Engine) {
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/routes/3", nil)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 400, w.Code)
// 	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
// }

// // misschien is het wel netter als dit allemaal in 1 functie staat maar zou ook ARRANGE kunnen zijn(prob het beste wel om alles hier te doen want anders kan je niet garanderen dat een andere functie in de weg zit), ook hier een goed voorbeeld waar een postRouteBody goed zou werken
// func testGetRouteByID(t *testing.T, r *gin.Engine) {
// 	route := models.Route{
// 		Name: "test",
// 	}
// 	createdRoute, _ := repositories.CreateEntity(&route, database.database.Db)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/routes/"+strconv.Itoa(int(createdRoute.ID)), nil)
// 	r.ServeHTTP(w, req)
// 	var retrievedRoute models.Route
// 	err := json.Unmarshal(w.Body.Bytes(), &retrievedRoute)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	assert.Equal(t, 200, w.Code)
// 	assert.EqualValues(t, createdRoute.Name, retrievedRoute.Name)
// }

// func testGetRouteByIDWithNonExistentRoute(t *testing.T, r *gin.Engine) {
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/routes/3", nil)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 400, w.Code)
// 	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
// }

// // An ID is of type uint
// func testGetRouteByIDWithInvalidID(t *testing.T, r *gin.Engine) {
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/routes/-3", nil)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 400, w.Code)
// 	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
// }

// func testDeleteRouteByIDWithNonexistentRoute(t *testing.T, r *gin.Engine) {
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/routes/3", nil)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 400, w.Code)
// 	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
// }

// // func TestPutRoute(t *testing.T) {
// // 	r := gin.Default()
// // 	r.PUT("/routes", handlers.PutRoute)

// // 	updatedRoute := models.Route{}
// // 	payload, _ := json.Marshal(updatedRoute)

// // 	w := httptest.NewRecorder()
// // 	req, _ := http.NewRequest("PUT", "/routes", bytes.NewBuffer(payload))
// // 	r.ServeHTTP(w, req)

// // 	assert.Equal(t, 200, w.Code)
// // 	assert.Equal(t, "hello", w.Body.String())
// // }
