package bbsclient_test

import (
	"bbs_api/infra/bbsclient"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUrlBuilder(t *testing.T) {
	b := bbsclient.NewUrlBuilder("example.com")

	t.Run("BuildBoardListUrl", func(t *testing.T) {
		want := "http://menu.example.com/bbsmenu.html"

		got := b.BuildBoardListUrl()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("differ\n%v\n", diff)
		}
	})

	t.Run("BuildThreadListUrl", func(t *testing.T) {
		want := "http://test-serverId.example.com/test-boardId/subback.html"

		got := b.BuildThreadListUrl("test-serverId", "test-boardId")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("differ\n%v\n", diff)
		}
	})

	t.Run("BuildThreadUrl", func(t *testing.T) {
		want := "http://test-serverId.example.com/test/read.cgi/test-boardId/test-threadId"

		got := b.BuildThreadUrl("test-serverId", "test-boardId", "test-threadId")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("differ\n%v\n", diff)
		}
	})
}
