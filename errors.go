package groupme

import (
	"errors"
	"net/http"
)

// StatusEnhanceYourCalm is "returned when you are being rate limited. Chill the heck out."
const StatusEnhanceYourCalm = 420

// Errors returned by GroupMe.
var (
	ErrNotModified         = errors.New("304 Not Modified")
	ErrBadRequest          = errors.New("400 Bad Request")
	ErrUnauthorized        = errors.New("401 Unauthorized")
	ErrForbidden           = errors.New("403 Forbidden")
	ErrNotFound            = errors.New("404 Not Found")
	ErrEnhanceYourCalm     = errors.New("420 Enhance Your Calm")
	ErrInternalServerError = errors.New("500 Internal Server Error")
	ErrBadGateway          = errors.New("502 Bad Gateway")
	ErrServiceUnavailable  = errors.New("503 Service Unavailable")
)

// Meta is the error response from the GroupMe API.
type Meta struct {
	Code   int      `json:"code"`
	Errors []string `json:"errors"`
}

func parseError(statusCode int, status string) error {
	switch statusCode {
	case http.StatusNotModified:
		return ErrNotModified
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusNotFound:
		return ErrNotFound
	case StatusEnhanceYourCalm:
		return ErrEnhanceYourCalm
	case http.StatusInternalServerError:
		return ErrInternalServerError
	case http.StatusBadGateway:
		return ErrBadGateway
	case http.StatusServiceUnavailable:
		return ErrServiceUnavailable
	default:
		return errors.New(status)
	}
}
