func SeedData(c *gin.Context) {
	courses := []models.Course{
		{Name: "Curso de Go", Description: "Aprende Go desde cero", Rating: 4.5},
		{Name: "Curso de Docker", Description: "Domina Docker y Kubernetes", Rating: 4.7},
	}
	for _, course := range courses {
		services.CreateCourse(context.Background(), &course)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Datos insertados"})
}

