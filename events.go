package goeloquent

import (
	"fmt"
	"gorm.io/gorm"
)

// BeforeCreate Hook يتم استدعاؤه قبل إنشاء السجل
func BeforeCreate(tx *gorm.DB) {
	fmt.Println("✅ BeforeCreate hook triggered!")
}

// AfterCreate Hook يتم استدعاؤه بعد إنشاء السجل
func AfterCreate(tx *gorm.DB) {
	fmt.Println("✅ AfterCreate hook triggered!")
}
