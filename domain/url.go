package domain

import (
	"strings"

	"github.com/pkg/errors"
)

type Url struct {
	Url string
}

var ErrUrl = errors.New("unexpected url")

func NewUrl(url string) (*Url, error) {
	if len(url) <= 0 {
		return nil, errors.Wrap(ErrUrl, "url is empty")
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return nil, errors.Wrap(ErrUrl, "invalid url")
	}

	return &Url{url}, nil
}
