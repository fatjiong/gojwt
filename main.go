package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/fatjiong/gojwt/controllers"
	"github.com/fatjiong/gojwt/model"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	//初始化数据库
	db, err := model.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	router := gin.Default()
	router.POST("/account/register", controllers.RegisterPost)
	router.POST("/account/login", controllers.LoginPost)

	//添加群组中间件
	authorized := router.Group("/user", TokenMiddelware())

	authorized.GET("/info", controllers.UserInfoGet)

	router.Run(":8080")
}

//token验证中间件
func TokenMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte("test"), nil
		})

		if err == nil {
			if token.Valid {
				c.Next()
			} else {
				c.JSON(http.StatusForbidden, gin.H{
					"code":    403,
					"message": "没有权限",
				})
				return
			}
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": err.Error(),
			})
		}
	}
}
