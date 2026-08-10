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
		er := pkg.BadRequest(err.Error())
		c.JSON(400, er)
		return
	}

	//check if the email is valid
	ok := pkg.IsValidEmail(req.Email)
	if !ok {
		er := pkg.BadRequest("invalid email format")
		c.JSON(400, er)
		return
	}

	//check if the password is valid
	err = pkg.IsValidPassword(req.Password)
	if err != nil {
		er := pkg.BadRequest(err.Error())
		c.JSON(400, er)
		return
	}

	err = s.CreateUser(c.Request.Context(), req)
	if err != nil {
		er := pkg.InternalServerError(err.Error())
		c.JSON(500, er)
		return
	}

	c.JSON(201, gin.H{"message": "user created successfully"})

}
