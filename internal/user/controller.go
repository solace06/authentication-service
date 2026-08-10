package user

import (
	"github.com/gin-gonic/gin"
	"github.com/solace06/auth-service/pkg"
)

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/signup", s.RegisterUser)
}

func (s *Scope) RegisterUser(c *gin.Context) {
	var req CreateUserRequest

	//validate request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//check if the email is valid
	ok := pkg.IsValidEmail(req.Email)
	if !ok {
		c.JSON(400, gin.H{"error": "invalid email format"})
		return
	}

	//check if the password is valid
	err = pkg.IsValidPassword(req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = s.CreateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "user created successfully"})

}
