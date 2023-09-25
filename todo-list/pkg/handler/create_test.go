package handler

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	client := resty.New()

	resp, _ := client.R().Post("http://localhost:8080/api/v1/tasks")

	assert.Equal(t, 200, resp.StatusCode())
}
