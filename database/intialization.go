package database

var connection string

func init() {
	connection = "mySQL"
}

func GetDatabase() string {
	return connection
}
