package models

import (
	"database/sql"
)

type SpecialDetail struct {
	Id          uint64       `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	SpecialId   uint64       `gorm:"column:special_id;type:bigint(20) unsigned;default:0;NOT NULL" json:"special_id"` // 专题id
	Description string       `gorm:"column:description;type:varchar(500)" json:"description"`                         // 专题简介
	JsonData    string       `gorm:"column:json_data;type:longtext" json:"json_data"`                                 // 专题json内容
	HtmlData    string       `gorm:"column:html_data;type:longtext" json:"html_data"`                                 // 专题html内容
	CreatedAt   sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt   sql.NullTime `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt   sql.NullTime `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

func (s SpecialDetail) table() string {
	return "special_details"
}
