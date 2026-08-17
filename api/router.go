package api

import (
	"github.com/gin-gonic/gin"
	"github.com/solace06/auth-service/internal/user"
)

func NewRouter() *gin.Engine {

	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.Use(Auth())
	user.UserRoutes(v1)

	return r
}
