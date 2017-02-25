package main

import (
	"github.com/amirilovic/snoop/http"
	"github.com/amirilovic/snoop/http/controllers/ping"
)

func main() {
	// Start server
	routes := ping.Routes()
	http.StartServer(http.Server{Port: "7777"}, routes)
}
