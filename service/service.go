package service

import (
	"bbs_api/domain"
	"bbs_api/domain/boardlist"
	"bbs_api/domain/threadlist"
	"bbs_api/openapi"
	"context"
	"errors"
	"net/http"
)

func NewBbsService(
	blRepo boardlist.BoardListRepository,
	thRepo threadlist.ThreadListRepository,
) openapi.DefaultApiServicer {
	return &bbsService{blRepo, thRepo}
}

type bbsService struct {
	boardListRepository  boardlist.BoardListRepository
	threadListRepository threadlist.ThreadListRepository
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
	threadList := s.threadListRepository.GetThreadList(domain.ServerId(server), domain.ThreadId(boardId))

	return openapi.Response(
		http.StatusOK,
		openapi.ThreadListResponse{
			ThreadList: func(l threadlist.ThreadList) []openapi.ThreadListResponseThreadList {
				res := make([]openapi.ThreadListResponseThreadList, len(l))
				for i, t := range threadList {
					res[i] = openapi.ThreadListResponseThreadList{
						Id:   string(t.ThreadId),
						Name: t.Name,
					}
				}
				return res
			}(threadList),
		},
	), nil
}

func (s *bbsService) ServerBoardIdThreadThreadIdGet(ctx context.Context, server string, boardId string, threadId string) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ServerBoardIdThreadThreadIdGet method not implemented")
}
