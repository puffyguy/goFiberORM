package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
# GetMongoDB - Setting up MongoDB Connection
*/
func GetMongoDB() *mongo.Database {
	credential := options.Credential{
		AuthSource: envVariable("MONGO_AUTHDB"), //the name of the database to use for authentication
		Username:   envVariable("MONGO_USERNAME"),
		Password:   envVariable("MONGO_PASSWORD"),
	}
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(envVariable("MONGO_DBNAME"))
}

/*
# - GetMySQLDB - Setting up MySQL Connection

- USER_NAME string
- PASSWORD string
- DBNAME string
- DBURL string
*/
func GetMySQLDB() *gorm.DB {
	uname, pass, dbname, dburl := envVariable("MYSQL_USERNAME"), envVariable("MYSQL_PASSWORD"), envVariable("MYSQL_DBNAME"), envVariable("MYSQL_DBURL")
	dsn := uname + ":" + pass + "@" + dburl + "/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return db
}

/*
# envVariable - Reads .env file and returns value of the environment variable named by the key
*/
func envVariable(key string) string {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error loading .env file")
		return ""
	}

	return os.Getenv(key)
}
