package main

import (
	"net/http"

	"github.com/rokaskuciauskas/goBasics/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
