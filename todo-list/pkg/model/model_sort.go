package model

import (
	"net/url"
)

type Sort struct {
	Param string
}

func (s *Sort) ParseQueryParams(url *url.URL) {
	s.Param = url.Query().Get("sort")
}
