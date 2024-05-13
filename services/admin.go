package services

import (
	"fmt"
	"go.uber.org/zap"
	"online-judge/dao/mysql"
	"online-judge/pkg/resp"
)

// DeleteUser 删除用户
func (u *UserService) DeleteUser() (response resp.Response) {
	// 检验是否有这个用户ID
	var UserIDCount int64
	err := mysql.CheckUserID(u.UserID, &UserIDCount)
	if err != nil {
		response.Code = resp.SearchDBError
		zap.L().Error("services-DeleteUser-CheckUserID ", zap.Error(err))
		return
	}
	if UserIDCount == 0 {
		response.Code = resp.NotExistUserID
		zap.L().Error("services-DeleteUser-CheckUserID "+
			fmt.Sprintf("do not have this userID %d ", u.UserID), zap.Error(err))
		return
	}
	// 删除用户
	err = mysql.DeleteUser(u.UserID)
	if err != nil {
		response.Code = resp.DBDeleteError
		zap.L().Error("services-DeleteUser-DeleteUser "+
			fmt.Sprintf("delete userID %d failed ", u.UserID), zap.Error(err))
		return
	}
	// 删除成功
	response.Code = resp.Success
	return
}

// AddAdmin 添加管理员
func (u *UserService) AddAdmin() (response resp.Response) {
	var UserIDCount int64
	err := mysql.CheckUserID(u.UserID, &UserIDCount)
	if err != nil {
		response.Code = resp.SearchDBError
		zap.L().Error("services-AddAdmin-CheckUserID ", zap.Error(err))
		return
	}
	if UserIDCount == 0 {
		response.Code = resp.NotExistUserID
		zap.L().Error("services-AddAdmin-CheckUserID "+
			fmt.Sprintf("do not have this userID %d ", u.UserID), zap.Error(err))
		return
	}
	err = mysql.AddAdminUser(u.UserID)
	if err != nil {
		response.Code = resp.SearchDBError
		zap.L().Error("services-AddAdmin-AddAdminUser ", zap.Error(err))
		return
	}
	response.Code = resp.Success

	return
}

// CreateProblem 创建题目
func (p *Problem) CreateProblem() (response resp.Response) {
	// 检查题目标题是否重复
	var problemNum int64
	err := mysql.CheckProblemTitle(p.Title, &problemNum)
	switch {
	case err != nil: // 搜索数据库错误
		response.Code = resp.SearchDBError
		zap.L().Error("services-CreateProblem-CheckProblemTitle ", zap.Error(err))
		return
	case problemNum > 0: // 题目已经存在
		response.Code = resp.ProblemAlreadyExist
		zap.L().Error("services-CreateProblem-CheckProblemTitle " +
			fmt.Sprintf("title %s aleardy exist", p.Title))
		return
	}
	var problems mysql.Problems
	problems.ProblemID = p.ProblemID
	problems.Title = p.Title
	problems.Content = p.Content
	problems.Difficulty = p.Difficulty
	problems.MaxRuntime = p.MaxRuntime
	problems.MaxMemory = p.MaxMemory

	// 提前转换类型
	var convertedTestCases []*mysql.TestCase
	for _, tc := range p.TestCases {
		// 进行类型转换
		convertedTestCases = append(convertedTestCases, &mysql.TestCase{
			TID:      tc.TID,
			PID:      tc.PID,
			Input:    tc.Input,
			Expected: tc.Expected,
		})
	}

	problems.TestCases = convertedTestCases
	// 创建题目
	err = mysql.CreateProblem(&problems)
	if err != nil {
		response.Code = resp.CreateProblemError
		zap.L().Error("services-CreateProblem-CreateProblem ", zap.Error(err))
		return
	}
	response.Code = resp.Success
	return
}

// UpdateProblem 更新题目
func (p *Problem) UpdateProblem() (response resp.Response) {
	// 检查题目id是否存在
	var idNum int64
	err := mysql.CheckProblemID(p.ProblemID, &idNum)
	switch {
	case err != nil: // 搜索数据库错误
		response.Code = resp.SearchDBError
		zap.L().Error("services-UpdateProblem-CheckProblemID ", zap.Error(err))
		return
	case idNum == 0: // 题目id不存在
		response.Code = resp.ProblemNotExist
		zap.L().Error("services-UpdateProblem-CheckProblemID " +
			fmt.Sprintf("problemID %s do not exist", p.ProblemID))
		return
	}

	// 检查题目标题是否存在
	var titleNum int64
	err = mysql.CheckProblemTitle(p.Title, &titleNum)
	switch {
	case err != nil: // 搜索数据库错误
		response.Code = resp.SearchDBError
		zap.L().Error("services-UpdateProblem-CheckProblemTitle", zap.Error(err))
		return
	case titleNum != 0: // 题目title已存在
		response.Code = resp.ProblemAlreadyExist
		zap.L().Error("services-UpdateProblem-CheckProblemTitle" +
			fmt.Sprintf("problem title %s already exist", p.Title))
		return
	}

	var problems mysql.Problems
	problems.ProblemID = p.ProblemID
	problems.Title = p.Title
	problems.Content = p.Content
	problems.Difficulty = p.Difficulty
	problems.MaxRuntime = p.MaxRuntime
	problems.MaxMemory = p.MaxMemory

	// 提前转换类型
	var convertedTestCases []*mysql.TestCase
	for _, tc := range p.TestCases {

		// 进行类型转换
		convertedTestCases = append(convertedTestCases, &mysql.TestCase{
			TID:      tc.TID,
			PID:      tc.PID,
			Input:    tc.Input,
			Expected: tc.Expected,
		})

	}
	problems.TestCases = convertedTestCases

	err = mysql.UpdateProblem(&problems)
	if err != nil {
		zap.L().Error("services-UpdateProblem-UpdateProblem ", zap.Error(err))
		response.Code = resp.InternalServerError
		return
	}
	response.Code = resp.Success
	return
}

// DeleteProblem 删除题目
func (p *Problem) DeleteProblem() (response resp.Response) {
	// 检查题目id是否存在
	var idNum int64
	err := mysql.CheckProblemID(p.ProblemID, &idNum)
	switch {
	case err != nil: // 搜索数据库错误
		response.Code = resp.SearchDBError
		zap.L().Error("services-DeleteProblem-CheckProblemID ", zap.Error(err))
		return
	case idNum == 0: // 题目id不存在
		response.Code = resp.ProblemNotExist
		zap.L().Error("services-DeleteProblem-CheckProblemID " +
			fmt.Sprintf("problemID %s do not exist", p.ProblemID))
		return
	}
	// 删除题目
	err = mysql.DeleteProblem(p.ProblemID)
	if err != nil {
		response.Code = resp.SearchDBError
		zap.L().Error("services-DeleteProblem-DeleteProblem  ", zap.Error(err))
		return
	}
	response.Code = resp.Success
	return
}