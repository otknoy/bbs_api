package response_test

import (
	"bbs_api/domain/boardlist"
	"bbs_api/domain/thread"
	"bbs_api/domain/threadlist"
	"bbs_api/interfaces/response"
	"bbs_api/openapi"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestCreateBoardListResponse(t *testing.T) {
	in := boardlist.BoardGroups{
		boardlist.BoardGroup{
			Name: "foo",
			BoardList: []boardlist.Board{
				{Name: "hoge", ServerId: "abc", BoardId: "def"},
				{Name: "huga", ServerId: "ghi", BoardId: "jkl"},
			},
		},
		boardlist.BoardGroup{
			Name:      "bar",
			BoardList: []boardlist.Board{},
		},
	}

	want := &openapi.BoardListResponse{
		BoardGroups: []openapi.BoardGroup{
			{
				Name: "foo",
				BoardList: []openapi.Board{
					{Name: "hoge", ServerId: "abc", BoardId: "def"},
					{Name: "huga", ServerId: "ghi", BoardId: "jkl"},
				},
			},
			{
				Name:      "bar",
				BoardList: []openapi.Board{},
			},
		},
	}

	got := response.CreateBoardListResponse(in)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("differs\n%v\n", diff)
	}
}

func TestCreateThreadListResponse(t *testing.T) {
	in := threadlist.ThreadList{
		{ThreadId: "foo", Name: "fooName"},
		{ThreadId: "bar", Name: "barName"},
	}

	want := &openapi.ThreadListResponse{
		ThreadList: []openapi.ThreadListResponseThreadList{
			{Id: "foo", Name: "fooName"},
			{Id: "bar", Name: "barName"},
		},
	}

	got := response.CreateThreadListResponse(in)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("differs\n%v\n", diff)
	}
}

func TestCreateThreadResponse(t *testing.T) {
	in := &thread.Thread{
		CommentList: []thread.Comment{
			{
				Meta: thread.Meta{
					Number:   1,
					UserName: "Kobe Bryant",
					UserId:   "black_mamba",
					PostedAt: time.Date(2020, 1, 2, 3, 4, 5, 6, time.UTC),
				},
				Text: "foobar",
			},
			{
				Meta: thread.Meta{
					Number:   2,
					UserName: "Lebron James",
					UserId:   "king_james",
					PostedAt: time.Date(2020, 1, 2, 3, 4, 5, 7, time.UTC),
				},
				Text: "hoge",
			},
		},
	}

	want := &openapi.ThreadResponse{
		CommentList: []openapi.Comment{
			{
				Meta: openapi.CommentMeta{
					Number:   1,
					UserName: "Kobe Bryant",
					UserId:   "black_mamba",
					PostedAt: time.Date(2020, 1, 2, 3, 4, 5, 6, time.UTC),
				},
				Text: "foobar",
			},
			{
				Meta: openapi.CommentMeta{
					Number:   2,
					UserName: "Lebron James",
					UserId:   "king_james",
					PostedAt: time.Date(2020, 1, 2, 3, 4, 5, 7, time.UTC),
				},
				Text: "hoge",
			},
		},
	}

	got := response.CreateThreadResponse(in)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("differs\n%v\n", diff)
	}

}
