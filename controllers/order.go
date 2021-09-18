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

	
	// customer_name := c.Param("customerName")
	
	// Call BindJSON to bind the received JSON to
    // newAlbum.
	if err := c.BindJSON(&itemOrder); err != nil {
		fmt.Println(order, "ordd")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	  }
	// order := models.Book{Title: input.Title, Author: input.Author}
	fmt.Println(order.CustomerName, ">>>")

	// item_code := c.PostForm("itemCode")
	// description := c.PostForm("description")
	// quantity := c.PostForm("quantity")
	order.CustomerName = itemOrder.CustomerName
	order.OrderedAt = time.Now()
	item.ItemCode = itemOrder.ItemCode
	item.Description = itemOrder.Description
	item.Quantity = itemOrder.Quantity

	err := db.Create(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		item.OrderId = order.ID
		fmt.Println(order.ID, ">>>> Order Id")
		err = db.Create(&item).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			result = gin.H {
				"orderedAt": order.OrderedAt,
				"CustomerName": order.CustomerName,
				"ItemCode": item.ItemCode,
				"Description": item.Description,
				"Quantity": item.Quantity,
			}
		}
	
	}
	
	c.JSON(http.StatusOK, result)
}

