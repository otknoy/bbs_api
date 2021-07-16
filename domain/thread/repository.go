package thread

import (
	"bbs_api/domain"
	"context"
)

type ThreadRepository interface {
	GetThread(context.Context, domain.ServerId, domain.BoardId, domain.ThreadId) *Thread
}
