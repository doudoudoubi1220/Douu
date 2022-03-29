package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	//err=
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
