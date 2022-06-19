package configs

import (
	"api/blueprints/users/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var TX *gorm.DB
var HTTP_PORT int

func InitDb() {
	var err error
	// connectionString := os.Getenv("CONNECTION_STRING")
	connectionString := "root:toor@tcp(localhost:3306)/simple_ecommerce?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	var err error

	// HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	HTTP_PORT = 8080
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.Users{})
}

func DBBegin() {
	TX = DB.Begin()
}

func DBCommit() {
	TX.Commit()

	DBClose()
}

func DBRollback() {
	TX.Rollback()

	DBClose()
}

func DBClose() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}

func ConfigTest() (*gorm.DB, error) {
	var err error
	connectionStringTest := "root:toor@tcp(localhost:3306)/simple_ecommerce_test?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(connectionStringTest), &gorm.Config{})
	if err != nil {
		return DB, err
	}
	return DB, err
}
