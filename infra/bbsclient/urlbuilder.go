package bbsclient

import (
	"bbs_api/domain"
	"fmt"
)

type UrlBuilder interface {
	BuildBoardListUrl() string
	BuildThreadListUrl(domain.ServerId, domain.BoardId) string
	BuildThreadUrl(domain.ServerId, domain.BoardId, domain.ThreadId) string
}

func NewUrlBuilder(domainSuffix string) UrlBuilder {
	return &urlBuilder{domainSuffix}
}

type urlBuilder struct {
	domainSuffix string
}

func (b *urlBuilder) BuildBoardListUrl() string {
	return "http://menu." + b.domainSuffix + "/bbsmenu.html"
}

func (b *urlBuilder) BuildThreadListUrl(serverId domain.ServerId, boardId domain.BoardId) string {
	return fmt.Sprintf("http://%s.%s/%s/subback.html", serverId, b.domainSuffix, boardId)
}

func (b *urlBuilder) BuildThreadUrl(serverId domain.ServerId, boardId domain.BoardId, threadId domain.ThreadId) string {
	return fmt.Sprintf("http://%s.%s/test/read.cgi/%s/%s", serverId, b.domainSuffix, boardId, threadId)
}
