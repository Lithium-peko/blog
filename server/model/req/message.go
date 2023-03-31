package req

// 新增留言
type AddMessage struct {
	Nickname string `json:"nickname" validate:"required"`
	Avatar   string `json:"avatar"`
	Content  string `json:"content" validate:"required"`
	Speed    int    `json:"speed"`
}

type GetMessages struct {
	PageQuery
	Nickname string `form:"nickname"`
	IsReview *int8  `form:"is_review"`
}
