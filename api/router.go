package api

import (
	"github.com/gin-gonic/gin"
	"github.com/solace06/auth-service/user"
)

func NewRouter() *gin.Engine {

	r := gin.Default()

	r1 := r.Group("/api/v1/users")
	r2 := r.Group("/api/v1")

	r1.Use(RateLimit())
	user.UserRoutes(r1)

	r2.Use(LoginRateLimit())
	user.AuthRoutes(r2)

	return r
}
