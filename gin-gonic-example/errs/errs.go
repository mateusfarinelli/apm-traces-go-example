package errs

import "net/http"

type Err struct {
	StatusCode int    `json:"status"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

func NotFound(message string) *Err {
	return &Err{
		StatusCode: http.StatusNotFound,
		Error: "Not found",
		Message: message,
	}
}

func InternalServerError(message string) *Err {
	return &Err{
		StatusCode: http.StatusInternalServerError,
		Error: "Internal Server Error",
		Message: message,
	}
}

func BadRequest(message string) *Err {
	return &Err{
		StatusCode: http.StatusBadRequest,
		Error: "Bad Request",
		Message: message,
	}
}

func Unauthorized(message string) *Err {
	return &Err{
		StatusCode: http.StatusUnauthorized,
		Error: "Unauthorized",
		Message: message,
	}
}

func Conflict(message string) *Err {
	return &Err{
		StatusCode: http.StatusConflict,
		Error: "Conflict",
		Message: message,
	}
}