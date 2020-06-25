package dig

import (
	"testing"

	"github.com/iyu/go-example/pkg/dig/fuga"
	"github.com/iyu/go-example/pkg/dig/hoge"
)

func TestRegister(t *testing.T) {
	container := Register()
	container.Invoke(func (hoge hoge.Hoge, fuga fuga.Fuga) {
		t.Log(hoge.Name())
		if hoge.Name() != "hoge" {
			t.Fail()
		}
		t.Log(fuga.Name())
		if fuga.Name() != "fuga" {
			t.Fail()
		}
	})
}
