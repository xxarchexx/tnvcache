package main

import (
	"fmt"

	"github.com/xxarchexx/tnvcache"
)

#####var connectionString = "mongodb://root:example@127.0.0.1:28019"




func main() {

	api := tnvcache.API{}
	//

    api.WithMongo("tnved", "cache", connectionString)
	
    //
    result, err := api.Request("яблоко")
	

    if err != nil {
		fmt.Println(err)
	}
     else {
		fmt.Println(result)
	}
    //
	defer api.Close()
}
