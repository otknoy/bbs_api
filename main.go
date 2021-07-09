package main

import (
	"bbs_api/infra"
	ihttp "bbs_api/infra/http"
	"bbs_api/interfaces"
	"bbs_api/openapi"
	"bbs_api/service"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	ub := ihttp.NewUrlBuilder("5ch.net")

	svc := interfaces.NewBbsController(
		service.NewBbsService(
			infra.NewBoardListRepository("http://menu.5ch.net"),
			infra.NewThreadListRepository(ub.BuildGetThradListUrl),
			infra.NewThreadRepository(ub.BuildGetThreadUrl),
		),
	)
	defaultApiController := openapi.NewDefaultApiController(svc)

	router := openapi.NewRouter(defaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
