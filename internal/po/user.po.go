package po

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `gorm:"column:uuid;type:uuid;default:gen_random_uuid();primaryKey"`
	Email    string    `gorm:"column:email;type:varchar(255);not null;uniqueIndex"`
	UserName string    `gorm:"column:user_name;type:varchar(255);not null;uniqueIndex"`
	Password string    `gorm:"column:password;type:varchar(255);not null"`
	Role     []Role    `gorm:"many2many:user_roles;"`

	CreateAt time.Time `gorm:"column:create_at;type:timestamp"`
	UpdateAt time.Time `gorm:"column:update_at;type:timestamp"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
