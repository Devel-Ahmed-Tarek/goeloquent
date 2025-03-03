package goeloquent

import (
	"gorm.io/gorm"
	"strconv"
)

// Paginate ينفذ الباجيناشن مع دعم page و pageSize كسلاسل نصية
func Paginate(db *gorm.DB, model interface{}, results interface{}, pageStr string, pageSizeStr string) (map[string]interface{}, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	var total int64
	db.Model(model).Count(&total)

	offset := (page - 1) * pageSize
	err = db.Limit(pageSize).Offset(offset).Find(results).Error
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"data":        results,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
	}
	return response, nil
}
