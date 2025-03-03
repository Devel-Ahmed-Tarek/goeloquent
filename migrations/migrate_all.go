package goeloquent

import (
	"log"
	"time"

	"https://github.com/Devel-Ahmed-Tarek/goeloquent/tree/main/app/models" // تأكد من تعديل المسار حسب مشروعك
)

// MigrateAllWithHistory يقوم بترحيل جميع الموديلات مع تسجيلها في جدول MigrationHistory
func MigrateAllWithHistory() {
	// ترحيل جدول MigrationHistory أولاً
	err := DB.AutoMigrate(&models.MigrationHistory{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate MigrationHistory table: %v", err)
	}

	// قائمة الميجريشنات التي نريد تطبيقها
	migrations := []struct {
		Name  string
		Model interface{}
	}{
		{"users", &models.User{}},
		{"posts", &models.Post{}},
		// أضف موديلات أخرى هنا إذا لزم الأمر
	}

	for _, m := range migrations {
		var count int64
		// التحقق من وجود السجل في جدول MigrationHistory
		if err := DB.Model(&models.MigrationHistory{}).Where("migration = ?", m.Name).Count(&count).Error; err != nil {
			log.Fatalf("❌ Failed to check migration '%s': %v", m.Name, err)
		}
		if count > 0 {
			log.Printf("ℹ️ Migration '%s' already applied, skipping.", m.Name)
			continue
		}

		// تنفيذ ترحيل الموديل
		if err := DB.AutoMigrate(m.Model); err != nil {
			log.Fatalf("❌ Migration for '%s' failed: %v", m.Name, err)
		}

		// تسجيل عملية الترحيل
		migrationRecord := models.MigrationHistory{
			Migration: m.Name,
			Batch:     1,
			AppliedAt: time.Now(),
		}
		if err := DB.Create(&migrationRecord).Error; err != nil {
			log.Fatalf("❌ Failed to record migration '%s': %v", m.Name, err)
		}
		log.Printf("✅ Migration '%s' applied successfully.", m.Name)
	}
}
