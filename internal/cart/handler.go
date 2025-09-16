package cart

import (
	"links/internal/cart/models"
	"links/pkg/api"
	"net/http"
	"strconv"
)

type CartHandler struct {
	CartRepository *CartRepository
}

func NewCartHandler(router *http.ServeMux, handler *CartHandler) {
	router.HandleFunc("POST /product/create", handler.CreateProduct())
	router.HandleFunc("GET /product/{uid}", handler.GetProduct())
	router.HandleFunc("PUT /product/{uid}", handler.UpdateProduct())
	router.HandleFunc("DELETE /product/{uid}", handler.DeleteProduct())
}

func (handler *CartHandler) CreateProduct() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		r, err := api.HandleBody[ProductCreateRequest](&writer, request)
		if err != nil {
			return
		}
		product, err := handler.CartRepository.CreateProduct(&models.Product{
			ProductId:   r.ProductId,
			Name:        r.Name,
			Description: r.Description,
			Price:       float64(r.Price),
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		api.Json(writer, product, 200)
	}
}

func (handler *CartHandler) UpdateProduct() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		r, err := api.HandleBody[ProductCreateRequest](&writer, request)
		if err != nil {
			return
		}
		product, err := handler.CartRepository.UpdateProduct(&models.Product{
			ProductId:   r.ProductId,
			Name:        r.Name,
			Description: r.Description,
			Price:       float64(r.Price),
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		api.Json(writer, product, 200)
	}
}

func (handler *CartHandler) GetProduct() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		productId := request.PathValue("uid")
		uid, err := strconv.Atoi(productId)
		if err != nil {
			api.Json(writer, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := handler.CartRepository.GetProduct(&models.Product{
			ProductId: uid,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		api.Json(writer, product, 200)
	}
}

func (handler *CartHandler) DeleteProduct() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		productId := request.PathValue("uid")
		uid, err := strconv.Atoi(productId)
		if err != nil {
			api.Json(writer, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := handler.CartRepository.DeleteProduct(&models.Product{
			ProductId: uid,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		api.Json(writer, product, 200)
	}
}
