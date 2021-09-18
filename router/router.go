package router

import(
	"Order-API/controllers"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)
func StartApp() *gin.Engine  {
	router := gin.Default()
	fmt.Println("SEVENNN")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "welcome to Order API by Seven"})    
	  })

	router.GET("/orders", controllers.GetAllOrders)
	router.POST("/orders", controllers.CreateOrder)
	router.PUT("/orders/:id", controllers.UpdateOrder)
	router.DELETE("/orders/:id", controllers.DeleteOrder)
	return router
}