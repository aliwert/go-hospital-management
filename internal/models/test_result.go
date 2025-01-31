package models

import "time"

type TestResult struct {
	ID              uint      `json:"id"`
	MedicalRecordID uint      `json:"medical_record_id"`
	TestName        string    `json:"test_name"`
	Result          string    `json:"result"`
	Unit            string    `json:"unit"`
	ReferenceRange  string    `json:"reference_range"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
