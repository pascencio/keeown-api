package server

import (
	"encoding/json"
	"net/http"

	secret "github.com/pascencio/keeown-api/secret"
)

func secretRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			s := secret.GetSecret()
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(s)
		}
	}
}

// RouteHandler ...
func RouteHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/secret", secretRoute)
	return mux
}
