package repository

import (
	"context"
	"lion/internal/domain/models"

	"gorm.io/gorm"
)

type SourceProductRepo interface {
	FindAll(ctx context.Context) ([]models.SourceProduct, error)
}

type sourceProductRepo struct {
	db *gorm.DB
}

// FindAll implements SourceProductRepo.
func (s *sourceProductRepo) FindAll(ctx context.Context) ([]models.SourceProduct, error) {
	var sourceProducts []models.SourceProduct
	if err := s.db.WithContext(ctx).Find(&sourceProducts).Error; err != nil {
		return nil, err
	}
	return sourceProducts, nil
}

func NewSourceProductRepo(db *gorm.DB) SourceProductRepo {
	return &sourceProductRepo{db}
}
