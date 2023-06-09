package model

import "time"

// 不包含逻辑删除的模型
type Universal struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id" mapstructure:"id"`
	CreatedAt time.Time `json:"created_at" mapstructure:"-"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"-"`
}
