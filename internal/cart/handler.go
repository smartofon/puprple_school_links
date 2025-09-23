package cart

import (
	"fmt"
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

	router.HandleFunc("POST /auth/login", handler.LoginHandler())
	router.HandleFunc("POST /auth/confirm", handler.Confirm())
	router.Handle("POST /auth/test", auth.IsAuthorized(handler.UserRepository, handler.Test()))

	router.Handle("POST /order", auth.IsAuthorized(handler.UserRepository, handler.CreateOrder()))
	router.Handle("POST /order/{id}", auth.IsAuthorized(handler.UserRepository, handler.GetOrder()))
	router.Handle("GET /my-orders", auth.IsAuthorized(handler.UserRepository, handler.OrderList()))
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

func (handler *CartHandler) Test() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		answer := struct {
			Operation  string
			Phone      string
			Authorized any
			ID         any
		}{
			Operation:  "test",
			Phone:      r.Context().Value("user_phone").(string),
			Authorized: r.Context().Value("authorized"),
			ID:         r.Context().Value("user_id"),
		}
		api.Json(writer, answer, 200)
	})
}

func (handler *CartHandler) CreateOrder() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		r, err := api.HandleBody[OrderCreateResponse](&writer, request)
		if err != nil {
			return
		}
		user := request.Context().Value("user_id").(int)
		order, err := handler.CartRepository.CreateOrder(r.Products, user)
		answer := AnswerTemplate{
			Result:  "ok",
			Message: "",
		}
		if err != nil {
			answer.Result = "error"
			answer.Message = err.Error()
			api.Json(writer, answer, http.StatusInternalServerError)
		}
		answer.Message = fmt.Sprintf("Create order number #%v", order.ID)
		api.Json(writer, answer, 200)
	}
}

func (handler *CartHandler) GetOrder() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		s := request.PathValue("id")
		orderID, err := strconv.Atoi(s)
		user := request.Context().Value("user_id").(int)
		if err != nil {
			api.Json(writer, struct{}{}, http.StatusBadRequest)
			return
		}

		answer := AnswerTemplate{
			Result:  "ok",
			Message: "",
		}

		order, err := handler.CartRepository.GetOrder(orderID, user)
		if err != nil {
			answer.Result = "error"
			answer.Message = err.Error()
			api.Json(writer, answer, http.StatusBadRequest)
			return
		}

		api.Json(writer, order, 200)
	}
}

func (handler *CartHandler) OrderList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user := request.Context().Value("user_id").(int)

		answer := AnswerTemplate{
			Result:  "ok",
			Message: "",
		}

		orders, err := handler.CartRepository.OrderList(user)
		if err != nil {
			answer.Result = "error"
			answer.Message = err.Error()
			api.Json(writer, answer, http.StatusBadRequest)
			return
		}

		api.Json(writer, orders, 200)
	}
}
