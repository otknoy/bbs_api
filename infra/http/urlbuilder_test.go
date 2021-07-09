package http_test

import (
	"bbs_api/infra/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUrlBuilder(t *testing.T) {
	b := http.NewUrlBuilder("example.com")

	t.Run("GetBuildBoardUrl", func(t *testing.T) {
		want := "http://menu.example.com/bbsmenu.html"

		got := b.BuildGetBoardUrl()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("differ\n%v\n", diff)
		}
	})

	t.Run("GetThreadListUrl", func(t *testing.T) {
		want := "http://test-serverId.example.com/test-boardId/subback.html"

		got := b.BuildGetThradListUrl("test-serverId", "test-boardId")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("differ\n%v\n", diff)
		}
	})

	t.Run("GetThreadUrl", func(t *testing.T) {
		want := "http://test-serverId.example.com/test/read.cgi/test-boardId/test-threadId"

		got := b.BuildGetThreadUrl("test-serverId", "test-boardId", "test-threadId")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("differ\n%v\n", diff)
		}
	})
}
