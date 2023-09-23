package model

import (
	"encoding/json"
	"io"
	"time"
)

type Task struct {
	ID          int64      `json:"task_id,omitempty"`
	Header      string     `json:"header,omitempty"`
	Description string     `json:"description,omitempty"`
	Date        CustomTime `json:"date,omitempty"`
	Status      string     `json:"status,omitempty"`
}

func (t *Task) DecodeJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&t)

	return err
}

func (t *Task) ParseRowsFromTable(values []interface{}) {
	t.ID = values[0].(int64)
	t.Header = values[1].(string)
	t.Description = values[2].(string)
	t.Date.Time = values[3].(time.Time)
	t.Status = values[4].(string)
}
