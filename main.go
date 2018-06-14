package main

import (
	"log"
	"net/http"
	"cricket_score/common"
	csr "cricket_score/router"
)

func main() {
	//db connection
	healthController := common.GetHealthController()

	router := csr.NewRouter(healthController)
	log.Fatal(http.ListenAndServe(":8080", router))
}