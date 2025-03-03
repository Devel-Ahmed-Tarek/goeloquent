package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LibraryConfig يحتوي على جميع إعدادات المكتبة
type LibraryConfig struct {
	DB_DSN      string
	RedisAddr   string
	Port        string
	SMTPHost    string
	SMTPPort    int
	SMTPUser    string
	SMTPPass    string
	FromAddress string
	FromName    string
	AppName     string
}

var GlobalConfig LibraryConfig

// LoadConfig يقوم بتحميل المتغيرات من ملف .env وتعيينها في GlobalConfig
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("❌ لم يتم تحميل ملف .env، سيتم استخدام المتغيرات المعرفة في البيئة")
	}
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		smtpPort = 587
	}
	GlobalConfig = LibraryConfig{
		DB_DSN:      os.Getenv("DB_DSN"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		Port:        os.Getenv("PORT"),
		SMTPHost:    os.Getenv("SMTP_HOST"),
		SMTPPort:    smtpPort,
		SMTPUser:    os.Getenv("SMTP_USER"),
		SMTPPass:    os.Getenv("SMTP_PASS"),
		FromAddress: os.Getenv("FROM_ADDRESS"),
		FromName:    os.Getenv("FROM_NAME"),
		AppName:     os.Getenv("APP_NAME"),
	}
	log.Println("✅ Loaded Library Config:", GlobalConfig)
}
