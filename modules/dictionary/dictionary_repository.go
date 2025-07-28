package dictionary

import (
	"kumande/models"
	"kumande/utils"

	"gorm.io/gorm"
)

// Dictionary Interface
type DictionaryRepository interface {
	FindAllDictionary(pagination utils.Pagination) ([]models.Dictionary, int, error)
}

// Dictionary Struct
type dictionaryRepository struct {
	db *gorm.DB
}

// Dictionary Constructor
func NewDictionaryRepository(db *gorm.DB) DictionaryRepository {
	return &dictionaryRepository{db: db}
}

func (r *dictionaryRepository) FindAllDictionary(pagination utils.Pagination) ([]models.Dictionary, int, error) {
	// Model
	var total int
	var dictionaries []models.Dictionary

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit

	// Query
	err := r.db.Order("dictionary_type ASC").
		Order("dictionary_name ASC").
		Limit(pagination.Limit).
		Offset(offset).
		Find(&dictionaries).Error

	if err != nil {
		return nil, 0, err
	}

	total = len(dictionaries)
	return dictionaries, total, nil
}
