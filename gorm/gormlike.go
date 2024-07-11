package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type Like struct {
// 	ID        int    `gorm:"primary_key"`
// 	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
// 	Ua        string `gorm:"type:varchar(256);not null;"`
// 	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
// 	Hash      uint64 `gorm:"unique_index:hash_idx;"`
// 	CreatedAt time.Time
// }

// func main() {
// 	db, err := gorm.Open("mysql", "root:@Xfq347@@(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	db.AutoMigrate(&Like{})

// 	if !db.HasTable(&Like{}) {
// 		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
// 			panic(err)
// 		}
// 	}

// 	like := &Like{
// 		Ip:        ip,
// 		Ua:        ua,
// 		Title:     title,
// 		Hash:      murmur3.Sum64([]byte(strings.Join([]string{ip, ua, title}, "-"))) >> 1,
// 		CreatedAt: time.Now(),
// 	}

// 	if err := db.Create(like).Error; err != nil {
// 		return err
// 	}

// 	if err := db.Where(&Like{Hash: hash}).Delete(Like{}).Error; err != nil {
// 		return err
// 	}

// 	var count int
// 	err := db.Model(&Like{}).Where(&Like{Ip: ip, Ua: ua, Title: title}).Count(&count).Error
// 	if err != nil {
// 		return false, err
// 	}

// 	db.Model(&Like).Update("name", "hello")
// 	db.Model(&Like).Updates(User{Name: "hello", Age: 18})
// 	db.Model(&Like).Updates(User{Name: "", Age: 0, Actived: false})
// }

func CreateAnimals(db *gorm.DB) err {
	tx := db.Begin()
	if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Craete(&Animal{Name: "Lion"}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
