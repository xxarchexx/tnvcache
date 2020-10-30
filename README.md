    import "github.com/xxarchexx/tnvcache"

   

    //set con string

    var connectionString = "mongodb://root:example@127.0.0.1:28019"


    func main() {
        
        //declare API
        api := tnvcache.API{}
        
        //initialize internal stuff
        api.Init()
       
        //create new  mongo connection
        api.ConnectRepository.NewConnection(dbanme, table, connectionString)
        
        //also could SetConnection (exists)
        //api.ConnectRepository.SetConnection(dbanme, table, mogno.Client)
        
        //usage
        var result = api.Request("яблоко")
        
        //close connection
        defer api.ConnectRepository.CloseConnection()
    }
