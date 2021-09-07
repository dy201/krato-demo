package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// 用来生成数据库

type User struct {
	gorm.Model
	UserName string `gorm:"column:user_name;index:idx_username;type:varchar(255); unique;not null; comment('用户名') "`
	Password string `gorm:"column:password;type:varchar(255);not null;comment '加密用&分开' "`
	Nickname string `gorm:"column:nickname;index:idx_nickname; type:varchar(255);comment('昵称')  "`
	Status   int32  `gorm:"column:status; type:int(3); default:0; comment '状态，1-启用，0-弃用'"`
	Email    string `gorm:"column:email;comment('邮箱') index type:varchar(255)"`
	Phone    string `gorm:"column:phone;index:idx_phone; type:varchar(255); comment('手机号')  "`
	Age      uint32 `gorm:"column:age; type:int(3); comment('年龄') "`
	RoleID   int32  `gorm:"column:role_id; type:int(10); comment('角色id') "`
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/doyle_blog?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	//options := &password.Options{16, 100, 32, sha512.New}
	//salt, encodedPwd := password.Encode("admin123", options)
	//newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s",salt,encodedPwd)

	db.AutoMigrate(&User{})
}
