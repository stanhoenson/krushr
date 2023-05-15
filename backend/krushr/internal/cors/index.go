package cors

import (
	"github.com/stanhoenson/krushr/internal/env"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeCors(r *gin.Engine) {

	defaultCorsConfig := cors.DefaultConfig()
	defaultCorsConfig.AllowCredentials = true
	defaultCorsConfig.AllowOrigins = env.AllowedOrigins
	r.Use(cors.New(defaultCorsConfig))
}
