package tnvcache

import (
	"github.com/xxarchexx/tnvcache/internal"
	"github.com/xxarchexx/tnvcache/internal/mongo"
)

type API struct {
	cache internal.CacheRepository
}

func (api *API) WithMongo(dbname, tablename, connectionstring string) {
	store := mongo.Store{}
	api.cache = store.Init(dbname, tablename, connectionstring)
}
