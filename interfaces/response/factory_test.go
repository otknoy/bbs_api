package response_test

import (
	"bbs_api/domain/boardlist"
	"bbs_api/interfaces/response"
	"bbs_api/openapi"
	"testing"

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
