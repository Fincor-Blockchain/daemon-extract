package main

import (
	"fmt"

	"github.com/fincor-Blockchain/daemon-extract/routes"
)

func main() {
	fmt.Println("AI  training Server Is Running...")
	routes.StartRoutes()
}
