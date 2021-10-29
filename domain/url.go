package domain

import (
	"fmt"
	"strings"
)

type Url struct {
	Url string
}

func NewUrl(url string) (*Url, error) {
	if len(url) <= 0 {
		return nil, fmt.Errorf("%v is empty", url)
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return nil, fmt.Errorf("%v is not url", url)
	}

	return &Url{url}, nil
}
