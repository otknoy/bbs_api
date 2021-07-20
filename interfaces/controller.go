package interfaces

import (
	"bbs_api/domain"
	"bbs_api/interfaces/response"
	"bbs_api/openapi"
	"bbs_api/service"
	"context"
	"net/http"
)

func NewBbsController(bbsService service.BbsService) openapi.DefaultApiServicer {
	return &bbsController{bbsService}
}

type bbsController struct {
	bbsService service.BbsService
}

func (s *bbsController) BoardListGet(ctx context.Context) (openapi.ImplResponse, error) {
	boardGroups := s.bbsService.GetBoardGroups(ctx)

	return openapi.Response(
		http.StatusOK,
		response.CreateBoardListResponse(boardGroups),
	), nil
}

func (s *bbsController) ServerIdBoardIdThreadListGet(ctx context.Context, serverId string, boardId string) (openapi.ImplResponse, error) {
	threadList := s.bbsService.GetThreadList(ctx, domain.ServerId(serverId), domain.BoardId(boardId))

	return openapi.Response(
		http.StatusOK,
		response.CreateThreadListResponse(threadList),
	), nil
}

func (s *bbsController) ServerIdBoardIdThreadThreadIdGet(ctx context.Context, serverId string, boardId string, threadId string) (openapi.ImplResponse, error) {
	thread := s.bbsService.GetThread(ctx, domain.ServerId(serverId), domain.BoardId(boardId), domain.ThreadId(threadId))

	return openapi.Response(
		http.StatusOK,
		response.CreateThreadResponse(thread),
	), nil
}
