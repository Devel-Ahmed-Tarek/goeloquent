package main

import (
	"log"
	"os"

	"github.com/username/goeloquent/config"     // تأكد من تعديل المسار حسب مشروعك
	"github.com/username/goeloquent/goeloquent"   // للوصول إلى MigrateAllWithHistory
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: go run artisan_migrate.go migrate")
		return
	}

	config.LoadConfig()
	config.Connect(config.GlobalConfig.DB_DSN)
	goeloquent.DB = config.DB

	switch os.Args[1] {
	case "migrate":
		goeloquent.MigrateAllWithHistory()
		log.Println("✅ All migrations applied successfully!")
	default:
		log.Println("Unknown command. Usage: go run artisan_migrate.go migrate")
	}
}
