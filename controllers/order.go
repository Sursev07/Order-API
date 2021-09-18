package controllers

import (
	"fmt"
	"Order-API/database"
	"Order-API/models"
	"net/http"
	"time"
	"strconv"

	"github.com/gin-gonic/gin"
)


func CreateOrder(c *gin.Context)  {
	var (
		itemOrder models.ItemOrder
		order models.Order
		item models.Item
		result gin.H
	)

	db := database.GetDB()
	
	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&itemOrder); err != nil {
		fmt.Println(order, "ordd")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	  }
	// order := models.Book{Title: input.Title, Author: input.Author}
	fmt.Println(order.CustomerName, ">>>")
	order.CustomerName = itemOrder.CustomerName
	order.OrderedAt = time.Now()
	
	

	err := db.Create(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		item.OrderId = order.ID
		fmt.Println(len(itemOrder.Items), ">>>> Order Id")
		itemslen := len(itemOrder.Items)
		var allItem []models.Item
		for i := 0; i < itemslen; i++ {
			item.ItemCode = itemOrder.Items[i].ItemCode
			item.Description = itemOrder.Items[i].Description
			item.Quantity = itemOrder.Items[i].Quantity

			err = db.Create(&item).Error
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			allItem = append(allItem,item)
		}
		fmt.Println(">>>> Masuk")
		result = gin.H {
			"orderedAt": order.OrderedAt,
			"CustomerName": order.CustomerName,
			"Items": allItem,
		}
		
	
	}
	
	c.JSON(http.StatusOK, result)
}

func GetAllOrders(c *gin.Context) {
	db := database.GetDB()
	Orders := []models.Order{}
	err := db.Debug().Preload("Items").Find(&Orders).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// fmt.Printf("%+v\n", Orders)
	for _, v := range Orders {
		fmt.Println(v.Items, ">>>")
	}
	result := gin.H {
		"data" : Orders,
	}
	 c.JSON(http.StatusOK, result)
}

func DeleteOrder(c *gin.Context) {
	db := database.GetDB()
	Item := models.Item{}
	Order := models.Order{}
	orderId, _ := strconv.Atoi(c.Param("id"))
	// fmt.Println(orderId, "OrderId")
	err := db.First(&Order, uint(orderId)).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     "Not Found",
			"message": "data Order not found",
		})
		return
	}

	err = db.Where("order_id = ?", orderId).Delete(&Item).Error
	err = db.Delete(Order, uint(orderId)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order has been successfully deleted",
	})
}

// func UpdateOrder(c *gin.Context) {
// 	db := database.GetDB()
// 	Item := models.Item{}
// 	Order := models.Order{}
// 	orderId, _ := strconv.Atoi(c.Param("id"))
// 	// fmt.Println(orderId, "OrderId")
// 	err := db.First(&Order, uint(orderId)).Error
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"err":     "Not Found",
// 			"message": "data Order not found",
// 		})
// 		return
// 	}

// 	err = db.Where("order_id = ?", orderId).Delete(&Item).Error
// 	err = db.Delete(Order, uint(orderId)).Error
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"err":     "Bad Request",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Order has been successfully deleted",
// 	})
// }



