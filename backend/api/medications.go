package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/aura/backend/models"
	"github.com/jacobmaizel/aura/backend/server"
)

func AddMedicationRoutes(rg *gin.RouterGroup) {
	MedicationsRoutes := rg.Group("/medications")
	MedicationsRoutes.POST("/", createMedicationForPatient)
	MedicationsRoutes.GET("/patient/:id", getMedicationForPatient)
	MedicationsRoutes.DELETE("/:id", deleteMedicationForPatient)

}

// type NewMedication struct {
// 	Name       string `json:"name"`
// 	Dosage     string `json:"dosage"`
// 	StartDate  string `json:"start_date"`
// 	EndDate    string `json:"end_date"`
// 	Patient_id int    `json:"patient_id"`
// }

// type Medication struct {
// 	Id         int       `json:"id"`
// 	Name       string    `json:"name"`
// 	Dosage     string    `json:"dosage"`
// 	StartDate  time.Time `json:"start_date"`
// 	EndDate    time.Time `json:"end_date"`
// 	Patient_id int       `json:"patient_id"`
// }

func createMedicationForPatient(c *gin.Context) {

	conn, err := server.Pool.Acquire(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": "Error acquiring connection"})
		return
	}

	defer conn.Release()

	var newMed models.NewMedication

	if err := c.BindJSON(&newMed); err != nil {
		c.JSON(400, gin.H{"error": "Error parsing request"})
		return
	}

	var res models.Medication

	if *newMed.StartDate == "" {
		newMed.StartDate = nil
	}

	if *newMed.EndDate == "" {
		newMed.EndDate = nil
	}

	conn.QueryRow(context.Background(), "INSERT INTO medications (name, dosage, start_date, end_date, patient_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, dosage, start_date, end_date, patient_id", newMed.Name, newMed.Dosage, newMed.StartDate, newMed.EndDate, newMed.Patient_id).Scan(&res.Id, &res.Name, &res.Dosage, &res.StartDate, &res.EndDate, &res.Patient_id)

	c.JSON(200, res)

}

func getMedicationForPatient(c *gin.Context) {

	conn, err := server.Pool.Acquire(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": "Error acquiring connection"})
		return
	}

	defer conn.Release()

	id := c.Param("id")

	rows, err := conn.Query(context.Background(), "SELECT * FROM medications WHERE patient_id = $1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error querying medications"})
		return
	}

	var medications []models.Medication
	for rows.Next() {
		var m models.Medication
		err := rows.Scan(&m.Id, &m.Name, &m.Dosage, &m.StartDate, &m.EndDate, &m.Patient_id)

		if err != nil {
			c.JSON(500, gin.H{"error": "Error scanning medications"})
			return
		}

		medications = append(medications, m)

	}

	c.JSON(200, medications)

}

func deleteMedicationForPatient(c *gin.Context) {

	conn, err := server.Pool.Acquire(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": "Error acquiring connection"})
		return
	}

	defer conn.Release()

	id := c.Param("id")

	_, err = conn.Exec(context.Background(), "DELETE FROM medications WHERE id = $1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting medication"})
		return
	}

	c.JSON(200, gin.H{"message": "Medication deleted"})

}
