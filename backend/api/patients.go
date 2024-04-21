package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/aura/backend/models"
	"github.com/jacobmaizel/aura/backend/server"
)

func AddPatientRoutes(rg *gin.RouterGroup) {
	PatientRoutes := rg.Group("/patients")
	PatientRoutes.GET("/", getPatients)
	PatientRoutes.GET("/:id", getPatient)

}

// CREATE TABLE IF NOT EXISTS patients(
//   id SERIAL PRIMARY KEY,
//   name VARCHAR(50) NOT NULL,
//   first_name VARCHAR(50) NOT NULL,
//   age INTEGER,
//   height INTEGER,
//   weight INTEGER,
//   gender VARCHAR(20) NOT NULL
// );

func getPatients(c *gin.Context) {
	conn, err := server.Pool.Acquire(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": "Error acquiring connection"})
		return
	}

	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT * FROM patients")
	if err != nil {
		c.JSON(500, gin.H{"error": "Error querying patients"})
		return
	}

	var patients []models.Patient
	for rows.Next() {
		var p models.Patient
		err := rows.Scan(&p.Id, &p.Name, &p.FirstName, &p.Age, &p.Height, &p.Weight, &p.Gender)

		if err != nil {
			c.JSON(500, gin.H{"error": "Error scanning patients"})
			return
		}

		patients = append(patients, p)

	}

	c.JSON(200, patients)

}

func getPatient(c *gin.Context) {
	conn, err := server.Pool.Acquire(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": "Error acquiring connection"})
		return
	}

	defer conn.Release()

	id := c.Param("id")

	var p models.Patient
	err = conn.QueryRow(context.Background(), "SELECT * FROM patients WHERE id = $1", id).Scan(&p.Id, &p.Name, &p.FirstName, &p.Age, &p.Height, &p.Weight, &p.Gender)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error scanning patient"})
		return
	}

	c.JSON(200, p)

}
