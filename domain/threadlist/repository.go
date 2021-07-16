package threadlist

import (
	"bbs_api/domain"
	"context"
)

type ThreadListRepository interface {
	GetThreadList(context.Context, domain.ServerId, domain.BoardId) ThreadList
}
