package domain

import (
	"time"

	domainCat "github.com/dwikikf/agviano-core-api-golang/internal/domain/category"
	"gorm.io/gorm"
)

// gorm sementara untuk percepat development

type Product struct {
	ID          uint64             `json:"id" gorm:"primaryKey;autoIncrement"`
	CategoryID  uint64             `json:"category_id" gorm:"not null"`
	Name        string             `json:"name" gorm:"type:varchar(150);not null"`
	Slug        string             `json:"slug" gorm:"type:varchar(100);not null;unique"`
	Description string             `json:"description" gorm:"type:text"`
	Price       float64            `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock       uint               `json:"stock" gorm:"not null"`
	ImageURL    string             `json:"image_url" gorm:"type:varchar(255)"`
	IsActive    bool               `json:"is_active" gorm:"default:false"`
	CreatedAt   time.Time          `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time          `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt     `json:"deleted_at" gorm:"index"`
	Category    domainCat.Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
