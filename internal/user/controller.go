package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solace06/auth-service/pkg"
	"github.com/solace06/auth-service/problem"
)

func UserRoutes(router *gin.RouterGroup) {
	router.POST("", s.RegisterUser)
}

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", s.UserLogin)
}

func (s *Scope) RegisterUser(c *gin.Context) {
	var req CreateUserRequest

	//validate request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		er := problem.BadRequest(err.Error())
		c.JSON(http.StatusBadRequest, er)
		return
	}

	//check if the email is valid
	ok := pkg.IsValidEmail(req.Email)
	if !ok {
		er := problem.BadRequest("invalid email format")
		c.JSON(http.StatusBadRequest, er)
		return
	}

	//check if the password is valid
	err = pkg.IsValidPassword(req.Password)
	if err != nil {
		er := problem.BadRequest(err.Error())
		c.JSON(http.StatusBadRequest, er)
		return
	}

	err = s.CreateUser(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "user already exists" {
			er := problem.Conflict(err.Error())
			c.JSON(http.StatusConflict, er)
			return
		}
		er := problem.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, er)
		return
	}

	response := Response{
		Data:    nil,
		Message: "user created successfully",
	}

	c.JSON(http.StatusCreated, response)

}

func (s *Scope) UserLogin(c *gin.Context) {
	var req UserLoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		er := problem.BadRequest("invalid request")
		c.JSON(http.StatusBadRequest, er)
		return
	}

	if req.Email == "" || req.Password == "" {
		er := problem.BadRequest("required fields missing")
		c.JSON(http.StatusBadRequest, er)
		return
	}

	ok := pkg.IsValidEmail(req.Email)
	if !ok {
		er := problem.BadRequest("invalid email format")
		c.JSON(http.StatusBadRequest, er)
		return
	}

	err = pkg.IsValidPassword(req.Password)
	if err != nil {
		er := problem.BadRequest(err.Error())
		c.JSON(http.StatusBadRequest, er)
		return
	}

	token, err := s.AuthenticateUser(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "user not found" || err.Error() == "incorrect password" {
			er := problem.Unauthorized(err.Error())
			c.JSON(http.StatusUnauthorized, er)
			return
		}
		er := problem.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, er)
		return
	}

	resp := &Response{
		Data:    token,
		Message: "login successful",
	}

	c.JSON(http.StatusOK, resp)
}
