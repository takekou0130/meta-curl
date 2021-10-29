package usecase

import (
	"fmt"
	"log"

	"github.com/takekou0130/meta-curl/application/inputPort"
	"github.com/takekou0130/meta-curl/application/repository"
	"github.com/takekou0130/meta-curl/domain"
)

type MetaInfoUsecase struct {
	repository repository.Repository
}

func NewMetaInfoUsecase(rp *repository.Repository) inputPort.InputPort {
	return &MetaInfoUsecase{
		repository: *rp,
	}
}

func (m *MetaInfoUsecase) MetaInfo(args []string) domain.MetaInfo {
	var urls []*domain.Url
	for _, arg := range args {
		url, err := domain.NewUrl(arg)
		if err != nil {
			log.Fatalf("%v is not url", arg)
		} else {
			urls = append(urls, url)
		}
	}

	// TODO: 複数の情報を取れるようにする
	doc, err := m.repository.Fetch(*urls[0])
	if err != nil {
		fmt.Errorf("%v", doc)
	}

	// TODO errorハンドリング
	info := doc2metaInfo(doc)
	return info
}

func doc2metaInfo(doc domain.Document) domain.MetaInfo {
	return domain.MetaInfo{
		Url:         doc.GetUrl(),
		Title:       doc.GetTitle(),
		Description: doc.GetDesc(),
		Keywords:    doc.GetKeywords(),
		Canonical:   doc.GetCanonicals(),
		Alternate:   doc.GetAlternates(),
	}
}
