package tnvcache

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)



//Request or get from mysql
func (api *API) Request(key string) (string, error) {
	c := api.cache
	fmt.Printf("%v", c)

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
		api.cache.Set(key, strData)

	} else {
		strData = result
	}

	return strData, nil
}
