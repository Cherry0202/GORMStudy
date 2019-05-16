package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	type User struct {
		gorm.Model
		Name string `gorm:"size:255"`
		Password string `gorm:"size:255"`
		Email string `gorm:"size:255"`
	}

	//接続
	db, err := gorm.Open("mysql", "root@/note?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "root:password@/dbname?charset=utf8&parseTime=True&loc=Local")

	//table
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	// マイグレーション
	db.AutoMigrate(&User{})


	if err != nil {
		panic("failed to connect database")
	}

	if err == nil {
		print("Success!!")
	}

	// 実行完了後DB接続を閉じる
	defer db.Close()

	// ログ出力を有効にする
	db.LogMode(true)


}
