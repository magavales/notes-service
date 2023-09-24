package model

import (
	"errors"
	"net/url"
)

type Status struct {
	Status string
}

func (s *Status) ParseQueryParams(url *url.URL) (err error) {
	if url.Query().Has("status") {
		s.Status = url.Query().Get("status")
		return nil
	} else {
		err = errors.New("don't have parameter 'status'")
		return err
	}
}
