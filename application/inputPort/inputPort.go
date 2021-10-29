package inputPort

import (
	"github.com/takekou0130/meta-curl/domain"
)

type InputPort interface {
	MetaInfo([]string) domain.MetaInfo
}
