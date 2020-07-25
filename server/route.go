package server

import (
	"encoding/json"
	"github.com/pascencio/keeown-api/secret"
	"net/http"
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
