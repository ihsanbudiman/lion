package customerror

import (
	"net/http"
)

type HTTPError struct {
	Message    string
	StatusCode int
}

func (e *HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(message string, statusCode int) *HTTPError {
	return &HTTPError{
		Message:    message,
		StatusCode: statusCode,
	}
}
func WriteHTTPResponse(w http.ResponseWriter, err error) {
	httpErr, ok := err.(*HTTPError)
	if !ok {
		httpErr = NewHTTPError(err.Error(), http.StatusInternalServerError)
	}

	http.Error(w, httpErr.Message, httpErr.StatusCode)
}
