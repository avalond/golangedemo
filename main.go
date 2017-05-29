package main

import (
	"fmt"
	"golangedemo/config"
)

func main() {
	SERVER_PORT := config.ConstonsKey("server_port")


	fmt.Println("hello golange "+ SERVER_PORT)
}
