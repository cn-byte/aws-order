package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//TestDB
var TestDB *gorm.DB

// InitMysqlConfig 初始化mysql
func InitMysqlConfig() {
	var err error
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	testDb, _ := sql.Open("mysql", dsn)
	testDb.SetMaxOpenConns(20)        //  最大空闲连接数
	testDb.SetMaxIdleConns(100)       // 最大在线连接数
	testDb.SetConnMaxLifetime(180000) //连接最大生命周期(单位：毫秒)
	TestDB, err = gorm.Open(mysql.New(mysql.Config{Conn: testDb}), &gorm.Config{})
	if err != nil {
		panic("connect mysql error, error=" + err.Error())
	}
}
