package models

import "time"

type TestResult struct {
	ID              uint      `gorm:"primary_key;auto_increment" json:"id"`
	MedicalRecordID uint      `gorm:"not null;index" json:"medical_record_id"`
	TestName        string    `gorm:"type:text;not null" json:"test_name"`
	Result          string    `gorm:"type:text;not null" json:"result"`
	Unit            string    `gorm:"type:text" json:"unit"`
	ReferenceRange  string    `gorm:"type:text" json:"reference_range"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
