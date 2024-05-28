package servicev1

import (
	"github.com/gin-gonic/gin"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/internal/api/dto"
	http "github.com/yituoshiniao/gin-api-http/internal/api/http"
	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2/application/service"
)

type UserScoreSrv struct {
	response     *http.Response
	userScoreSrv *service.UserScoreSrv
}

func NewUserScoreSrv(
	response *http.Response,
	userScoreSrv *service.UserScoreSrv,

) *UserScoreSrv {
	return &UserScoreSrv{
		response:     response,
		userScoreSrv: userScoreSrv,
	}
}

// Add
//
//	@Summary		添加数据
//	@Description	gorm 添加数据
//	@Tags			user_score
//	@Param			object	query		dto.UserScoreRequest	1	"request"
//	@Success		200		{object}	http.ResponseData		"请求成功"
//	@Router			/v1/userScore/add [POST]
func (p *UserScoreSrv) Add(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.UserScoreRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		xlog.S(ctx).Errorw("参数绑定错误", "err", err)
		p.response.Process(c, "", err)
		return
	}

	var err error
	id, err := p.userScoreSrv.AddOne(c.Request.Context(), "22")
	if err != nil {
		p.response.Error(c, err.Error(), id)
		return
	}
	p.response.Process(c, id, nil)
}

// Del
//
//	@Summary		删除数据
//	@Description	gorm 删除数据
//	@Tags			user_score
//	@Param			object	query		dto.UserScoreRequest	1	"request"
//	@Success		200		{object}	http.ResponseData		"请求成功"
//	@Router			/v1/userScore/del [POST]
func (p *UserScoreSrv) Del(c *gin.Context) {
	var err error
	err = p.userScoreSrv.DelId(c.Request.Context(), 11)
	if err != nil {
		p.response.Error(c, err.Error(), nil)
		return
	}
	p.response.Process(c, nil, nil)
}

// Update
//
//	@Summary		更新数据
//	@Description	gorm 更新数据
//	@Tags			user_score
//	@Param			object	query		dto.UserScoreRequest	1	"request"
//	@Success		200		{object}	http.ResponseData		"请求成功"
//	@Router			/v1/userScore/update [POST]
func (p *UserScoreSrv) Update(c *gin.Context) {
	var err error
	var scoreResult int32 = 11
	var score int32 = 22
	err = p.userScoreSrv.UpdateStatus(c.Request.Context(), "22", scoreResult, score)
	if err != nil {
		p.response.Error(c, err.Error(), nil)
		return
	}
	p.response.Process(c, nil, nil)
}

// Find
//
//	@Summary		查询一条数据
//	@Description	gorm 查询一条数据
//	@Tags			user_score
//	@Param			object	query		dto.UserScoreRequest		1	"request"
//	@Success		200		{object}	dto.UserScoreFindResponse	"请求成功"
//	@Router			/v1/userScore/find [GET]
func (p *UserScoreSrv) Find(c *gin.Context) {
	var err error
	var userId string = "22"
	item, err := p.userScoreSrv.FindOne(c.Request.Context(), userId)
	if err != nil {
		p.response.Error(c, err.Error(), nil)
		return
	}
	p.response.Process(c, item, nil)
}

// List
//
//	@Summary		查询列表
//	@Description	gorm 查询列表
//	@Tags			user_score
//	@Param			object	query		dto.UserScoreRequest		1	"request"
//	@Success		200		{object}	dto.UserScoreListResponse	"请求成功"
//	@Router			/v1/userScore/list [GET]
func (p *UserScoreSrv) List(c *gin.Context) {
	var err error
	var userId string = "22"
	var ids []int32 = []int32{1, 3, 4}
	list, err := p.userScoreSrv.Find(c.Request.Context(), userId, ids)
	if err != nil {
		p.response.Error(c, err.Error(), nil)
		return
	}
	p.response.Process(c, list, nil)
}
