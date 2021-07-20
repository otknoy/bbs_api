package response

import (
	"bbs_api/domain/boardlist"
	"bbs_api/domain/thread"
	"bbs_api/domain/threadlist"
	"bbs_api/openapi"
)

func CreateBoardListResponse(boardGroups boardlist.BoardGroups) *openapi.BoardListResponse {
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

	return &openapi.BoardListResponse{
		BoardGroups: bgs,
	}
}

func CreateThreadListResponse(threadList threadlist.ThreadList) *openapi.ThreadListResponse {
	tl := make([]openapi.ThreadListResponseThreadList, len(threadList))
	for i, t := range threadList {
		tl[i] = openapi.ThreadListResponseThreadList{
			Id:   string(t.ThreadId),
			Name: t.Name,
		}
	}

	return &openapi.ThreadListResponse{
		ThreadList: tl,
	}
}

func CreateThreadResponse(thread *thread.Thread) *openapi.ThreadResponse {
	l := make([]openapi.Comment, len(thread.CommentList))
	for i, c := range thread.CommentList {
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

	return &openapi.ThreadResponse{
		CommentList: l,
	}
}
