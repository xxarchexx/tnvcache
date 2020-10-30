package internal

//ConnectionRepository for new connection, close connection and set exixts connection
type ConnectionRepository interface {
	NewConnection(dbname, tableName, uri string) error
	SetConnection(dbname, tableName string, client interface{})
	CloseConnection() error
}

//CacheRepository for new connection, close connection and set exixts connection
type CacheRepository interface {
	Get(key string) (string, error)
	Set(key, value string) error
}
