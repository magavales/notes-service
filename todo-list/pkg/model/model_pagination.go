package model

import (
	"net/url"
	"strconv"
)

type Pagination struct {
	Limit  int
	Offset int
}

func (p *Pagination) ParseQueryParams(url *url.URL) {
	if url.Query().Has("limit") {
		p.Limit, _ = strconv.Atoi(url.Query().Get("limit"))
	} else {
		p.Limit = 10
	}
	if url.Query().Has("offset") {
		p.Offset, _ = strconv.Atoi(url.Query().Get("offset"))
	} else {
		p.Offset = 0
	}
}
