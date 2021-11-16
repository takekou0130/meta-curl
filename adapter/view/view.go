package view

import "github.com/takekou0130/meta-curl/domain"

type View interface {
	Render(*domain.MetaInfo) error
}
