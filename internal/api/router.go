package api

import (
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/constants"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(deps app.Dependencies) *mux.Router {
	router := mux.NewRouter()

	// UserRoutes
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/login", LoginUserHandler(deps.UserService)).Methods(http.MethodPost)
	userRouter.HandleFunc("/register", RegisterUserHandler(deps.UserService)).Methods(http.MethodPost)

	userRouter.HandleFunc("", middleware.CheckAuth(GetUserListHandler(deps.UserService), constants.ADMIN)).Methods(http.MethodGet)
	userRouter.HandleFunc("/profile", middleware.CheckAuth(GetUserProfileHandler(deps.UserService), constants.ALL)).Methods(http.MethodGet)

	// Seller
	sellerRouter := router.PathPrefix("/seller").Subrouter()
	sellerRouter.HandleFunc("", middleware.CheckAuth(GetAllSellersHandler(deps.SellerService), constants.ADMIN)).Methods(http.MethodGet)
	sellerRouter.HandleFunc("/register", middleware.CheckAuth(RegisterSellerHandler(deps.SellerService), constants.ADMIN)).Methods(http.MethodPost)
	sellerRouter.HandleFunc("/{id}", middleware.CheckAuth(DeleteSellerHandler(deps.SellerService), constants.ADMIN)).Methods(http.MethodDelete)

	// Product
	productRouter := router.PathPrefix("/product").Subrouter()
	productRouter.HandleFunc("", GetProductWithFilterHandler(deps.ProductService)).Methods(http.MethodGet)
	productRouter.HandleFunc("", middleware.CheckAuth(CreateProductHandler(deps.ProductService), constants.SELLER)).Methods(http.MethodPost)
	productRouter.HandleFunc("/{id}", GetProductHandler(deps.ProductService)).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", middleware.CheckAuth(UpdateProductHandler(deps.ProductService), constants.SELLER)).Methods(http.MethodPatch)

	// Order / Cart
	orderRouter := router.PathPrefix("/cart").Subrouter()
	orderRouter.HandleFunc("", middleware.CheckAuth(CreateOrderHandler(deps.OrderService), constants.BUYER)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/{id}", middleware.CheckAuth(GetAllOrderItemsHandler(deps.OrderService), constants.BUYER)).Methods(http.MethodGet)
	orderRouter.HandleFunc("/{id}/order", middleware.CheckAuth(PlaceOrderHandler(deps.OrderService), constants.BUYER)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/{id}/product/{productDetailId}", middleware.CheckAuth(AddOrderHandler(deps.OrderService), constants.BUYER)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/{id}/product/{productDetailId}", middleware.CheckAuth(UpdateOrderItemHandler(deps.OrderService), constants.BUYER)).Methods(http.MethodPut)
	// TODO cart/{id}/count service and api remaining

	orderPlacedRouted := router.PathPrefix("/order").Subrouter()
	orderPlacedRouted.HandleFunc("", middleware.CheckAuth(GetUserPlacedOrderHandler(deps.OrderService), constants.BUYER)).Methods(http.MethodGet)
	orderPlacedRouted.HandleFunc("/{id}", middleware.CheckAuth(GetPlacedOrderDetailsHandler(deps.OrderService), constants.BUYER)).Methods(http.MethodGet)

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/register", middleware.CheckAuth(RegisterAdminHandler(deps.UserService), constants.ADMIN)).Methods(http.MethodPost)

	return router
}
