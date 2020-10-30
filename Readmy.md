import "github.com/xxarchexx/tnvcache"

###Usage 
//set con string
var connectionString = "mongodb://root:example@127.0.0.1:28019"


func main() {
    //declare API
	api := tnvcache.API{}
    //initaialize internal stuff
	api.Init()
    //set mongo connection
	api.ConnectRepository.NewConnection(dbanme, table, connectionString)
	//also could SetConnection
    //api.ConnectRepository.SetConnection(dbanme, table, mogno.Client)
    
    //useage
    var result = api.Request("яблоко")
    //close connection
	defer api.ConnectRepository.CloseConnection()
}
