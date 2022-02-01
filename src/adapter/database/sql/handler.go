package sql

import (
	"fmt"
	"gofiber-example/src/adapter/config"
	"gofiber-example/src/internal/model/dbmodel"
	"log"
	"strconv"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbMutex sync.Mutex

type GetGormDBFn func() (*gorm.DB, error)

func GetGormDB() (*gorm.DB, error) {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("port error")
	}
	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	dbMutex.Lock()
	if DB == nil {
		fmt.Println("Creating DB Instance...")

		DB, err = gorm.Open(postgres.Open(dsn))
		if err == nil {
			sqlDb, _ := DB.DB()
			fmt.Println("Pinging DB Instance...")
			err = sqlDb.Ping()
			if err != nil {
				fmt.Println("Error Pinging DB Instance:", err.Error())
				return nil, err
			}
		}
	}
	dbMutex.Unlock()

	return DB, err
}

func DoMigration() error {
	DB, err := GetGormDB()

	if err != nil {
		panic(err.Error())
	}

	err = DB.AutoMigrate(&dbmodel.User{})

	if err != nil {
		panic(err.Error())
	}

	return nil
}
