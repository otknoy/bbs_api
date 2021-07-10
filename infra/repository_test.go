package infra_test

import (
	"bbs_api/domain"
	"bbs_api/domain/boardlist"
	"bbs_api/domain/thread"
	"bbs_api/domain/threadlist"
	"bbs_api/infra"
	"bbs_api/infra/bbsclient"
	_ "embed"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var (
	//go:embed test_data/bbsmenu.html
	boardListHtml string

	//go:embed test_data/hayabusa9_news_subback.html
	threadListHtml string

	//go:embed test_data/hayabusa9_news_1622978925.html
	threadHtml string

	jst, _ = time.LoadLocation("Asia/Tokyo")
)

type urlBuilderMock struct {
	bbsclient.UrlBuilder
	MockBuildBoardListUrl  func() string
	MockBuildThreadListUrl func(domain.ServerId, domain.BoardId) string
	MockBuildThreadUrl     func(domain.ServerId, domain.BoardId, domain.ThreadId) string
}

func (m *urlBuilderMock) BuildBoardListUrl() string {
	return m.MockBuildBoardListUrl()
}

func (m *urlBuilderMock) BuildThreadListUrl(sId domain.ServerId, bId domain.BoardId) string {
	return m.MockBuildThreadListUrl(sId, bId)
}

func (m *urlBuilderMock) BuildThreadUrl(sId domain.ServerId, bId domain.BoardId, tId domain.ThreadId) string {
	return m.MockBuildThreadUrl(sId, bId, tId)
}

func TestBoardRepository_GetBoardGroups(t *testing.T) {
	s := NewStubServer(t, "/bbsmenu.html", boardListHtml)
	defer s.Close()

	r := infra.NewBoardListRepository(&urlBuilderMock{
		MockBuildBoardListUrl: func() string {
			return s.URL + "/bbsmenu.html"
		},
	})

	got := r.GetBoardGroups()

	if diff := cmp.Diff(
		got[0],
		boardlist.BoardGroup{
			Name: "地震",
			BoardList: []boardlist.Board{
				{
					Name:     "地震headline",
					ServerId: "headline",
					BoardId:  "bbynamazu",
				},
				{
					Name:     "地震速報",
					ServerId: "egg",
					BoardId:  "namazuplus",
				},
				{
					Name:     "臨時地震",
					ServerId: "mao",
					BoardId:  "eq",
				},
				{
					Name:     "臨時地震+",
					ServerId: "himawari",
					BoardId:  "eqplus",
				},
				{
					Name:     "緊急自然災害",
					ServerId: "rio2016",
					BoardId:  "lifeline",
				},
			},
		},
	); diff != "" {
		t.Errorf("differs: %v", diff)
	}
}

func TestBoardRepository_GetBoardGroups_number_of_boards(t *testing.T) {
	s := NewStubServer(t, "/bbsmenu.html", boardListHtml)
	defer s.Close()

	r := infra.NewBoardListRepository(&urlBuilderMock{
		MockBuildBoardListUrl: func() string {
			return s.URL + "/bbsmenu.html"
		},
	})

	got := r.GetBoardGroups()

	tests := []struct {
		index int
		name  string
		count int
	}{
		{0, "地震", 5},
		{1, "おすすめ", 7},
		{2, "特別企画", 1},
		{17, "家電製品", 23},
		{41, "ＰＣ等", 28},
		{46, "隔離", 1},
		{47, "運営案内", 2},
		{48, "BBSPINK", 116},
	}

	for _, tt := range tests {
		boardGroup := got[tt.index]

		if v := boardGroup.Name; v != tt.name {
			t.Errorf("name differs.\n got=%s\nwant=%s\n", v, tt.name)
		}

		if v := len(boardGroup.BoardList); v != tt.count {
			t.Errorf("number of board list differs.\n got=%d\nwant=%d\n", v, tt.count)
		}
	}
}

func TestThreadListRepository_GetThreadList(t *testing.T) {
	serverId := domain.ServerId("test-serverId")
	boardId := domain.BoardId("test-boardId")

	s := NewStubServer(t, "/test-boardId/subback.html", threadListHtml)
	defer s.Close()

	r := infra.NewThreadListRepository(&urlBuilderMock{
		MockBuildThreadListUrl: func(sId domain.ServerId, bId domain.BoardId) string {
			if sId != serverId || bId != boardId {
				t.Fatalf("invalid args of urlBuilderFunc: serverId=%s, boardId=%s", sId, bId)
			}
			return fmt.Sprintf("%s/%s/subback.html", s.URL, boardId)
		},
	})

	got := r.GetThreadList(serverId, boardId)

	if diff := cmp.Diff(
		got[:3],
		threadlist.ThreadList{
			{ThreadId: "1622633270", Name: "1: 【東京五輪】竹島表記問題が過激化…駐韓日本大使館付近で旭日旗を燃やした大学生３人逮捕 (1)"},
			{ThreadId: "1622632537", Name: "2: かっぱ寿司では (17)"},
			{ThreadId: "1622582170", Name: "3: 「身に付ける」ドラレコ登場　首にかけて360度をくまなく録画 (594)"},
		},
	); diff != "" {
		t.Errorf("differs: %v", diff)
	}

	if diff := cmp.Diff(
		got[431:434],
		threadlist.ThreadList{
			{ThreadId: "9245000000", Name: "432: ★ 5ちゃんねるからのお知らせ (3)"},
			{ThreadId: "9240200226", Name: "433: 【ご案内】新型コロナウイルスについて About COVID-19 (4)"},
			{ThreadId: "9240180801", Name: "434: 【すすコイン】期間限定キャンペーンのお知らせ！ (37)"},
		},
	); diff != "" {
		t.Errorf("differs: %v", diff)
	}

}

func TestThreadRepository_GetThread(t *testing.T) {
	serverId := domain.ServerId("test-serverId")
	boardId := domain.BoardId("test-boardId")
	threadId := domain.ThreadId("test-threadId")

	s := NewStubServer(t, "/test-boardId/thread/test-threadId", threadHtml)
	defer s.Close()

	r := infra.NewThreadRepository(&urlBuilderMock{
		MockBuildThreadUrl: func(sId domain.ServerId, bId domain.BoardId, tId domain.ThreadId) string {
			if sId != serverId || bId != boardId || tId != threadId {
				t.Fatalf("invalid args of urlBuilderFunc: serverId=%s, boardId=%s, threadId=%s", sId, bId, tId)
			}

			return fmt.Sprintf("%s/%s/thread/%s", s.URL, boardId, threadId)
		},
	})

	got := r.GetThread(serverId, boardId, threadId)

	if diff := cmp.Diff(
		got.CommentList[:3],
		[]thread.Comment{
			{
				Meta: thread.Meta{
					Number:   1,
					UserName: "セドナ(神奈川県) [ﾆﾀﾞ]",
					UserId:   "ID:8BHhTYQx0",
					PostedAt: time.Date(2021, 6, 6, 20, 28, 45, 980*1e6, jst),
				},
				Text: `<img src="//img.5ch.net/ico/1fu.gif"/>
大きな販路を失った台湾産パイナップルを応援しようと、愛媛県松山市職員有志が職員に
共同購入を呼び掛けたところ約９トン分の注文があり、３日に現物が到着した。
台湾から同市の学校給食へのパイナップル無償提供の話が進んでおり、お礼の意味もあるという。
同市は２０１４年に台北市と友好交流協定を結ぶなど、台湾と関係を深めている。

１９年７月に松山―台北の定期航空便が就航し、同年の松山市観光客推定では
国・地域別の外国人宿泊客数は台湾が６万３２００人と最も多かった。しかし、
新型コロナウイルス感染拡大で航空便は欠航が続いており、台湾の東京パラ柔道代表選手の松山での事前合宿は中止が決まった。


同課の窪田勝彦国際交流担当課長は「行き来は難しくなっているが、新型コロナ収束後を見据え、
さまざまな形で交流を深めたい」と話した。

<a href="http://jump.5ch.net/?https://www.ehime-np.co.jp/article/news202106040014" target="_blank">https://www.ehime-np.co.jp/article/news202106040014</a>`,
			},
			{
				Meta: thread.Meta{
					Number:   2,
					UserName: "環状星雲(光) [DE]",
					UserId:   "ID:KSqiCSi70",
					PostedAt: time.Date(2021, 6, 6, 20, 29, 31, 700*1e6, jst),
				},
				Text: "さよならと書いた毛蟹～",
			},
			{
				Meta: thread.Meta{
					Number:   3,
					UserName: "水メーザー天体(東京都) [EU]",
					UserId:   "ID:N7fBmzZc0",
					PostedAt: time.Date(2021, 6, 6, 20, 31, 8, 590*1e6, jst),
				},
				Text: "( ´∀`)二個買ったよ～",
			},
		},
	); diff != "" {
		t.Errorf("differs: %v", diff)
	}

}

func NewStubServer(t *testing.T, wantPath, response string) *httptest.Server {
	t.Helper()

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if path := r.URL.Path; path != wantPath {
			t.Fatalf("path must be %s, but %s", wantPath, path)
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, response)
	}))

	return s
}
