package interfaces

import (
	"bbs_api/domain"
	"bbs_api/domain/boardlist"
	"bbs_api/domain/thread"
	"bbs_api/domain/threadlist"
	"bbs_api/openapi"
	"context"
	"net/http"
)

func NewBbsController(
	blRepo boardlist.BoardListRepository,
	thRepo threadlist.ThreadListRepository,
	tRepo thread.ThreadRepository,
) openapi.DefaultApiServicer {
	return &bbsController{blRepo, thRepo, tRepo}
}

type bbsController struct {
	boardListRepository  boardlist.BoardListRepository
	threadListRepository threadlist.ThreadListRepository
	threadRepository     thread.ThreadRepository
}

func (s *bbsController) BoardListGet(ctx context.Context) (openapi.ImplResponse, error) {
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
			BoardGroups: bgs,
		},
	), nil
}

func (s *bbsController) ServerIdBoardIdThreadListGet(ctx context.Context, serverId string, boardId string) (openapi.ImplResponse, error) {
	threadList := s.threadListRepository.GetThreadList(domain.ServerId(serverId), domain.BoardId(boardId))

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

func (s *bbsController) ServerIdBoardIdThreadThreadIdGet(ctx context.Context, serverId string, boardId string, threadId string) (openapi.ImplResponse, error) {
	t := s.threadRepository.GetThread(domain.ServerId(serverId), domain.BoardId(boardId), domain.ThreadId(threadId))

	l := make([]openapi.Comment, len(t.CommentList))
	for i, c := range t.CommentList {
		l[i] = openapi.Comment{
			Meta: openapi.CommentMeta{
				Number:   int32(c.Meta.Number),
				UserName: c.Meta.UserName,
				UserId:   c.Meta.UserId,
				PostedAt: c.Meta.PostedAt,
			},
			Text: c.Text,
		}
	}

	return openapi.Response(
		http.StatusOK,
		openapi.ThreadResponse{
			CommentList: l,
		},
	), nil
}
