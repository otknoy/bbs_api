package main

import (
	"bbs_api/openapi"
	"bbs_api/service"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	svc := service.NewBbsService()
	DefaultApiController := openapi.NewDefaultApiController(svc)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
