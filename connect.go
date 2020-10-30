package tnvcache

import (
	"github.com/xxarchexx/tnvcache/internal"
	"github.com/xxarchexx/tnvcache/internal/mongo"
)

type API struct {
	cache  internal.CacheRepository
	iclose internal.IClose
}

func (api *API) WithMongo(dbname, tablename, connectionstring string) {
	api.cache, api.Close = mongo.Init(dbname, tablename, connectionstring)
}

func (api *API) Close() {
	api.iclose.Close()
}
