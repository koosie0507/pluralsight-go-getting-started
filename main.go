package main

import (
	"net/http"

	"github.com/koosie0507/pluralsight-go-getting-started/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
