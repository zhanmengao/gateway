package global

import "context"

type Node interface {
	GetHttpAddr() string
	GetWsAddr() string
}

type IRoute interface {
	Route(ctx context.Context, srv, key string) (n Node, err error)
}
