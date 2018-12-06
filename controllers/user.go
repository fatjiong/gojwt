package controllers

import (
	"fmt"
	"github.com/fatjiong/gojwt/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//账户注册处理
func RegisterPost(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")

	if user, err := model.UserRegister(account, password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "添加成功",
			"user":    user,
		})
	}
}

//登录方法
func LoginPost(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")

	if user, err := model.Userdetail(account, password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
			"user":    user,
			"token":   model.CreateToken(),
		})
	}
}

func UserInfoGet(c *gin.Context) {
	fmt.Println("123")
}
