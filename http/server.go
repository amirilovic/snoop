package http

import "net/http"
import "fmt"
import "github.com/julienschmidt/httprouter"

// Server is a http server
type Server struct {
        Port string
        Host string
}

// StartServer starts the server
func StartServer(s Server, routes []Route) {
        router := initHTTPRouter(routes)
        addr := fmt.Sprintf("%v:%v", s.Host, s.Port)
        fmt.Printf("Listening at: %v", addr)
        http.ListenAndServe(addr, router)
}

func initHTTPRouter(routes []Route) *httprouter.Router {
        router := httprouter.New()
        for _, r := range routes {
                router.Handle(r.Method, r.Path, r.Handle)
        }

        return router
}
