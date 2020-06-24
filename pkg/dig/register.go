package dig

import (
	"go.uber.org/dig"

	"github.com/iyu/go-example/pkg/dig/client"
	"github.com/iyu/go-example/pkg/dig/fuga"
	"github.com/iyu/go-example/pkg/dig/hoge"
)

func Register() *dig.Container {
	container := dig.New()
	container.Provide(func () hoge.HogeClient {
		return client.New("hoge")
	})
	container.Provide(func () fuga.FugaClient {
		return client.New("fuga")
	})
	container.Provide(hoge.New)
	container.Provide(fuga.New)

	return container
}
