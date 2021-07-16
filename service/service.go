package service

import (
	"bbs_api/domain"
	"bbs_api/domain/boardlist"
	"bbs_api/domain/thread"
	"bbs_api/domain/threadlist"
	"context"
)

type BbsService interface {
	GetBoardGroups(context.Context) boardlist.BoardGroups
	GetThreadList(context.Context, domain.ServerId, domain.BoardId) threadlist.ThreadList
	GetThread(context.Context, domain.ServerId, domain.BoardId, domain.ThreadId) *thread.Thread
}

func NewBbsService(
	blRepo boardlist.BoardListRepository,
	tlRepo threadlist.ThreadListRepository,
	tRepo thread.ThreadRepository,
) BbsService {
	return &bbsService{blRepo, tlRepo, tRepo}
}

type bbsService struct {
	boardListRepository  boardlist.BoardListRepository
	threadListRepository threadlist.ThreadListRepository
	threadRepository     thread.ThreadRepository
}

func (s *bbsService) GetBoardGroups(ctx context.Context) boardlist.BoardGroups {
	return s.boardListRepository.GetBoardGroups(ctx)
}

func (s *bbsService) GetThreadList(ctx context.Context, serverId domain.ServerId, boardId domain.BoardId) threadlist.ThreadList {
	return s.threadListRepository.GetThreadList(ctx, serverId, boardId)
}

func (s *bbsService) GetThread(ctx context.Context, serverId domain.ServerId, boardId domain.BoardId, threadId domain.ThreadId) *thread.Thread {
	return s.threadRepository.GetThread(ctx, serverId, boardId, threadId)
}
