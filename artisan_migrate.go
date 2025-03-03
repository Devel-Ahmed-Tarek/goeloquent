package main

import (
	"log"
	"os"

	"github.com/Devel-Ahmed-Tarek/goeloquent/config"
	"github.com/Devel-Ahmed-Tarek/goeloquent/goeloquent"
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
