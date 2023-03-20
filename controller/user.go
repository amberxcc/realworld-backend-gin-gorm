package controller

import (
	"github.com/awesomexu/go-realworld/models"
	"github.com/awesomexu/go-realworld/utils"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {

	// 1. validate request
	var registerBody struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBind(&registerBody); err != nil {
		ctx.JSON(422, gin.H{
			"message": "参数校验失败",
			"detail":  err.Error(),
		})
		return
	}

	// 2. validate email
	var user models.User
	utils.DB.Where("email = ?", registerBody.Email).First(&user)
	if user.ID != 0 {
		ctx.JSON(400, gin.H{
			"message": "email has been register!",
		})
		return
	}

	// 3. create user
	user.Email = registerBody.Email
	user.Username = registerBody.Username
	user.Password = registerBody.Password
	result := utils.DB.Create(&user)
	if result.Error != nil {
		panic("创建数据失败 =>" + result.Error.Error())
	}

	// 4. generate token
	token, err := utils.GenToken(user.ID)
	if err != nil {
		panic("生成token失败" + err.Error())
	}

	// 5. response
	ctx.JSON(200, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"bio":      user.Bio,
		"image":    user.Image,
		"token":    token,
	})
}

func Login(ctx *gin.Context) {

	// 1. validate request body
	var loginBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBind(&loginBody); err != nil {
		ctx.JSON(422, gin.H{
			"message": "参数校验失败",
			"detail":  err.Error(),
		})
		return
	}

	// 2. validate Email
	var user models.User
	utils.DB.Find(&user, "email=?", loginBody.Email)
	if user.ID == 0 {
		ctx.JSON(400, gin.H{
			"message": "Email不存在",
		})
		return
	}

	// 3. generate token
	token, err := utils.GenToken(user.ID)
	if err != nil {
		panic("生成token失败" + err.Error())
	}

	// 4. response
	ctx.JSON(200, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"bio":      user.Bio,
		"image":    user.Image,
		"token":    token,
	})
}

func GetCurrentUser(ctx *gin.Context) {
	userId, _ := ctx.Get("user")
	var user models.User
	utils.DB.First(&user, userId)

	ctx.JSON(200, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"bio":      user.Bio,
		"image":    user.Image,
	})
}

func UpdateUserInfo(ctx *gin.Context) {
	userId, _ := ctx.Get("user")
	var user models.User
	utils.DB.First(&user, userId)

	var updateUserBody struct {
		Username string `json:"username"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	}
	ctx.ShouldBind(&updateUserBody)

	if updateUserBody.Username != "" {
		user.Username = updateUserBody.Username
	}
	if updateUserBody.Bio != "" {
		user.Bio = updateUserBody.Bio
	}
	if updateUserBody.Image != "" {
		user.Image = updateUserBody.Image
	}
	utils.DB.Save(&user)

	ctx.JSON(200, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"bio":      user.Bio,
		"image":    user.Image,
	})
}
