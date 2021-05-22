package boardlist

type BoardListRepository interface {
	GetBoardGroups() BoardGroups
}
