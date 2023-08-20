package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Sign(db *gorm.DB, r *gin.Engine) {

	r.POST("/register", func(ctx *gin.Context) {

		//获取参数
		name := ctx.PostForm("name")
		password := ctx.PostForm("password")

		//数据验证
		if len(name) == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "用户名不能为空",
			})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "密码不能少于6位",
			})
			return
		}

		//判断用户名是否存在
		var user User
		db.Where(" name= ?", name).First(&user)
		if user.ID != 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "用户已存在",
			})
			return
		}

		//创建用户
		hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    500,
				"message": "密码加密错误",
			})
			return
		}
		newUser := User{
			Name:     name,
			Password: string(hasedPassword),
		}
		db.Create(&newUser)

		//返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "注册成功",
		})

		return
	})
	return
}

// 登录
func Log(db *gorm.DB, r *gin.Engine) {
	r.POST("/login", func(ctx *gin.Context) {

		//获取参数
		id := ctx.PostForm("id")
		password := ctx.PostForm("password")

		//数据验证
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "密码不能少于6位",
			})
			return
		}

		//判断id是否存在
		var user User
		db.Where("id = ?", id).First(&user)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "用户不存在",
			})
			return
		}

		//判断密码是否正确
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "密码错误",
			})
		}

		//返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "登录成功",
		})
	})
}
