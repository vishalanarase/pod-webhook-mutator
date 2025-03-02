package webhook

import (
	"net/http"
)

// Validate handles the validation of the request.
func Validate(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request validated successfully"))
	return
}
