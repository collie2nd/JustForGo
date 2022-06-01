package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	var (
		db  *gorm.DB
		err error
	)
	debug := flag.Bool("d", false, "Pg调试开关")
	flag.Parse()
	dsn := "host=%s port=%d user=%s dbname=%s password=%s sslmode=disable"
	dsn = fmt.Sprintf(dsn, "localhost", 5432, "postgres", "postgres", 123456)
	fmt.Println("Connect Database Str: ", dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,         // 禁用彩色打印
		},
	)
	for {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			fmt.Println("Connect to PgSQL failed.", err)
			time.Sleep(30 * time.Second)
		} else {
			break
		}
	}

	if *debug {
		db = db.Debug()
	}

	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "枯藤", "男", "篮球"}
	u2 := UserInfo{2, "topgoer.com", "女", "足球"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)

	db.Table("user_info").CreateInBatches([]UserInfo{
		{3, "neo", "男", "xx"},
		{4, "bob", "女", "xx"},
	}, 10000)

	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)
	var uu []UserInfo
	// order 要在 find 前
	//db.Where(query, args...)
	db.Order("id").Find(&uu, "hobby = ?", "xx")
	fmt.Printf("%#v\n", uu)
	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	// ids 是id主键的列表
	ids := []int{1, 2}
	db.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&u, ids)
	db.Unscoped().Delete(&u, "name = ?", "枯藤")
}
