package api

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/aura/backend/models"
	"github.com/jacobmaizel/aura/backend/server"
)

func AddTemperatureRoutes(rg *gin.RouterGroup) {
	TemperatureRoutes := rg.Group("/temperatures")
	TemperatureRoutes.POST("/", createTemperatureReading)
	TemperatureRoutes.GET("/patient/:id", getTemperatureForPatient)

}

func createTemperatureReading(c *gin.Context) {
	conn, err := server.Pool.Acquire(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": "Error acquiring connection"})
		return
	}

	defer conn.Release()

	var newTemp models.NewBodyTemperature

	if err := c.BindJSON(&newTemp); err != nil {
		c.JSON(400, gin.H{"error": "Error parsing request"})
		return
	}
	var existingTemp bool

	conn.QueryRow(context.Background(), "SELECT EXISTS (SELECT 1 FROM body_temperatures WHERE date = $1 AND patient_id = $2)", newTemp.Date, newTemp.Patient_id).Scan(&existingTemp)

	if existingTemp {
		c.JSON(400, gin.H{"error": "Temperature reading already exists for today"})
		return
	}

	var res models.BodyTemperature

	conn.QueryRow(context.Background(), "INSERT INTO body_temperatures (date, temperature, patient_id) VALUES ($1, $2, $3) RETURNING id, date, temperature, patient_id", newTemp.Date, newTemp.Temperature, newTemp.Patient_id).Scan(&res.Id, &res.Date, &res.Temperature, &res.Patient_id)

	c.JSON(200, res)

}

// get temperature data for a patient, with option to change the scale from 1,3 or 6 months
func getTemperatureForPatient(c *gin.Context) {
	conn, err := server.Pool.Acquire(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": "Error acquiring connection"})
		return
	}

	defer conn.Release()

	var temps []models.BodyTemperature
	var patient_id = c.Param("id")
	var scale = c.DefaultQuery("scale", "1")
	scaleInt, _ := strconv.Atoi(scale)

	startTime := time.Now().AddDate(0, -scaleInt, 0)
	rows, err := conn.Query(context.Background(), "SELECT * FROM body_temperatures WHERE patient_id = $1 AND date > $2", patient_id, startTime)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error querying temperatures"})
		return
	}

	for rows.Next() {
		var t models.BodyTemperature
		err := rows.Scan(&t.Id, &t.Date, &t.Temperature, &t.Patient_id)

		if err != nil {
			c.JSON(500, gin.H{"error": "Error scanning temperatures"})
			return
		}

		temps = append(temps, t)

	}

	c.JSON(200, temps)

}
