package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pascencio/keeown-api/secret"
)

func getSecret(w http.ResponseWriter, r *http.Request) {
	s := secret.GetSecret()
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

// RouteHandler ...
func RouteHandler() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/secret", getSecret).
		Methods("GET")
	return r
}
