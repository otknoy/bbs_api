package infra

import (
	"bufio"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func NewShiftJISDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	reader := transform.NewReader(
		bufio.NewReader(res.Body),
		japanese.ShiftJIS.NewDecoder(),
	)

	return goquery.NewDocumentFromReader(reader)
}
