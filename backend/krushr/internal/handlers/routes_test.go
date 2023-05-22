package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/middleware"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TODO prob some issues when .env variables are used somewhere?
func TestRoutesRoutes(t *testing.T) {
	// common setup
	r := gin.Default()
	r.Use(middleware.Authorization())
	handlers.RegisterRouteRoutes(r)
	database.InitializeDatabase("test.db", "test/")

	// parallel
	t.Run("routes", func(t *testing.T) {
		t.Run("testDeleteRouteByID", func(t *testing.T) {
			testDeleteRouteByID(t, r)
		})
		t.Run("testDeleteRouteUnauthorized", func(t *testing.T) {
			testDeleteRouteUnauthorized(t, r)
		})
		t.Run("testDeleteRouteByIDWithInvalidID", func(t *testing.T) {
			testDeleteRouteByIDWithInvalidID(t, r)
		})
		t.Run("testGetRouteByIDWithInvalidID", func(t *testing.T) {
			testDeleteRouteByIDWithInvalidID(t, r)
		})
		t.Run("testGetRouteByID", func(t *testing.T) {
			testGetRouteByID(t, r)
		})
	})

	// teardown
	os.Remove("test/test.db")
	os.Remove("test")
}

func createDummyRoute() {
	dummyRoute := &models.PostRouteBody{
		Name: "De Boswandeling",
		PointsOfInterest: []models.PostPointOfInterestBody{
			{Name: "Grote Boom", GetPointOfInterestBody.Longitude: 32.785944, Latitude: 78.843984},
			{Name: "Grote Boom", Longitude: 32.785944, Latitude: 78.843984},
		},
	}
}

// TODO hashed password
func testAdminGetAllRoutes(t *testing.T, r *gin.Engine) {
	user, err := services.CreateUser(&models.PostUserBody{"test@test.com", "test", 2}, database.Db)
	if err != nil {
		t.Error(err)
	}
	route, err := services.CreateRoute(&models.PostRouteBody{"De Boswandeling"})
}

func testDeleteRouteByID(t *testing.T, r *gin.Engine) {
	createdRoute := createDummyRoute()
	createdRouteJSON, _ := json.Marshal(createdRoute)

	createdUser := createDummyUser()
	expirationTime := time.Now().Add(constants.TokenValidityPeriod)
	jwt, _ := services.GenerateJWTWithUser(createdUser, expirationTime)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/routes/"+strconv.Itoa(int(createdRoute.ID)), nil)
	req.Header.Add("Authorization", jwt)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(createdRouteJSON), w.Body.String())
}

func testDeleteRouteUnauthorized(t *testing.T, r *gin.Engine) {
	createdRoute := createDummyRoute()
	createdRouteJSON, _ := json.Marshal(createdRoute)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/routes/"+strconv.Itoa(int(createdRoute.ID)), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(createdRouteJSON), w.Body.String())
}

func testDeleteRouteByIDWithInvalidID(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/routes/3", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
}

// misschien is het wel netter als dit allemaal in 1 functie staat maar zou ook ARRANGE kunnen zijn(prob het beste wel om alles hier te doen want anders kan je niet garanderen dat een andere functie in de weg zit), ook hier een goed voorbeeld waar een postRouteBody goed zou werken
func testGetRouteByID(t *testing.T, r *gin.Engine) {
	route := models.Route{
		Name: "test",
	}
	createdRoute, _ := repositories.CreateEntity(&route, database.Db)

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
