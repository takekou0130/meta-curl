package domain

import (
	"github.com/PuerkitoBio/goquery"
)

type Document struct {
	url *Url
	doc *goquery.Document
}

func NewDocument(url *Url, doc *goquery.Document) *Document {
	return &Document{url, doc}
}

func (doc *Document) GetUrl() Url {
	return *doc.url
}

func (doc *Document) GetTitle() []string {
	title := doc.doc.Find("title").Text()
	return []string{title}
}

func (doc *Document) GetDesc() []string {
	desc, _ := doc.doc.Find("meta[name='description']").Attr("content")
	return []string{desc}
}
func (doc *Document) GetKeywords() []string {
	key, _ := doc.doc.Find("meta[name='keywords']").Attr("content")
	return []string{key}
}
func (doc *Document) GetCanonicals() []string {
	cano, _ := doc.doc.Find("link[rel='canonical']").Attr("href")
	return []string{cano}
}
func (doc *Document) GetAlternates() []string {
	alt, _ := doc.doc.Find("link[rel='alternate']").Attr("href")
	return []string{alt}
}
