package service

import (
	"bbs_api/infra"
	"bbs_api/openapi"
	"context"
	"errors"
	"net/http"
)

func NewBbsService() openapi.DefaultApiServicer {
	return &bbsService{}
}

type bbsService struct {
}

func (s *bbsService) BoardListGet(ctx context.Context) (openapi.ImplResponse, error) {
	svc := infra.NewBoardListRepository("http://menu.5ch.net")

	boardGroups := svc.GetBoardGroups()

	bgs := make([]openapi.BoardGroup, len(boardGroups))
	for i, bg := range boardGroups {
		bl := make([]openapi.Board, len(bg.BoardList))
		for j, b := range bg.BoardList {
			bl[j] = openapi.Board{
				Name:     b.Name,
				ServerId: string(b.ServerId),
				BoardId:  string(b.BoardId),
			}
		}

		bgs[i] = openapi.BoardGroup{
			Name:      bg.Name,
			BoardList: bl,
		}
	}

	return openapi.Response(
		http.StatusOK,
		openapi.BoardListResponse{
			LastUpdated: "no implementation",
			BoardGroups: bgs,
		},
	), nil
}

func (s *bbsService) ServerBoardIdThreadListGet(ctx context.Context, server string, boardId string) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ServerBoardIdThreadListGet method not implemented")
}

func (s *bbsService) ServerBoardIdThreadThreadIdGet(ctx context.Context, server string, boardId string, threadId string) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ServerBoardIdThreadThreadIdGet method not implemented")
}
