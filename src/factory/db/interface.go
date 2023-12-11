package db

import (
	"tracker_backend/src/adapter"
	"tracker_backend/src/factory"
)

type AbsDbGatewayFactory interface {
	Build(factory.CtxDeps) (adapter.DbGateway, error)
}
