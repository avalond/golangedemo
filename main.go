package main

import (
	"fmt"
	"awesomeProject/config"

)

func main() {
	SERVER_PORT := config.ConstonsKey("server_port")


	fmt.Println("hello golange "+ SERVER_PORT)
}
