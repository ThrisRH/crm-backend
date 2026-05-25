package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          uuid.UUID `gorm:"column:id;type:int;not null;primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name;type:varchar(255);not null;uniqueIndex"`
	Description string    `gorm:"column:description;type:varchar(255);default:''"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
