package db

import (
	"fmt"
	"github.com/rasyidmm/RumahMakan/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbd *gorm.DB
var err error

func Init() *gorm.DB{
	conf := config.GetConfig()
	//connectionString := "postgres://"+conf.DB_USERNAME+":"+conf.DB_PASSWORD+"@"+conf.DB_HOST+":"+conf.DB_PORT+"/"+conf.DB_NAME+"?sslmode=disable"
	connectionString := "host="+conf.DB_HOST+" user="+conf.DB_USERNAME+" password="+conf.DB_PASSWORD+" dbname="+conf.DB_NAME+" port="+conf.DB_PORT+" sslmode=disable"
	dbd, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	fmt.Println(connectionString)
	if err != nil{
		panic("Db Error")
	}
	fmt.Println("connection DB sukses")
	return dbd
}

