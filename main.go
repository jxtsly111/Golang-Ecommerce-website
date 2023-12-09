package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jxtsly111/ecommerce-yt/controllers"
	"github.com/jxtsly111/ecommerce-yt/database"
	"github.com/jxtsly111/ecommerce-yt/middleware"
	"github.com/jxtsly111/ecommerce-yt/routes"
)

func main()  {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    app := controllers.NewApplication(database.ProductData(database.Client, "Products"),database.UserData(database.Client, "Users"))

    router :=gin.New()
    router.Use(gin.Logger())

    routes.UserRoutes(router)
    router.Use(middleware.Authentication())

    router.GET("/addtocart", app.AddToCart())
    router.GET("/removeitem", app.Removeitem())
    router.GET("/cartcheckout", app.BuyFromCart())
    router.GET("/instantbuy", app.InstantBuy())

    log.Fatal(router.Run(":" + port))

}