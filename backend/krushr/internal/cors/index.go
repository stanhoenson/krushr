package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeCors(r *gin.Engine) {

	// TODO fix cors things
	defaultCorsConfig := cors.DefaultConfig()
	defaultCorsConfig.AllowCredentials = true
	defaultCorsConfig.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:42305", "https://krushr.hoenson.xyz"}
	r.Use(cors.New(defaultCorsConfig))
}
