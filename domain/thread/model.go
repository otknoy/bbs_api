package thread

import "time"

type Thread struct {
	CommentList []Comment
}

type Comment struct {
	Meta Meta
	Text string
}

type Meta struct {
	Number   uint
	UserName string
	UserId   string
	PostedAt time.Time
}
