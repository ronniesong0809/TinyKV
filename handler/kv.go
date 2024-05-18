package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ronniesong0809/tinyKv/store"
)

func GetValue(c *gin.Context) {
	key := c.Param("key")
	value, expirationTime, err := store.Get(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": value, "ttlLeft": expirationTime.Seconds()})
}

func SetValue(c *gin.Context) {
	key := c.Param("key")
	var json struct {
		Value interface{}   `json:"value" binding:"required"`
		TTL   time.Duration `json:"ttl"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ttlDuration := time.Duration(json.TTL) * time.Second
	if json.TTL == 0 {
		ttlDuration = 1 * time.Minute
	}
	store.Set(key, json.Value, ttlDuration)
	c.JSON(http.StatusOK, gin.H{"status": "set"})
}

func UpdateValue(c *gin.Context) {
	key := c.Param("key")
	var json struct {
		Value interface{}   `json:"value" binding:"required"`
		TTL   time.Duration `json:"ttl"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ttlDuration := time.Duration(json.TTL) * time.Second
	err := store.Update(key, json.Value, ttlDuration)
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
