package threadlist

type ThreadListRepository interface {
	GetThreadList() ThreadList
}
