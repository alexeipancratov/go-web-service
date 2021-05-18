package main

import (
	"net/http"

	"github.com/alexeipancratov/web-service-advanced/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
