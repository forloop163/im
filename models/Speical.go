package models

import (
	"database/sql"
)

type Special struct {
	ID           uint64         `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	SiteID       uint64         `gorm:"column:site_id;default:0;NOT NULL" json:"site_id"`               // 站点id
	Platform     int            `gorm:"column:platform;default:1;NOT NULL" json:"platform"`             // 专题类型：1 PC，2 H5
	Name         string         `gorm:"column:name;NOT NULL" json:"name"`                               // 专题名称
	PreviewUrl   string         `gorm:"column:preview_url;NOT NULL" json:"preview_url"`                 // 预览地址
	CloneID      uint64         `gorm:"column:clone_id;default:0;NOT NULL" json:"clone_id"`             // 克隆专题ID
	Status       int            `gorm:"column:status;default:1;NOT NULL" json:"status"`                 // 状态：1待审核，2审核通过，3发布
	IsPublic     int            `gorm:"column:is_public;default:0;NOT NULL" json:"is_public"`           // 是否公开：1是0否
	ClientID     int            `gorm:"column:client_id;default:1;NOT NULL" json:"client_id"`           // 客户端ID:1原创官网；2整装官网；3直投官网；4自媒体官网
	ViewsNum     int64          `gorm:"column:views_num;default:0;NOT NULL" json:"views_num"`           // 浏览量
	SignNum      int64          `gorm:"column:sign_num;default:0;NOT NULL" json:"sign_num"`             // 报名人数
	UserID       int64          `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`               // 添加人id
	UpdateUserID int64          `gorm:"column:update_user_id;default:0;NOT NULL" json:"update_user_id"` // 最后编辑人id
	CreatedAt    sql.NullTime   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    sql.NullTime   `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    sql.NullTime   `gorm:"column:deleted_at" json:"deleted_at"`
	Detail       *SpecialDetail `gorm:"FOREIGNKEY:special_id;ASSOCIATION_FOREIGNKEY:id" json:"Detail,omitempty"`
}
