package http

import (
	customerror "lion/internal/domain/custom_error"
	"lion/internal/usecase"
	"net/http"
)

type ProductHandler interface {
	SyncProduct(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	productUsecase usecase.ProductUsecase
}

func (p *productHandler) SyncProduct(w http.ResponseWriter, r *http.Request) {
	err := p.productUsecase.SyncProduct(r.Context())
	if err != nil {
		customerror.WriteHTTPResponse(w, err)
		return
	}

	JSON(w, "sync product success", nil)
}

func NewProductHandler(productUsecase usecase.ProductUsecase) ProductHandler {
	return &productHandler{
		productUsecase,
	}
}
