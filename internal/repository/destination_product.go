package repository

import (
	"context"
	"lion/internal/domain/models"

	"gorm.io/gorm"
)

type DestinationProductRepo interface {
	SyncAll(context.Context, []models.SourceProduct) error
}

type destinationProductRepo struct {
	db *gorm.DB
}

// FindAll implements SourceProductRepo.
func (s *destinationProductRepo) SyncAll(ctx context.Context, sourceProducts []models.SourceProduct) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, sourceProduct := range sourceProducts {
		destinationProduct := models.DestinationProduct{
			ID:           sourceProduct.ID,
			ProductName:  sourceProduct.ProductName,
			Qty:          sourceProduct.Qty,
			SellingPrice: sourceProduct.SellingPrice,
			PromoPrice:   sourceProduct.PromoPrice,
		}
		if err := tx.Save(&destinationProduct).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func NewDestinationProductRepo(db *gorm.DB) DestinationProductRepo {
	return &destinationProductRepo{db}
}
