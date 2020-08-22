package responses

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse represents a failing response error
type ErrorResponse struct {
	HTTPStatusCode int    `json:"httpStatusCode"`
	Message        string `json:"message"`
}

// Render is a rendered response of an error
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// NewErrorResponse returns an initializer ErrorResponse with given error
func NewErrorResponse(statusCode int, err error) render.Renderer {
	return &ErrorResponse{
		HTTPStatusCode: statusCode,
		Message:        err.Error(),
	}
}
