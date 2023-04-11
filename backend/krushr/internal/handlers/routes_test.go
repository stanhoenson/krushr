package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// TODO prob some issues when .env variables are used somewhere?
func TestRoutesRoutes(t *testing.T) {
	// common setup
	r := gin.Default()
	handlers.RoutesRoutes(r)
	db := database.InitializeDatabase("test.db")

	// sequentially
	t.Run("routes", func(t *testing.T) {
		testDeleteRouteByIDWithNonexistentRoute(t, r)
		testGetRouteByIDWithNonExistentRoute(t, r)
		testGetRouteByID(t, r, db)
	})

	// parallel
	// t.Run("routes", func(t *testing.T) {
	// 	t.Run("testDeleteRouteByIDWithInvalidID", func(t *testing.T) {
	// 		testDeleteRouteByIDWithInvalidID(t, r)
	// 	})
	// 	t.Run("testGetRouteByIDWithInvalidID", func(t *testing.T) {
	// 		testDeleteRouteByIDWithInvalidID(t, r)
	// 	})
	// 	t.Run("testGetRouteByID", func(t *testing.T) {
	// 		testGetRouteByID(t, r, db)
	// 	})
	// })

	// teardown
	os.Remove("test.db")
}

// misschien is het wel netter als dit allemaal in 1 functie staat maar zou ook ARRANGE kunnen zijn(prob het beste wel om alles hier te doen want anders kan je niet garanderen dat een andere functie in de weg zit), ook hier een goed voorbeeld waar een postRouteBody goed zou werken
func testGetRouteByID(t *testing.T, r *gin.Engine, db *gorm.DB) {
	route := models.Route{
		Title: "test",
	}
	createdRoute, _ := repositories.CreateEntity(&route)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/routes/"+strconv.Itoa(int(createdRoute.ID)), nil)
	r.ServeHTTP(w, req)
	var retrievedRoute models.Route
	err := json.Unmarshal(w.Body.Bytes(), &retrievedRoute)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.EqualValues(t, createdRoute.Title, retrievedRoute.Title)
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

// func TestDeleteRouteByID(t *testing.T) {
// 	r := gin.Default()
// 	r.DELETE("/routes/:id", handlers.DeleteRouteByID)

// 	route := models.Route{
// 		Title:    "test",
// 		StatusID: 1,
// 		UserID:   1,
// 	}
// 	repositories.CreateEntity(&route)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/routes/:1", nil)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "hello", w.Body.String())
// }

func testDeleteRouteByIDWithNonexistentRoute(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/routes/3", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"Invalid ID parameter\"}", w.Body.String())
}

// An ID is of type uint
func testDeleteRouteByIDWithInvalidID(t *testing.T, r *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/routes/-3", nil)
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
