package api

import (
	ports "walls-user-service/internal/port"
)

// Httphander for the api
type HTTPHandler struct {
	userService ports.UserService
}

func NewHTTPHandler(
	countryService ports.UserService) *HTTPHandler {
	return &HTTPHandler{
		userService: countryService,
	}
}
