package cart

import (
	"links/internal/auth"
	"links/internal/cart/models"
	"links/internal/user"
	"links/pkg/api"
	"net/http"
	"strconv"
)

type CartHandler struct {
	CartRepository *CartRepository
	UserRepository *user.UserRepository
	AuthHService   *auth.AuthHService
}

func NewCartHandler(router *http.ServeMux, handler *CartHandler) {
	router.HandleFunc("POST /product/create", handler.CreateProduct())
	router.HandleFunc("GET /product/{uid}", handler.GetProduct())
	router.HandleFunc("PUT /product/{uid}", handler.UpdateProduct())
	router.HandleFunc("DELETE /product/{uid}", handler.DeleteProduct())

	router.HandleFunc("GET /auth/login", handler.LoginHandler())
	router.HandleFunc("POST /auth/confirm", handler.Confirm())
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
			Images:      r.Images,
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

type AuthHandler struct {
	AuthHService *auth.AuthHService
}

func (handler *CartHandler) LoginHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		r, err := api.HandleBody[auth.PhoneAuthRequest](&writer, request)
		if err != nil {
			return
		}

		user, err := handler.AuthHService.Login(r.Phone)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		answer := auth.PhoneAuthResponce{SessionId: user.SessionId}
		api.Json(writer, answer, 200)
		api.SendSMS(user.Phone, user.ConfirmCode)
	}
}

func (handler *CartHandler) Confirm() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		r, err := api.HandleBody[auth.PhoneAuthCodeRequest](&writer, request)
		if err != nil {
			return
		}
		token, err := handler.AuthHService.Confirm(r.SessionId, r.Code)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return
		}
		answer := auth.PhoneAuthCodeResponce{Token: token}
		api.Json(writer, answer, 200)
	}
}
