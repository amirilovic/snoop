package ping

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	h "github.com/amirilovic/snoop/http"
)

func Routes() []h.Route {
	var routes []h.Route
	routes = append(routes, h.Route{Method: "GET", Path: "/ping", Handle: ping})
	return routes
}

func ping(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	type response struct {
		Message string `json:message`
	}

	json.NewEncoder(w).Encode(&response{Message: "PONG"})
}
