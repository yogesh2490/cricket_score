package main

import (
	"cricket_score/common"
	csr "cricket_score/router"
	"log"
	"net/http"
)

func main() {
	//db connection
	healthController := common.GetHealthController()

	router := csr.NewRouter(healthController)
	log.Fatal(http.ListenAndServe(":8080", router))
}
