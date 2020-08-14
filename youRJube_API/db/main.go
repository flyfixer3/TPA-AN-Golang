package database
import (
	"fmt"
	"github.com/jinzhu/gorm"
	
)

type dbConfig struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
}

var config = dbConfig{"localhost", 5432, "postgres", "youRJube", "alvin123"}
const dbConnection = "postgres://postgres:alvin123@localhost:5432/youRJube?sslmode=disable"
//getDatabaseUrl -> get database url
func getDatabaseURL() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		config.host, config.port, config.user, config.dbname, config.password)
}
//GetDatabase -> return database connection
func GetDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("postgres",dbConnection)
	return db, err
}

