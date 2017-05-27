package core

func NewRouter()*Router  {
	return  &Router{mux.NewRouter}
}