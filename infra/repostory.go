package infra

import (
	"bbs_api/domain"
	"bbs_api/domain/boardlist"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type BoardListRepository struct {
}

func NewBoardListRepository(host string) boardlist.BoardListRepository {
	return &boardListRepository{host}
}

type boardListRepository struct {
	host string
}

func (r *boardListRepository) GetBoardGroups() boardlist.BoardGroups {
	doc, _ := newShiftJISDocument(r.host + "/bbsmenu.html")

	m := make(map[string]boardlist.BoardList)
	l := make([]string, 0)
	groupName := ""
	doc.Find("html > body > font").Children().Each(func(_ int, s *goquery.Selection) {
		if s.Is("br") {
			return
		}

		if s.Is("b") {
			groupName = s.Text()
			l = append(l, s.Text())
		}

		if groupName == "" {
			return
		}

		if _, ok := m[groupName]; !ok {
			m[groupName] = make(boardlist.BoardList, 0)
		}

		if s.Is("a") {
			href, _ := s.Attr("href")
			u, _ := url.Parse(href)

			if u.Path == "" || u.Path == "/" {
				return
			}

			serverId := strings.Split(u.Host, ".")[0]
			boardId := u.Path[1 : len(u.Path)-1]

			m[groupName] = append(
				m[groupName],
				boardlist.Board{
					Name:     s.Text(),
					ServerId: domain.ServerId(serverId),
					BoardId:  domain.BoardId(boardId),
				},
			)
		}
	})

	bgs := boardlist.BoardGroups{}
	for _, k := range l {
		bgs = append(
			bgs,
			boardlist.BoardGroup{
				Name:      k,
				BoardList: m[k],
			},
		)
	}

	return boardlist.BoardGroups(bgs)
}
