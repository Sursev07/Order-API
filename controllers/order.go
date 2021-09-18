package controllers

import (
	"fmt"
	"Order-API/database"
	"Order-API/models"
	"net/http"
	"time"

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
			allItem = append(allItem,itemOrder.Items[i])
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

// func GetOrders(c *gin.Context) {
	 
// }

