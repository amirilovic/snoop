package main

import (
	"github.com/amirilovic/snoop/http"
	"github.com/amirilovic/snoop/http/controllers/ping"
)

func main() {
	routes := ping.Routes()
	http.StartServer(http.Server{Port: "7777"}, routes)
}
