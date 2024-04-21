package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jacobmaizel/aura/backend/models"
)

func LoadData() error {

	file, err := os.Open("patient_data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	var patients []models.Patient
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&patients); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err
	}

	for _, p := range patients {

		ctx := context.Background()
		conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

		if err != nil {
			fmt.Println("Error acquiring connection:", err)
			return err
		}

		var dbPatientId int

		conn.QueryRow(ctx, "INSERT INTO patients (name, first_name, age, height, weight, gender) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", p.Name, p.FirstName, p.Age, p.Height, p.Weight, p.Gender).Scan(&dbPatientId)

		for _, m := range p.Medications {
			_, err = conn.Exec(ctx, "INSERT INTO medications (name, dosage, start_date, end_date, patient_id) VALUES ($1, $2, $3, $4, $5)", m.Name, m.Dosage, m.StartDate, m.EndDate, dbPatientId)
			if err != nil {
				fmt.Println("Error inserting medication:", err)
				return err
			}

		}

		for _, b := range p.Body_temperatures {
			_, err = conn.Exec(ctx, "INSERT INTO body_temperatures (date, temperature, patient_id) VALUES ($1, $2, $3)", b.Date, b.Temperature, dbPatientId)
			if err != nil {
				fmt.Println("Error inserting body temperature:", err)
				return err
			}
		}

	}
	return nil
}
