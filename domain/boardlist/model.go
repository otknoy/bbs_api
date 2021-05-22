package boardlist

import "bbs_api/domain"

type BoardGroups []BoardGroup

type BoardGroup struct {
	Name      string
	BoardList BoardList
}

type BoardList []Board

type Board struct {
	Name     string
	ServerId domain.ServerId
	BoardId  domain.BoardId
}
