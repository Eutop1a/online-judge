package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"online-judge/pkg/resp"
	"online-judge/services"
)

// GetProblemList 获取题目列表接口
// @Tags Problem API
// @Summary 获取题目列表
// @Description 获取题目列表接口
// @Param Authorization header string true "token"
// @Success 200 {object} _Response "获取题目列表成功"
// @Failure 200 {object} _Response "服务器内部错误"
// @Router /problem/list [GET]
func GetProblemList(c *gin.Context) {
	var getProblemList services.Problem
	data, err := getProblemList.GetProblemList()
	if err != nil {
		resp.ResponseError(c, resp.CodeInternalServerError)
		zap.L().Error("controller-GetProblemList-GetProblemList ", zap.Error(err))
		return
	}
	resp.ResponseSuccess(c, data)
}

// GetProblemDetail 获取单个题目详细接口
// @Tags Problem API
// @Summary 获取单个题目详细
// @Description 获取单个题目详细接口
// @Accept multipart/form-data
// @Produce json,multipart/form-data
// @Param Authorization header string true "token"
// @Param problem_id path string true "题目ID"
// @Success 200 {object} _Response "获取成功"
// @Failure 200 {object} _Response "题目ID不存在"
// @Failure 200 {object} _Response "服务器内部错误"
// @Router /problem/{problem_id} [GET]
func GetProblemDetail(c *gin.Context) {
	var getProblemDetail services.Problem
	pid := c.Param("problem_id")
	getProblemDetail.ProblemID = pid

	data, err := getProblemDetail.GetProblemDetail()
	if err != nil {
		resp.ResponseError(c, resp.CodeProblemIDNotExist)
		zap.L().Error("controller-GetProblemDetail-GetProblemDetail ", zap.Error(err))
		return
	}
	resp.ResponseSuccess(c, data)
}

// GetProblemID 获取题目ID接口
// @Tags Problem API
// @Summary 获取题目ID
// @Description 获取题目ID接口
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "token"
// @Param title formData string true "标题"
// @Success 200 {object} _Response "获取题目ID成功"
// @Failure 200 {object} _Response "题目title不存在"
// @Router /problem/id [POST]
func GetProblemID(c *gin.Context) {
	var getProblemID services.Problem
	title := c.PostForm("title")
	getProblemID.Title = title
	uid, err := getProblemID.GetProblemID()
	if err != nil {
		resp.ResponseError(c, resp.CodeProblemTitleNotExist)
		//zap.L().Error("controller-GetProblemID-GetProblemID ", zap.Error(err))
		return
	}
	resp.ResponseSuccess(c, uid)
}
