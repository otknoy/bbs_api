package main

import (
	"bbs_api/domain"
	"bbs_api/infra"
	"bbs_api/openapi"
	"bbs_api/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	svc := service.NewBbsService(
		infra.NewBoardListRepository("http://menu.5ch.net"),
		infra.NewThreadListRepository(func(serverId domain.ServerId, boardId domain.BoardId) string {
			return fmt.Sprintf("http://%s.5ch.net/%s/subback.html", serverId, boardId)
		}),
	)
	DefaultApiController := openapi.NewDefaultApiController(svc)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
