package bbsclient

import (
	"bufio"
	"context"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func NewShiftJISDocument(ctx context.Context, url string) (*goquery.Document, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header.Add("User-Agent", "curl/7.74.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	log.Printf("http get: url=%s, status=%s\n", url, res.Status)

	reader := transform.NewReader(
		bufio.NewReader(res.Body),
		japanese.ShiftJIS.NewDecoder(),
	)

	return goquery.NewDocumentFromReader(reader)
}
