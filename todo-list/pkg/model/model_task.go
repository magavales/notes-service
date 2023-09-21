package model

import (
	"encoding/json"
	"io"
)

type Task struct {
	ID          int64      `json:"task_id"`
	Header      string     `json:"header"`
	Description string     `json:"description"`
	Date        CustomTime `json:"date"`
	Status      string     `json:"status"`
}

func (t *Task) DecodeJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&t)

	return err
}
