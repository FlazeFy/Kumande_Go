package admin

import (
	"kumande/models"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByEmail(email string) (*models.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) FindByEmail(email string) (*models.Admin, error) {
	// Models
	var admin models.Admin

	// Query
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}

	return &admin, nil
}
