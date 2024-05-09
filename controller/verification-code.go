package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"online-judge/pkg/resp"
	"online-judge/pkg/utils"
	"online-judge/services"
)

// SendEmailCode 发送邮箱验证码接口
// @Tags Verification Code API
// @Summary 发送邮箱验证码
// @Description 发送邮箱验证码接口
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "邮箱"
// @Success 200 {object} _Response "发送邮箱验证码成功"
// @Failure 200 {object} _Response "邮箱格式错误"
// @Failure 200 {object} _Response "服务器内部错误"
// @Router /send-email-code [POST]
func SendEmailCode(c *gin.Context) {
	userEmail := c.PostForm("email") //从前端获取email信息
	// 判断email是否合法

	if !utils.ValidateEmail(userEmail) {
		resp.ResponseError(c, resp.CodeInvalidateEmailFormat)
		zap.L().Error("controller-SendEmailCode-ValidateEmail " +
			fmt.Sprintf("invalid email %s ", userEmail))
		return
	}
	resCode := services.SendEmailCode(userEmail)
	switch resCode {
	// 成功
	case resp.Success:
		resp.ResponseSuccess(c, resp.CodeSuccess)

	// 邮箱格式错误
	case resp.InvalidateEmailFormat:
		resp.ResponseError(c, resp.CodeInvalidateEmailFormat)

	case resp.SearchDBError, resp.EncryptPwdError:
		resp.ResponseError(c, resp.CodeInternalServerError)

	default:
		resp.ResponseError(c, resp.CodeInternalServerError)
	}
}

// SendCode 发送图片验证码接口
// @Tags Verification Code API
// @Summary 发送图片验证码
// @Description 发送图片验证码接口
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "用户名"
// @Success 200 {object} _Response "发送图片验证码成功"
// @Failure 200 {object} _Response "服务器内部错误"
// @Router /send-code [POST]
func SendCode(c *gin.Context) {
	username := c.PostForm("username")
	b64s, err := services.SendCode(username)
	// 生成图片验证码失败
	if err != nil {
		resp.ResponseError(c, resp.CodeInternalServerError)
		return
	}
	resp.ResponseSuccess(c, b64s)

}

// CheckPictureCode 检查图片验证码接口
// @Tags Verification Code API
// @Summary 检查图片验证码
// @Description 检查图片验证码
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "用户名"
// @Param code formData string true "图片验证码"
// @Success 200 {object} _Response "图片验证码正确"
// @Failure 200 {object} _Response "图片验证码错误"
// @Failure 200 {object} _Response "服务器内部错误"
// @Router /check-picture-code [POST]
func CheckPictureCode(c *gin.Context) {
	username := c.PostForm("username")
	code := c.PostForm("code")
	ok, err := services.CheckCode(username, code)
	// 从 redis 中获取失败
	if err != nil {
		resp.ResponseError(c, resp.CodeUsernameNotExist)
		return
	}
	if !ok {
		resp.ResponseError(c, resp.CodePictureError)
		return
	}
	resp.ResponseSuccess(c, resp.CodeSuccess)

}
