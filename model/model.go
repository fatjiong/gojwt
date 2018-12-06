package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

type Category struct {
	gorm.Model
	Pid  uint
	Name string
	Sort uint
}

type Adver struct {
	gorm.Model
	Title     string
	Thumb     string
	Sort      uint
	LinkUrl   string
	ArticleId uint
}
type Article struct {
	gorm.Model
	CategoryId  uint
	Title       string
	Thumb       string
	Tags        string
	Author      string
	Description string
	Status      uint
	IsRecommend uint
	Hits        uint
	Body        string
	Up          uint
	Down        uint
}

type Comment struct {
	gorm.Model
	ArticleId uint
	CommentId uint
	IsCheck   uint
	Title     string
	Body      string
}

type Feedback struct {
	gorm.Model
	Name   string
	Email  string
	Title  string
	Body   string
	Remark string
}

type Sharelink struct {
	gorm.Model
	Name    string
	Thumb   string
	IsCheck uint
	LinkUrl string
}

type User struct {
	gorm.Model
	Account   string
	Password  string
	Salt      string
	Gender    uint
	TrueName  string
	NickName  string
	Status    uint
	LoginTime time.Time
}

/**
初始化数据库
*/
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:123456@/gblog?charset=utf8mb4&parseTime=True&loc=Local")
	if err == nil {
		DB = db
		//设置默认表名前缀
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return "gb_" + defaultTableName
		}
		db.AutoMigrate(&Category{}, &Article{}, &Adver{}, &Comment{}, &Feedback{}, &Sharelink{}, &User{})
		return db, nil
	}
	return nil, err
}

//用户注册
func UserRegister(account string, password string) (*User, error) {
	user := &User{
		Account:  account,
		Password: password,
	}
	if err := DB.Create(user).Error; err != nil {
		return nil, err
	} else {
		return user, err
	}
}

//用户查询
func Userdetail(account string, password string) (*User, error) {
	var user User

	if err := DB.First(&user, "account = ? and password = ?", account, password).Error; err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

//创建token
func CreateToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	// 根据信息生成token
	if tokenString, err := token.SignedString([]byte("test")); err != nil {
		return ""
	} else {
		return tokenString
	}
}
