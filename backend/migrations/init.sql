CREATE TABLE IF NOT EXISTS patients(
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  age INTEGER,
  height INTEGER,
  weight INTEGER,
  gender VARCHAR(20) NOT NULL 
);

CREATE TABLE IF NOT EXISTS medications(
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  dosage VARCHAR(50) NOT NULL,
  start_date DATE,
  end_date DATE,
  patient_id INTEGER REFERENCES patients(id)
);

CREATE TABLE IF NOT EXISTS body_temperatures(
  id SERIAL PRIMARY KEY,
  date DATE,
  temperature numeric,
  patient_id INTEGER REFERENCES patients(id)
);
