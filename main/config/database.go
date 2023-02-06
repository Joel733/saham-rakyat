package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
  func InitDatabase() (*gorm.DB, error){

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	os.Getenv("username"),
	os.Getenv("password"),
	os.Getenv("db_host"),
	os.Getenv("db_port"),
	os.Getenv("db_schema"),
	)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	return db, err
}