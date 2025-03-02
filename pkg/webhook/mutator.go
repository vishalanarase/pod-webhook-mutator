package webhook

import (
	"net/http"
)

// Mutate mutates the pod spec
func Mutate(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request mutated successfully"))
	return
}
