package dto

const (
	GenerateIDMaxNum int = 1000
)

type GenerateIDRequest struct {
	Num int `form:"num" binding:"required,min=1,max=1000"` // 生成id数量,最多一次一千
}

type GenerateIDResponse struct {
	Numbers []int64 `json:"numbers"`
}
