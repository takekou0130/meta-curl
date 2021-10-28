package repository

import "github.com/takekou0130/meta-curl/domain"

type Repository interface {
	Fetch(domain.Url) (domain.Document, error)
}
