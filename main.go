package main

import (
	"bbs_api/infra"
	"bbs_api/openapi"
	"bbs_api/service"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	svc := service.NewBbsService(
		infra.NewBoardListRepository("http://menu.5ch.net"),
		infra.NewThreadListRepository(),
	)
	DefaultApiController := openapi.NewDefaultApiController(svc)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
