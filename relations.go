package goeloquent

import "gorm.io/gorm"

// HasMany يقوم بتحميل علاقة HasMany باستخدام Preload
func HasMany(db *gorm.DB, model interface{}, relation string) *gorm.DB {
	return db.Preload(relation)
}

// BelongsTo يقوم بتحميل علاقة BelongsTo باستخدام Preload
func BelongsTo(db *gorm.DB, model interface{}, relation string) *gorm.DB {
	return db.Preload(relation)
}

// ManyToMany يقوم بتحميل علاقة ManyToMany باستخدام Preload
func ManyToMany(db *gorm.DB, model interface{}, relation string) *gorm.DB {
	return db.Preload(relation)
}
