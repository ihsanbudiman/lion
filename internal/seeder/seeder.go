package seeder

import (
	"fmt"
	"lion/internal/domain/models"
	"math"
	"math/rand/v2"

	"gorm.io/gorm"
)

func Seed(dbSource, dbDestination *gorm.DB) {

	sourceProducts := []models.SourceProduct{}
	destinationProducts := []models.DestinationProduct{}

	for i := 0; i < 500; i++ {
		productName := fmt.Sprintf("Product No.%d", i+1)

		qty := rand.IntN(1000-10+1) + 10
		sellingPrice := math.Ceil(10000 + rand.Float64()*(50000-10000))
		promoPrice := math.Ceil(sellingPrice * 0.9)
		sourceProducts = append(sourceProducts, models.SourceProduct{
			ID:           uint(i + 1),
			ProductName:  productName,
			Qty:          qty,
			SellingPrice: sellingPrice,
			PromoPrice:   promoPrice,
		})

		destinationProducts = append(destinationProducts, models.DestinationProduct{
			ID:           uint(i + 1),
			ProductName:  productName,
			Qty:          0,
			SellingPrice: 0,
			PromoPrice:   0,
		})
	}

	txSource := dbSource.Begin()
	txDestination := dbDestination.Begin()

	err := txSource.Save(&sourceProducts).Error
	if err != nil {
		txSource.Rollback()
		txDestination.Rollback()
		panic(err)
	}

	err = txDestination.Save(&destinationProducts).Error
	if err != nil {
		txSource.Rollback()
		txDestination.Rollback()
		panic(err)
	}

	err = txSource.Commit().Error
	if err != nil {
		txSource.Rollback()
		txDestination.Rollback()
		panic(err)
	}

	err = txDestination.Commit().Error
	if err != nil {
		txSource.Rollback()
		txDestination.Rollback()
		panic(err)
	}

}
