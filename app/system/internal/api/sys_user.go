package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/api"
	"github.com/mars-projects/mars/common/middleware/authentication"
	"gorm.io/gorm"
	"net/http"
)

type SysUserHandler struct {
	api.Api
	biz     *biz.SysUser
	roleBiz *biz.SysRole
}

func NewSysUserHandler(userBiz *biz.SysUser) *SysUserHandler {
	h := api.Api{}
	return &SysUserHandler{biz: userBiz, Api: h}
}

func (e *SysUserHandler) Get(ctx *gin.Context) {
	var err error
	var data dto.SysUserById
	var model models.SysUser

	err = e.MakeContext(ctx).
		Bind(&data, nil).
		Errors
	if err != nil {
		e.InternalErrorResult(err)
	}
	err = e.Bind(&data, nil).Errors
	if err != nil {
		e.ErrorResult(400, err, "参数错误")
		return
	}
	err = e.biz.Get(&data, &model)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		e.ErrorResult(404, err, "该用户不存在")
		return
	}
	e.Result(&model, "查询成功")
}

// GetPage
// @Summary 列表用户信息数据
// @Description 获取JSON
// @Tags 用户
// @Param username query string false "username"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user [get]
// @Security Bearer
func (e SysUserHandler) GetPage(c *gin.Context) {
	req := dto.SysUserGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Query).
		Errors
	if err != nil {
		e.InternalErrorResult(err)
		return
	}

	list := make([]models.SysUser, 0)
	var count int64

	err = e.biz.GetPage(&req, &list, &count)

	if err != nil {
		e.ErrorResult(400, err, err.Error())
		return
	}
	e.PageResult(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Insert
// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserInsertReq true "用户数据"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user [post]
// @Security Bearer
func (e SysUserHandler) Insert(c *gin.Context) {
	fmt.Println("创建成功")
	req := dto.SysUserInsertReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(c.GetInt(authentication.UserId))
	err = e.biz.Insert(&req)
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	e.Result(req.GetId(), "创建成功")
}

// Update
// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [put]
// @Security Bearer
func (e SysUserHandler) Update(c *gin.Context) {
	req := dto.SysUserUpdateReq{}
	err := e.MakeContext(c).
		Bind(&req).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	req.SetCreateBy(c.GetInt(authentication.UserId))

	err = e.biz.Update(&req)
	if err != nil {
		return
	}
	e.Result(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Param userId path int true "userId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [delete]
// @Security Bearer
func (e SysUserHandler) Delete(c *gin.Context) {
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.InternalErrorResult(err)
		return
	}
	// 设置编辑人
	req.SetCreateBy(c.GetInt(authentication.UserId))

	err = e.biz.Remove(&req)
	if err != nil {
		e.ErrorResult(403, err, err.Error())
		return
	}
	e.Result(req.GetId(), "删除成功")
}

// InsetAvatar
// @Summary 修改头像
// @Description 获取JSON
// @Tags 个人中心
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/avatar [post]
// @Security Bearer
func (e SysUserHandler) InsetAvatar(c *gin.Context) {
	req := dto.UpdateSysUserAvatarReq{}
	err := e.MakeContext(c).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	// 数据权限检查
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	guid := uuid.New().String()
	filPath := "static/uploadfile/" + guid + ".jpg"
	for _, file := range files {
		// 上传文件至指定目录
		err = c.SaveUploadedFile(file, filPath)
		if err != nil {
			e.ErrorResult(500, err, "")
			return
		}
	}
	req.SetCreateBy(c.GetInt(authentication.UserId))

	req.Avatar = "/" + filPath

	err = e.biz.UpdateAvatar(&req)
	if err != nil {
		return
	}
	e.Result(filPath, "修改成功")
}

// UpdateStatus 修改用户状态
// @Summary 修改用户状态
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpdateSysUserStatusReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/status [put]
// @Security Bearer
func (e SysUserHandler) UpdateStatus(c *gin.Context) {
	req := dto.UpdateSysUserStatusReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	req.SetCreateBy(c.GetInt(authentication.UserId))

	err = e.biz.UpdateStatus(&req)
	if err != nil {
		return
	}
	e.Result(req.GetId(), "更新成功")
}

// ResetPwd 重置用户密码
// @Summary 重置用户密码
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.ResetSysUserPwdReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/pwd/reset [put]
// @Security Bearer
func (e SysUserHandler) ResetPwd(c *gin.Context) {
	req := dto.ResetSysUserPwdReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	req.SetCreateBy(c.GetInt(authentication.UserId))

	err = e.biz.ResetPwd(&req)
	if err != nil {
		e.ErrorResult(403, err, err.Error())
		return
	}
	e.Result(req.GetId(), "更新成功")
}

// UpdatePwd
// @Summary 重置密码
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.PassWord true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/pwd/set [put]
// @Security Bearer
func (e SysUserHandler) UpdatePwd(c *gin.Context) {
	req := dto.PassWord{}
	err := e.MakeContext(c).
		Bind(&req).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	// 数据权限检查
	err = e.biz.UpdatePwd(c.GetInt(authentication.UserId), req.OldPassword, req.NewPassword)
	if err != nil {
		e.ErrorResult(http.StatusForbidden, err, "密码修改失败")
		return
	}
	e.Result(nil, "密码修改成功")
}

// GetProfile
// @Summary 获取个人中心用户
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/profile [get]
// @Security Bearer
func (e SysUserHandler) GetProfile(c *gin.Context) {
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	req.SetCreateBy(c.GetInt(authentication.UserId))

	sysUser := models.SysUser{}
	roles := make([]models.SysRole, 0)
	posts := make([]models.SysPost, 0)
	err = e.biz.GetProfile(&req, &sysUser, &roles, &posts)
	if err != nil {
		e.ErrorResult(500, err, "获取用户信息失败")
		return
	}
	e.Result(gin.H{
		"user":  sysUser,
		"roles": roles,
		"posts": posts,
	}, "查询成功")
}

// GetInfo
// @Summary 获取个人信息
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/getinfo [get]
// @Security Bearer
func (e SysUserHandler) GetInfo(c *gin.Context) {
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	var roles = make([]string, 1)
	//roles[0] = user.GetRoleName(c)
	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"
	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"
	roles[0] = "admin"
	var mp = make(map[string]interface{})
	mp["roles"] = roles
	//if user.GetRoleName(c) == "admin" || user.GetRoleName(c) == "系统管理员" {
	//	mp["permissions"] = permissions
	//	mp["buttons"] = buttons
	//} else {
	//	list, _ := e.roleBiz.GetById(user.GetRoleId(c))
	//	mp["permissions"] = list
	//	mp["buttons"] = list
	//}
	mp["permissions"] = permissions
	mp["buttons"] = buttons
	sysUser := models.SysUser{}
	req.Id = c.GetInt(authentication.UserId)
	fmt.Println(req.GetId())
	err = e.biz.Get(&req, &sysUser)
	if err != nil {
		e.ErrorResult(http.StatusUnauthorized, err, "登录失败")
		return
	}
	mp["introduction"] = " am a super administrator"
	mp["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if sysUser.Avatar != "" {
		mp["avatar"] = sysUser.Avatar
	}
	mp["userName"] = sysUser.NickName
	mp["userId"] = sysUser.UserId
	mp["deptId"] = sysUser.DeptId
	mp["name"] = sysUser.NickName
	mp["code"] = 200
	e.Result(mp, "")
}
