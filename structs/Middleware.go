package structs

//Middleware holds a path and handler for a middleware
type Middleware struct {
	path    string
	handler http.Handler
}
