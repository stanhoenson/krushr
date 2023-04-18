package handlers_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/middleware"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var Db *gorm.DB

// TODO prob some issues when .env variables are used somewhere?
func TestRoutesRoutes(t *testing.T) {
	// common setup
	r := gin.Default()
	r.Use(middleware.Authorization())
	handlers.RegisterRouteRoutes(r)
	Db = database.InitializeDatabase("test.db")

	// sequentially
	// t.Run("routes", func(t *testing.T) {
	// 	testDeleteRouteByID(t, r)
	// 	testDeleteRouteByIDWithInvalidID(t, r)
	// 	// testDeleteRouteByIDWithNonexistentRoute(t, r)
	// 	testGetRouteByIDWithInvalidID(t, r)
	// 	// testGetRouteByIDWithNonExistentRoute(t, r)
	// 	testGetRouteByID(t, r)
	// })

	// parallel
	t.Run("routes", func(t *testing.T) {
		t.Run("testDeleteRouteByID", func(t *testing.T) {
			testDeleteRouteByID(t, r)
		})
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
	os.Remove("test.db")
}

func createDummyRoute() *models.Route {
	route := models.Route{
		Name: "De Boswandeling",
	}
	createdRoute, err := repositories.CreateEntity(&route, Db)
	if err != nil {
		log.Fatal(err)
	}
	return createdRoute
}

func createDummyUser() *models.User {
	user := models.User{
		Email:    "s.hoenson@protonmail.com",
		Password: "stanaap2",
		RoleID:   1,
	}
	createdUser, err := repositories.CreateEntity(&user, Db)
	if err != nil {
		log.Fatal(err)
	}
	return createdUser
}

func testDeleteRouteByID(t *testing.T, r *gin.Engine) {
	createdRoute := createDummyRoute()
	createdRouteJSON, _ := json.Marshal(createdRoute)

	createdUser := createDummyUser()
	jwt, _ := services.GenerateJWTWithUser(createdUser, 24*7)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/routes/"+strconv.Itoa(int(createdRoute.ID)), nil)
	req.Header.Add("Authorization", jwt)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(createdRouteJSON), w.Body.String())
}

// An ID should be of type uint
func testDeleteRouteByIDWithInvalidID(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/routes/-3", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
}

// misschien is het wel netter als dit allemaal in 1 functie staat maar zou ook ARRANGE kunnen zijn(prob het beste wel om alles hier te doen want anders kan je niet garanderen dat een andere functie in de weg zit), ook hier een goed voorbeeld waar een postRouteBody goed zou werken
func testGetRouteByID(t *testing.T, r *gin.Engine) {
	route := models.Route{
		Name: "test",
	}
	createdRoute, _ := repositories.CreateEntity(&route, Db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/routes/"+strconv.Itoa(int(createdRoute.ID)), nil)
	r.ServeHTTP(w, req)
	var retrievedRoute models.Route
	err := json.Unmarshal(w.Body.Bytes(), &retrievedRoute)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.EqualValues(t, createdRoute.Name, retrievedRoute.Name)
}

func testGetRouteByIDWithNonExistentRoute(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/routes/3", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
}

// An ID is of type uint
func testGetRouteByIDWithInvalidID(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/routes/-3", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
}

func testDeleteRouteByIDWithNonexistentRoute(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/routes/3", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
}

// func TestPutRoute(t *testing.T) {
// 	r := gin.Default()
// 	r.PUT("/routes", handlers.PutRoute)

// 	updatedRoute := models.Route{}
// 	payload, _ := json.Marshal(updatedRoute)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("PUT", "/routes", bytes.NewBuffer(payload))
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "hello", w.Body.String())
// }
