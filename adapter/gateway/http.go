package gateway

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/takekou0130/meta-curl/domain"
)

type Gateway struct {
	client *http.Client
}

func NewGateway(c *http.Client) *Gateway {
	return &Gateway{
		client: c,
	}
}

func (gw *Gateway) Fetch(url domain.Url) (domain.Document, error) {
	req, err := http.NewRequest(http.MethodGet, url.Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "")

	res, err := gw.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	return *domain.NewDocument(&url, doc), err
}
