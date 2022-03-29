package models

import (
	"time"
)

type User struct {
	Id        uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(100)" json:"name"`                   // 用户名
	Password  string    `gorm:"column:password;type:varchar(100)" json:"password,omitempty"` // 密码
	Email     string    `gorm:"column:email;type:varchar(100)" json:"email"`                 // 邮箱
	Mobile    string    `gorm:"column:mobile;type:varchar(32)" json:"mobile"`                // 手机号码
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at,omitempty"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}
