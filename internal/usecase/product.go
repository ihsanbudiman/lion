package usecase

import (
	"context"
	customerror "lion/internal/domain/custom_error"
	"lion/internal/repository"

	"log"
)

type ProductUsecase interface {
	SyncProduct(ctx context.Context) error
}

type productUsecase struct {
	sourceProductRepo      repository.SourceProductRepo
	destinationProductRepo repository.DestinationProductRepo
}

func (p *productUsecase) SyncProduct(ctx context.Context) error {
	sourceProducts, err := p.sourceProductRepo.FindAll(ctx)
	if err != nil {
		return customerror.NewHTTPError(err.Error(), 500)
	}

	go func() {
		// Recover panic from goroutine, because panic on goroutine will stop the program
		defer func() {
			if r := recover(); r != nil {
				log.Default().Println("recovered from panic", r)
			}
		}()
		err = p.destinationProductRepo.SyncAll(ctx, sourceProducts)
		if err != nil {
			log.Default().Println(err)
		}

		log.Default().Println("sync done")
	}()

	return nil
}

func NewProductUsecase(sourceProductRepo repository.SourceProductRepo, destinationProductRepo repository.DestinationProductRepo) ProductUsecase {
	return &productUsecase{
		sourceProductRepo,
		destinationProductRepo,
	}
}
