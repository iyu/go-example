package fuga

import (
	"github.com/iyu/go-example/pkg/dig/client"
)

type FugaClient = client.Client

type Fuga interface {
	Name() string
}

type fuga struct {
	client FugaClient
}

func New(c FugaClient) Fuga {
	return &fuga{c}
}

func (f *fuga) Name() string {
	return f.client.Name()
}
