package domain

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

type Document struct {
	url *Url
	doc *goquery.Document
}

func NewDocument(url *Url, doc *goquery.Document) *Document {
	return &Document{url, doc}
}

var ErrDontExistAttr = errors.New("Attributes don't exist")

func (doc *Document) GetUrl() Url {
	return *doc.url
}

func (doc *Document) GetTitle() []string {
	title := doc.doc.Find("title").Text()
	return []string{title}
}

func (doc *Document) GetDesc() ([]string, error) {
	desc, ok := doc.doc.Find("meta[name='description']").Attr("content")
	if !ok {
		return nil, errors.Wrap(ErrDontExistAttr, "description don't exist")
	}
	return []string{desc}, nil
}
func (doc *Document) GetKeywords() ([]string, error) {
	key, ok := doc.doc.Find("meta[name='keywords']").Attr("content")
	if !ok {
		return nil, errors.Wrap(ErrDontExistAttr, "keywords don't exist")
	}
	return []string{key}, nil
}
func (doc *Document) GetCanonicals() ([]string, error) {
	cano, ok := doc.doc.Find("link[rel='canonical']").Attr("href")
	if !ok {
		return nil, errors.Wrap(ErrDontExistAttr, "canonicals don't exist")
	}
	return []string{cano}, nil
}
func (doc *Document) GetAlternates() ([]string, error) {
	alt, ok := doc.doc.Find("link[rel='alternate']").Attr("href")
	if !ok {
		return nil, errors.Wrap(ErrDontExistAttr, "alternates don't exist")
	}
	return []string{alt}, nil
}
