package threadlist

import "bbs_api/domain"

type ThreadList []Thread

type Thread struct {
	ThreadId domain.ThreadId
	Name     string
}
