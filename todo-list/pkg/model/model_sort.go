package model

import (
	"errors"
	"net/url"
)

type Sort struct {
	Sort string
}

func (s *Sort) ParseQueryParams(url *url.URL) (err error) {
	if url.Query().Has("sort") {
		s.Sort = url.Query().Get("sort")
		return nil
	} else {
		err = errors.New("don't have parameter 'sort'")
		return err
	}
}
