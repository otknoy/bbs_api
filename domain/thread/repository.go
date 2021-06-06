package thread

import "bbs_api/domain"

type ThreadRepository interface {
	GetThread(domain.ServerId, domain.BoardId, domain.ThreadId) *Thread
}
