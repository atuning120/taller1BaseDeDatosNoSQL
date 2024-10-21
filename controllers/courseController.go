package controllers

import (
	"backend/models"
	"backend/services"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
	courses, err := services.GetAllCourses(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateCourse(context.Background(), &course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, course)
}
