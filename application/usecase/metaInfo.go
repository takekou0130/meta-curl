package usecase

import (
	"github.com/pkg/errors"
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

func (m *MetaInfoUsecase) MetaInfo(args []string) (*domain.MetaInfo, error) {
	var urls []*domain.Url
	for _, arg := range args {
		url, err := domain.NewUrl(arg)
		if err != nil {
			return nil, errors.Wrap(err, "failed to domain.NewUrl")
		} else {
			urls = append(urls, url)
		}
	}

	// TODO: 複数の情報を取れるようにする
	doc, err := m.repository.Fetch(*urls[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to repository.Fetch")
	}

	// TODO errorハンドリング
	info, err := doc2metaInfo(doc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert doc to metaInfo")
	}
	return info, nil
}

func doc2metaInfo(doc *domain.Document) (*domain.MetaInfo, error) {
	Url := doc.GetUrl()
	Title := doc.GetTitle()
	Description, err := doc.GetDesc()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get info")
	}
	Keywords, err := doc.GetKeywords()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get info")
	}
	Canonical, err := doc.GetCanonicals()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get info")
	}
	Alternate, err := doc.GetAlternates()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get info")
	}
	return &domain.MetaInfo{
		Url,
		Title,
		Description,
		Keywords,
		Canonical,
		Alternate,
	}, nil
}
