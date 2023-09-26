package model

import (
	"encoding/json"
	"io"
)

type TaskReq struct {
	Header      string     `json:"header,omitempty"`
	Description string     `json:"description,omitempty"`
	Date        CustomTime `json:"date,omitempty"`
	Status      string     `json:"status,omitempty"`
}

func (tq *TaskReq) DecodeJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&tq)

	return err
}
