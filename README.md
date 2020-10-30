package main

import (
	"github.com/xxarchexx/tnvcache"
)

var connectionString = "mongodb://root:example@127.0.0.1:28019"

type tt struct {
}

func main() {
		api := tnvcache.API{}
		api.WithMongo("tnved", "cache", connectionString)
		api.Request("яблоко")
		defer api.Close()
}
