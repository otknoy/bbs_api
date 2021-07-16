package infra

import (
	"bbs_api/domain"
	"bbs_api/domain/boardlist"
	"bbs_api/domain/thread"
	"bbs_api/domain/threadlist"
	"bbs_api/infra/bbsclient"
	"context"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	_ "time/tzdata"
)

func NewBoardListRepository(ub bbsclient.UrlBuilder) boardlist.BoardListRepository {
	return &boardListRepository{ub}
}

type boardListRepository struct {
	ub bbsclient.UrlBuilder
}

func (r *boardListRepository) GetBoardGroups(ctx context.Context) boardlist.BoardGroups {
	doc, _ := bbsclient.NewShiftJISDocument(ctx, r.ub.BuildBoardListUrl())

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

func NewThreadListRepository(ub bbsclient.UrlBuilder) threadlist.ThreadListRepository {
	return &threadListRepository{ub}
}

type threadListRepository struct {
	ub bbsclient.UrlBuilder
}

func (r *threadListRepository) GetThreadList(ctx context.Context, serverId domain.ServerId, boardId domain.BoardId) threadlist.ThreadList {
	url := r.ub.BuildThreadListUrl(serverId, boardId)
	doc, _ := bbsclient.NewShiftJISDocument(ctx, url)

	threadList := make([]threadlist.Thread, 0)
	doc.Find("div > small#trad > a").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			return
		}

		tokens := strings.Split(href, "/")
		if len(tokens) != 2 {
			return
		}
		threadId := tokens[0]

		threadList = append(threadList, threadlist.Thread{
			ThreadId: domain.ThreadId(threadId),
			Name:     s.Text(),
		})
	})

	return threadlist.ThreadList(threadList)
}

func NewThreadRepository(ub bbsclient.UrlBuilder) thread.ThreadRepository {
	return &threadRepository{ub}
}

type threadRepository struct {
	ub bbsclient.UrlBuilder
}

var jst, _ = time.LoadLocation("Asia/Tokyo")

func (r *threadRepository) GetThread(ctx context.Context, serverId domain.ServerId, boardId domain.BoardId, threadId domain.ThreadId) *thread.Thread {
	url := r.ub.BuildThreadUrl(serverId, boardId, threadId)
	doc, _ := bbsclient.NewShiftJISDocument(ctx, url)

	l := make([]thread.Comment, 0)
	doc.Find("body > div > div.thread > div.post").Each(func(i int, s *goquery.Selection) {
		number, _ := strconv.ParseInt(s.Find("div.meta > span.number").Text(), 0, 32)
		name := s.Find("div.meta > span.name").Text()
		id := s.Find("div.meta > span.uid").Text()

		postedAt := func() time.Time {
			dayNames := map[string]string{
				"日": "Sun",
				"月": "Mon",
				"火": "Tue",
				"水": "Wed",
				"木": "Thu",
				"金": "Fri",
				"土": "Sat",
			}

			date := s.Find("div.meta > span.date").Text()
			for jp, en := range dayNames {
				date = strings.Replace(date, jp, en, 1)
			}

			postedAt, err := time.ParseInLocation(
				"2006/01/02(Mon) 15:04:05.99",
				date,
				jst,
			)
			if err != nil {
				log.Println(err)
			}

			return postedAt
		}()

		text := func() string {
			t, _ := s.Find("div.message > span.escaped").Html()

			lines := make([]string, 0)
			for _, l := range strings.Split(t, "<br/>") {
				l = strings.TrimLeft(l, " ")
				l = strings.TrimRight(l, " ")

				lines = append(lines, l)
			}

			return strings.Join(lines, "\n")
		}()

		l = append(
			l,
			thread.Comment{
				Meta: thread.Meta{
					Number:   uint(number),
					UserName: name,
					UserId:   id,
					PostedAt: postedAt,
				},
				Text: text,
			})
	})

	return &thread.Thread{
		CommentList: l,
	}
}
