package threadlist

import "bbs_api/domain"

type ThreadListRepository interface {
	GetThreadList(domain.ServerId, domain.BoardId) ThreadList
}
