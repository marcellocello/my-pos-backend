package router

import (
	"mypos-backend/internal/handler"
	"mypos-backend/internal/middleware"
	"mypos-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Use(middleware.EnableCORS())

	r.GET("/api/health", func(c *gin.Context) {
		response.Success(c, "MyPOS Backend API is healthy", gin.H{
			"status":    "running",
			"framework": "gin",
			"offline":   true,
		})
	})

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			_ = auth
			// auth.POST("/login", authHandler.Login)
			// auth.POST("/login", authHandler.Login)
			// auth.GET("/me", authHandler.Me)
		}

		user := api.Group("/users")
		{
			_ = user
			user.POST("/register", userHandler.Register)
			user.GET("/:username", userHandler.GetByUsername)
			user.PUT("/change-password", userHandler.ChangePassword)
		}

		products := api.Group("/products")
		{
			_ = products
			// products.GET("", productHandler.GetAll)
			// products.GET("/:id", productHandler.GetByID)
			// products.POST("", productHandler.Create)
			// products.PUT("/:id", productHandler.Update)
			// products.POST("/:id/restock", productHandler.Restock)
		}

		transactions := api.Group("/transactions")
		{
			_ = transactions
			// transactions.POST("/checkout", transactionHandler.Checkout)
			// transactions.GET("", transactionHandler.GetAll)
			// transactions.GET("/:id", transactionHandler.GetByID)
		}

		// Reports & Analytics Routes (Owner Only)
		reports := api.Group("/reports")
		{
			_ = reports
			// reports.GET("/daily", reportHandler.DailyReport)
			// reports.GET("/monthly", reportHandler.MonthlyReport)
			// reports.GET("/top-selling", reportHandler.TopSelling)
		}

		// Thermal USB Printer Routes
		printer := api.Group("/printer")
		{
			_ = printer
			// printer.POST("/test", printerHandler.TestPrint)
			// printer.POST("/receipt/:invoice", printerHandler.PrintReceipt)
		}
	}

	return r
}
