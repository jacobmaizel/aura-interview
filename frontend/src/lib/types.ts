// type Patient struct {
// 	Id                int               `json:"id"`
// 	Name              string            `json:"name"`
// 	FirstName         string            `json:"first_name"`
// 	Age               int               `json:"age"`
// 	Height            int               `json:"height"`
// 	Weight            int               `json:"weight"`
// 	Gender            string            `json:"gender"`
// 	Medications       []Medication      `json:"medications"`
// 	Body_temperatures []BodyTemperature `json:"body_temperatures"`
// }

// type Medication struct {
// 	Id         int        `json:"id"`
// 	Name       string     `json:"name"`
// 	Dosage     string     `json:"dosage"`
// 	StartDate  *time.Time `json:"start_date"`
// 	EndDate    *time.Time `json:"end_date"`
// 	Patient_id int        `json:"patient_id"`
// }

// type BodyTemperature struct {
// 	Id          int       `json:"id"`
// 	Date        time.Time `json:"date"`
// 	Temperature float64   `json:"temperature"`
// 	Patient_id  int       `json:"patient_id"`
// }

export type Patient = {
  id: number;
  name: string;
  first_name: string;
  age: number;
  height: number;
  weight: number;
  gender: string;
};

export type Medication = {
  id: number;
  name: string;
  dosage: string;
  start_date: string;
  end_date: string;
  patient_id: number;
};

export type BodyTemperature = {
  id: number;
  date: string;
  temperature: number;
  patient_id: number;
};
