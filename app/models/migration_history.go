package models

import "time"

// MigrationHistory يسجل كل عملية ترحيل تم تطبيقها
type MigrationHistory struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Migration string    `json:"migration" gorm:"uniqueIndex;not null"`
	Batch     int       `json:"batch" gorm:"not null"`
	AppliedAt time.Time `json:"applied_at" gorm:"autoCreateTime"`
}
