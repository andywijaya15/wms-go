package models

import (
	"time"
)

type Factory struct {
	ID          int        `gorm:"primaryKey;type:int;not null"`
	CreatedAt   *time.Time `gorm:"type:timestamp(0)"`
	UpdatedAt   *time.Time `gorm:"type:timestamp(0)"`
	IsActive    bool       `gorm:"type:boolean;default:true;not null"`
	AdOrgID     int        `gorm:"type:int;not null"`
	FactoryName string     `gorm:"type:varchar(191);not null"`
	IsFabric    bool       `gorm:"type:boolean;not null"`
}
