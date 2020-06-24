package client

type Client interface {
	Name() string
}

type client struct {
	name string
}

func New(name string) Client {
	return &client{name}
}

func (d *client) Name() string {
	return d.name
}
