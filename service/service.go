package service

import (
	"bbs_api/domain/boardlist"
	"bbs_api/openapi"
	"context"
	"errors"
	"net/http"
)

func NewBbsService(blRepo boardlist.BoardListRepository) openapi.DefaultApiServicer {
	return &bbsService{blRepo}
}

type bbsService struct {
	boardListRepository boardlist.BoardListRepository
}

func (s *bbsService) BoardListGet(ctx context.Context) (openapi.ImplResponse, error) {
	boardGroups := s.boardListRepository.GetBoardGroups()

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
