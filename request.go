package tnvcache

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/xxarchexx/tnvcache/internal"
)

type api struct {
	cache internal.CacheRepository
}

//Request or get from mysql
func (api *api) Request(key string) (string, error) {
	result, err := api.cache.Get(key)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var strData string = ""

	if result == "" {

		v := url.Values{}
		v.Set("Query", key)

		resp, err := http.Get(fmt.Sprintf("https://api.tnved.s-3.dev/api/Search/Search?%v", v.Encode()))

		if err != nil {
			log.Fatal(err)
			return "", err
		}

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		strData = string(responseData)

	} else {
		strData = result
	}

	return strData, nil
}
