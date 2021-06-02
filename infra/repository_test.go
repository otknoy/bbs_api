package infra_test

import (
	"bbs_api/domain"
	"bbs_api/domain/boardlist"
	"bbs_api/domain/threadlist"
	"bbs_api/infra"
	_ "embed"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	//go:embed test_data/bbsmenu.html
	boardListHtml string

	//go:embed test_data/hayabusa9_news_subback.html
	threadListHtml string
)

func TestBoardRepository_GetBoardGroups(t *testing.T) {
	s := NewStubServer(t, "/bbsmenu.html", boardListHtml)
	defer s.Close()

	r := infra.NewBoardListRepository(s.URL)

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

	r := infra.NewBoardListRepository(s.URL)

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

	r := infra.NewThreadListRepository(func(sId domain.ServerId, bId domain.BoardId) string {
		if sId != serverId || bId != boardId {
			t.Fatalf("invalid args of urlBuilderFunc: serverId=%s, boardId=%s", sId, bId)
		}

		return fmt.Sprintf("%s/%s/subback.html", s.URL, boardId)
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
