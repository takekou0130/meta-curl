package gateway

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/takekou0130/meta-curl/application/repository"
	"github.com/takekou0130/meta-curl/domain"
)

type Gateway struct {
	client *http.Client
}

func NewGateway(c *http.Client) repository.Repository {
	return &Gateway{
		client: c,
	}
}

func (gw *Gateway) Fetch(url domain.Url) (*domain.Document, error) {
	req, err := http.NewRequest(http.MethodGet, url.Url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to http.NewRequest")
	}
	req.Header.Set("User-Agent", "")

	res, err := gw.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to client.Do")
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to goquery.NewDocumentFromReader")
	}
	return domain.NewDocument(&url, doc), nil
}
