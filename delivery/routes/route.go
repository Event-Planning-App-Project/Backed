package routes

import (
	"event/delivery/controller/category"
	"event/delivery/controller/comment"

	"event/delivery/controller/transaction"
	"event/delivery/controller/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Path(e *echo.Echo, u user.ControllerUser, t transaction.TransController, cat category.CategoryControl, co comment.CommentControll) {

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// Login
	e.POST("/login", u.Login())
	// ROUTES USER
	user := e.Group("/user")
	user.POST("", u.InsertUser()) // Register
	// user.GET("", u.GetAllUser, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.GET("/:id", u.GetUserbyID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.PUT("/:id", u.UpdateUserID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.DELETE("/:id", u.DeleteUserID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))

	category := e.Group("/category")
	category.PUT("/:id", cat.UpdateCat(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	category.DELETE("/:id", cat.DeleteCat(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	category.POST("", cat.CreateCategory(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	category.GET("", cat.GetAllCategory())
	category.GET("/:id", cat.GetCategoryID())

	comment := e.Group("/comment")
	comment.PUT("/:id", co.UpdateComment(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	comment.DELETE("/:id", co.DeleteComment(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	comment.POST("", co.CreateComment(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	comment.GET("", co.GetAllComment())
	comment.GET("/:id", co.GetCommentID())

	// ROUTES TRANSACTION
	Transaction := e.Group("/transaction")
	Transaction.POST("", t.CreateTransaction(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Transaction.GET("", t.GetAllTransaction(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Transaction.GET("/:order_id", t.GetTransactionDetail(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Transaction.POST("/:order_id/pay", t.PayTransaction(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Transaction.POST("/:order_id/cancel", t.CancelTransaction(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Transaction.GET("/:order_id/finish_payment", t.FinishPayment())
}
