package structs

import (
	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
	middlewares [] Middleware
}
