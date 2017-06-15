package main

import (
	"net/http"
	"golangedemo/code/error"
)

func main() {
	setPort()
	setHandler()

}
func setPort() {
	const SERVER_PORT string = ":8088"

	println("start server at" + SERVER_PORT)
	//port
	e := http.ListenAndServe(SERVER_PORT, nil)
	if e != nil {
		panic(e)
	}

}

func setHandler() {
	http.HandleFunc("/", error.NotFoundHandler)
}
