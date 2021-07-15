package main

import (
	"bbs_api/infra"
	"bbs_api/infra/bbsclient"
	"bbs_api/interfaces"
	"bbs_api/openapi"
	"bbs_api/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := func() http.Handler {
		ub := bbsclient.NewUrlBuilder("5ch.net")

		defaultApiController := openapi.NewDefaultApiController(
			interfaces.NewBbsController(
				service.NewBbsService(
					infra.NewBoardListRepository(ub),
					infra.NewThreadListRepository(ub),
					infra.NewThreadRepository(ub),
				),
			),
		)

		return openapi.NewRouter(defaultApiController)
	}()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: handler,
	}

	log.Println("server start")
	defer log.Println("server stop")

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
