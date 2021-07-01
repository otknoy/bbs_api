package main

import (
	"bbs_api/domain"
	"bbs_api/infra"
	"bbs_api/interfaces"
	"bbs_api/openapi"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	svc := interfaces.NewBbsController(
		infra.NewBoardListRepository("http://menu.5ch.net"),
		infra.NewThreadListRepository(func(serverId domain.ServerId, boardId domain.BoardId) string {
			return fmt.Sprintf("http://%s.5ch.net/%s/subback.html", serverId, boardId)
		}),
		infra.NewThreadRepository(func(serverId domain.ServerId, boardId domain.BoardId, threadId domain.ThreadId) string {
			return fmt.Sprintf("http://%s.5ch.net/test/read.cgi/%s/%s", serverId, boardId, threadId)
		}),
	)
	defaultApiController := openapi.NewDefaultApiController(svc)

	router := openapi.NewRouter(defaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
