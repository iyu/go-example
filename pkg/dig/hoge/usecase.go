package hoge

import (
	"github.com/iyu/go-example/pkg/dig/client"
)

type HogeClient client.Client

type Hoge interface {
	Name() string
}

type hoge struct {
	client HogeClient
}

func New(c HogeClient) Hoge {
	return &hoge{c}
}

func (h *hoge) Name() string {
	return h.client.Name()
}
