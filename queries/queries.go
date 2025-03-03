package queries

import (
	"gorm.io/gorm"
	"strings"
)

// QueryWhere يُطبق شرط where باستخدام خريطة
func QueryWhere(db *gorm.DB, conditions map[string]interface{}) *gorm.DB {
	return db.Where(conditions)
}

// QueryWhereIn يُطبق شرط where in على عمود معين
func QueryWhereIn(db *gorm.DB, column string, values []interface{}) *gorm.DB {
	return db.Where(column+" IN ?", values)
}

// QueryGroupBy يُطبق groupBy على الاستعلام
func QueryGroupBy(db *gorm.DB, columns ...string) *gorm.DB {
	groupByClause := strings.Join(columns, ", ")
	return db.Group(groupByClause)
}

// QueryOrderBy يُطبق orderBy على الاستعلام
func QueryOrderBy(db *gorm.DB, column string, order string) *gorm.DB {
	order = strings.ToLower(order)
	if order != "asc" && order != "desc" {
		order = "asc"
	}
	return db.Order(column + " " + order)
}
