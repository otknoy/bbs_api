package boardlist

import "context"

type BoardListRepository interface {
	GetBoardGroups(context.Context) BoardGroups
}
