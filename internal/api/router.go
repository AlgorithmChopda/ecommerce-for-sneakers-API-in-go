package api

import (
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app"
	"github.com/gorilla/mux"
)

func NewRouter(deps app.Dependencies) *mux.Router {
	router := mux.NewRouter()

	// UserRoutes
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/login", LoginUserHandler(deps.UserService)).Methods(http.MethodPost)
	userRouter.HandleFunc("/register", RegisterUserHandler(deps.UserService)).Methods(http.MethodPost)

	// Admin
	adminRouter := router.PathPrefix("/seller").Subrouter()
	adminRouter.HandleFunc("", RegisterSellerHandler(deps.SellerService)).Methods(http.MethodPost)

	// Product
	productRouter := router.PathPrefix("/product").Subrouter()
	productRouter.HandleFunc("", CreateProductHandler(deps.ProductService)).Methods(http.MethodPost)
	productRouter.HandleFunc("/{id}", GetProductHandler(deps.ProductService)).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", UpdateProductHandler(deps.ProductService)).Methods(http.MethodPatch)

	// Order / Cart
	orderRouter := router.PathPrefix("/cart").Subrouter()
	orderRouter.HandleFunc("", CreateOrderHandler(deps.OrderService)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/{id}/order", PlaceOrderHandler(deps.OrderService)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/{id}/{productDetailId}", AddOrderHandler(deps.OrderService)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/{id}/{productDetailId}", UpdateOrderItemHandler(deps.OrderService)).Methods(http.MethodPut)

	//PlaceOrderHandler
	return router
}
