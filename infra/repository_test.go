package infra_test

import (
	"bbs_api/domain/boardlist"
	"bbs_api/infra"
	_ "embed"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed test_data/bbsmenu.html
var html string

func TestBoardRepository_GetBoardGroups(t *testing.T) {
	s := NewStubServer(t, html)
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
	s := NewStubServer(t, html)
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

func NewStubServer(t *testing.T, response string) *httptest.Server {
	t.Helper()

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantPath := "/bbsmenu.html"

		if path := r.URL.Path; path != wantPath {
			t.Fatalf("path must be %s, but %s", wantPath, path)
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, response)
	}))

	return s
}
