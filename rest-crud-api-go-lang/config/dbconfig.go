package config

import (
	"fmt"
	"log"
	"rest-crud-api-go-lang/entity"

	"github.com/jinzhu/gorm"
)

func InitDB() {
	config :=
		Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "userdb",
		}

	connectionString := GetConnectionString(config)
	err := Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	//Migrate database table
	Connector.AutoMigrate(&entity.User{})
	log.Println("Table migrated")
}

type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.
		Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&"+
			"collation=utf8mb4_unicode_ci&"+
			"parseTime=true&multiStatements=true",
			config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}
var Connector *gorm.DB

//MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}
