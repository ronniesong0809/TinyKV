package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronniesong0809/tinyKv/store"
)

func GetValue(c *gin.Context) {
	key := c.Param("key")
	value, err := store.Get(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": value})
}

func SetValue(c *gin.Context) {
	key := c.Param("key")
	var json struct {
		Value string `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	store.Set(key, json.Value)
	c.JSON(http.StatusOK, gin.H{"status": "set"})
}

func UpdateValue(c *gin.Context) {
	key := c.Param("key")
	var json struct {
		Value string `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := store.Update(key, json.Value)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func DeleteValue(c *gin.Context) {
	key := c.Param("key")
	err := store.Delete(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
