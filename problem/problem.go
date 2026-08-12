package problem

import (
	"net/http"
)

type Problem struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail,omitempty"`
}

func New(status int, problemType, title, detail string) Problem {
	return Problem{
		Type:   problemType,
		Title:  title,
		Status: status,
		Detail: detail,
	}
}

func BadRequest(detail string) Problem {
	return New(
		http.StatusBadRequest,
		"https://your-api-domain.com/problems/bad-request",
		"Bad Request",
		detail,
	)
}

func Unauthorized(detail string) Problem {
	return New(
		http.StatusUnauthorized,
		"https://your-api-domain.com/problems/unauthorized",
		"Unauthorized",
		detail,
	)
}

func Forbidden(detail string) Problem {
	return New(
		http.StatusForbidden,
		"https://your-api-domain.com/problems/forbidden",
		"Forbidden",
		detail,
	)
}

func NotFound(detail string) Problem {
	return New(
		http.StatusNotFound,
		"https://your-api-domain.com/problems/not-found",
		"Not Found",
		detail,
	)
}

func Conflict(detail string) Problem {
	return New(
		http.StatusConflict,
		"https://your-api-domain.com/problems/conflict",
		"Conflict",
		detail,
	)
}

func InternalServerError(detail string) Problem {
	return New(
		http.StatusInternalServerError,
		"https://your-api-domain.com/problems/internal-server-error",
		"Internal Server Error",
		detail,
	)
}
