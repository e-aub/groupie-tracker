package global

import "net/http"

var (
	NotFoundErr         = Error{Code: http.StatusNotFound, Message: "Page not found!"}
	MethodNotAllowedErr = Error{Code: http.StatusMethodNotAllowed, Message: "Method not allowed!"}
	InternalServerErr   = Error{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
)
