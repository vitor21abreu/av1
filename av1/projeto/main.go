package main

import "projeto/internal"

func main() {
	router := internal.SetupRoutes()
	router.Run(":8080")
}
