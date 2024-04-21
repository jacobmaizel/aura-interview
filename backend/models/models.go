package models

import "time"

type Patient struct {
	Id                int               `json:"id"`
	Name              string            `json:"name"`
	FirstName         string            `json:"first_name"`
	Age               int               `json:"age"`
	Height            int               `json:"height"`
	Weight            int               `json:"weight"`
	Gender            string            `json:"gender"`
	Medications       []JsonMedication  `json:"medications"`
	Body_temperatures []JsonTemperature `json:"body_temperatures"`
}

type Medication struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Dosage     string     `json:"dosage"`
	StartDate  *time.Time `json:"start_date"`
	EndDate    *time.Time `json:"end_date"`
	Patient_id int        `json:"patient_id"`
}

type BodyTemperature struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Temperature float64   `json:"temperature"`
	Patient_id  int       `json:"patient_id"`
}

type JsonMedication struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Dosage     string  `json:"dosage"`
	StartDate  *string `json:"start_date"`
	EndDate    *string `json:"end_date"`
	Patient_id int     `json:"patient_id"`
}

type JsonTemperature struct {
	Id          int     `json:"id"`
	Date        *string `json:"date"`
	Temperature float64 `json:"temperature"`
	Patient_id  int     `json:"patient_id"`
}

type NewPatient struct {
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	Gender    string `json:"gender"`
}

type NewMedication struct {
	Name       string  `json:"name"`
	Dosage     string  `json:"dosage"`
	StartDate  *string `json:"start_date"`
	EndDate    *string `json:"end_date"`
	Patient_id int     `json:"patient_id"`
}

type NewBodyTemperature struct {
	Date        string  `json:"date"`
	Temperature float64 `json:"temperature"`
	Patient_id  int     `json:"patient_id"`
}
