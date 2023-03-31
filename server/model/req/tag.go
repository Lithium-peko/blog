package req

// 新增 / 编辑 标签对象
type AddOrEditTag struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required" label:"标签名称"`
}
