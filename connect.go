package tnvcache

import (
	"github.com/xxarchexx/tnvcache/internal"
	"github.com/xxarchexx/tnvcache/internal/mongo"
)

type API struct {
	ConnectRepository internal.ConnectionRepository
	storeRepository   *mongo.Store
	cache             internal.CacheRepository
}

func (apiStore *API) Init() {
	apiStore.storeRepository = mongo.Init()
	apiStore.ConnectRepository = apiStore.storeRepository.GetConnectionRepository()
	var setInterface internal.CacheRepository
	setInterface = internal.CacheRepository(apiStore.storeRepository)
	apiStore.cache = setInterface
}
