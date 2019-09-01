package main

import (
	"teste/routes"
)

func main() {

	e := routes.Index()

	e.Start(":3000")
}
