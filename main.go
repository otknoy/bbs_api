package main

import (
	"bbs_api/infra"
	"bbs_api/infra/bbsclient"
	"bbs_api/interfaces"
	"bbs_api/openapi"
	"bbs_api/service"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	ub := bbsclient.NewUrlBuilder("5ch.net")

	svc := interfaces.NewBbsController(
		service.NewBbsService(
			infra.NewBoardListRepository(ub),
			infra.NewThreadListRepository(ub),
			infra.NewThreadRepository(ub),
		),
	)
	defaultApiController := openapi.NewDefaultApiController(svc)

	router := openapi.NewRouter(defaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
