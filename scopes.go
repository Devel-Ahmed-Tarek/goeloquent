package goeloquent

import "gorm.io/gorm"

// ScopeActive مثال على Scope لإرجاع السجلات النشطة
func ScopeActive(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", "active")
}
