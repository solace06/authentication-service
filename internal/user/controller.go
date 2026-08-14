package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solace06/auth-service/pkg"
	"github.com/solace06/auth-service/problem"
)

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/signup", s.RegisterUser)
	r.POST("/login", s.UserLogin)
}

func (s *Scope) RegisterUser(c *gin.Context) {
	var req CreateUserRequest

	//validate request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		er := problem.BadRequest(err.Error())
		c.JSON(400, er)
		return
	}

	//check if the email is valid
	ok := pkg.IsValidEmail(req.Email)
	if !ok {
		er := problem.BadRequest("invalid email format")
		c.JSON(400, er)
		return
	}

	//check if the password is valid
	err = pkg.IsValidPassword(req.Password)
	if err != nil {
		er := problem.BadRequest(err.Error())
		c.JSON(400, er)
		return
	}

	err = s.CreateUser(c.Request.Context(), req)
	if err != nil {
		er := problem.InternalServerError(err.Error())
		c.JSON(500, er)
		return
	}

	response := Response{
		Data:    nil,
		Message: "user created successfully",
	}

	c.JSON(201, response)

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
		er:=problem.BadRequest("invalid email format")
		c.JSON(http.StatusBadRequest, er)
		return
	}

	err = pkg.IsValidPassword(req.Password)
	if err != nil {
		er:=problem.BadRequest(err.Error())
		c.JSON(http.StatusBadRequest, er)
		return
	}

	token, err:=s.AuthenticateUser(c.Request.Context(),req)
	if err != nil {
		er:=problem.InternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError,er)
		return
	}

	resp:=&Response{
		Data: token,
		Message: "Login Successful",
	}

	c.JSON(200, resp)
}
