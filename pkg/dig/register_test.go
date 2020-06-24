package dig

import (
	"testing"

	"github.com/iyu/go-example/pkg/dig/fuga"
	"github.com/iyu/go-example/pkg/dig/hoge"
)

func TestRegister(t *testing.T) {
	container := Register()
	container.Invoke(func (hoge hoge.Hoge, fuga fuga.Fuga) {
		if hoge.Name() != "hoge" {
			t.Fail()
		}
		if fuga.Name() != "fuga" {
			t.Fail()
		}
	})
}
