package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:@Xfq347@@(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{Name: "枯藤", Gender: "男", Hobby: "篮球"}
	u2 := UserInfo{Name: "topgoer.com", Gender: "女", Hobby: "足球"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)

	// 查询
	var u UserInfo
	db.First(&u) // 查找第一条记录（按主键排序）
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby = ?", "足球") // 查找 hobby 为 足球 的记录
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("Hobby", "双色球")

	// 删除
	db.Delete(&u)
}
