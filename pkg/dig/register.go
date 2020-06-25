package dig

import (
	"go.uber.org/dig"

	"github.com/iyu/go-example/pkg/dig/client"
	"github.com/iyu/go-example/pkg/dig/fuga"
	"github.com/iyu/go-example/pkg/dig/hoge"
)

func Register() *dig.Container {
	container := dig.New()
	err := container.Provide(func () hoge.HogeClient {
		return client.New("hoge")
	})
	if err != nil {
		panic(err)
	}
	err = container.Provide(func () fuga.FugaClient {
		return client.New("fuga")
	})
	if err != nil {
		panic(err)
	}
	err = container.Provide(hoge.New)
	if err != nil {
		panic(err)
	}
	err = container.Provide(fuga.New)
	if err != nil {
		panic(err)
	}

	return container
}
