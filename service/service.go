package service

import (
	"bbs_api/domain"
	"bbs_api/domain/boardlist"
	"bbs_api/domain/thread"
	"bbs_api/domain/threadlist"
)

type BbsService interface {
	GetBoardGroups() boardlist.BoardGroups
	GetThreadList(domain.ServerId, domain.BoardId) threadlist.ThreadList
	GetThread(domain.ServerId, domain.BoardId, domain.ThreadId) *thread.Thread
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

func (s *bbsService) GetBoardGroups() boardlist.BoardGroups {
	return s.boardListRepository.GetBoardGroups()
}

func (s *bbsService) GetThreadList(serverId domain.ServerId, boardId domain.BoardId) threadlist.ThreadList {
	return s.threadListRepository.GetThreadList(serverId, boardId)
}

func (s *bbsService) GetThread(serverId domain.ServerId, boardId domain.BoardId, threadId domain.ThreadId) *thread.Thread {
	return s.threadRepository.GetThread(serverId, boardId, threadId)
}
