package main

import (
	"go-sea-crm/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8080")
}