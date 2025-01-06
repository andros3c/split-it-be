package postgres_driver

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Host string
	DB_User string
	DB_Password string
	DB_Name string
	DB_Port int
	DB_SSLMode string
	DB_TimeZone string
}


func (conf *ConfigDB)InitialDB()*gorm.DB{
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		conf.DB_Host,
		conf.DB_User,
		conf.DB_Password,
		conf.DB_Name,
		conf.DB_Port,
		conf.DB_SSLMode,
		conf.DB_TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	
	return db
}

	
	
		
			
			