package model

import (
	"net/url"
	"strconv"
)

type QueryParams struct {
	Status string
	Limit  int
	Offset int
}

func (qp *QueryParams) ParseQueryParams(url *url.URL) {
	qp.Status = url.Query().Get("status")
	qp.Limit, _ = strconv.Atoi(url.Query().Get("limit"))
	qp.Offset, _ = strconv.Atoi(url.Query().Get("offset"))
}
