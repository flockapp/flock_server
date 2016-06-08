package main

import (
	"fmt"
	"github.com/flockapp/flock_server/controllers"
	"github.com/flockapp/flock_server/models"
	"net/http"
	"log"
)

func main() {
	if err := models.Setup(); err != nil {
		panic(err)
	}
	fmt.Println("Listening on port", models.Conf.Port[1:])
	if err := http.ListenAndServe(models.Conf.Port, controllers.GetRoutes()); err != nil {
		log.Fatal(err)
	}
}
