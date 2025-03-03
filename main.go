package main

import (
	"log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/username/goeloquent/config"      // تأكد من تعديل المسار حسب مشروعك
	"github.com/username/goeloquent/goeloquent"    // للوصول إلى وظائف المكتبة
	"github.com/username/goeloquent/app/models"    // للوصول إلى النماذج
	"github.com/username/goeloquent/routes"        // تعريف Routes
)

func main() {
	// تحميل إعدادات المكتبة من .env
	config.LoadConfig()

	// الاتصال بقاعدة البيانات
	config.Connect(config.GlobalConfig.DB_DSN)
	goeloquent.DB = config.DB

	// تنفيذ الميجريشنات مع تسجيلها
	goeloquent.MigrateAllWithHistory()

	// الاتصال بـ Redis
	goeloquent.ConnectRedis(config.GlobalConfig.RedisAddr)

	// تجربة إنشاء سجل User
	user := models.User{
		Name:   "أحمد",
		Email:  "ahmed@example.com",
		Status: "active",
	}
	if err := goeloquent.DB.Create(&user).Error; err != nil {
		log.Printf("❌ Error creating user: %v", err)
	} else {
		log.Printf("✅ User created: %+v", user)
	}

	// تجربة إنشاء سجل Post مرتبط بالمستخدم
	post := models.Post{
		Title:  "عنوان المشاركة",
		Body:   "محتوى المشاركة التجريبي",
		UserID: user.ID,
	}
	if err := goeloquent.DB.Create(&post).Error; err != nil {
		log.Printf("❌ Error creating post: %v", err)
	} else {
		log.Printf("✅ Post created: %+v", post)
	}

	// تجربة Pagination لجلب المستخدمين
	var usersList []models.User
	paginationResult, err := goeloquent.Paginate(goeloquent.DB, &models.User{}, &usersList, "1", "5")
	if err != nil {
		log.Printf("❌ Pagination error: %v", err)
	} else {
		fmt.Printf("Paginated Users: %+v\n", paginationResult)
	}

	// تجربة ScopeActive
	var activeUsers []models.User
	if err := goeloquent.ScopeActive(goeloquent.DB).Find(&activeUsers).Error; err != nil {
		log.Printf("❌ Error fetching active users: %v", err)
	} else {
		fmt.Printf("Active Users: %+v\n", activeUsers)
	}

	// تجربة إرسال بريد إلكتروني
	emailConfig := goeloquent.EmailConfig{
		SMTPHost:    config.GlobalConfig.SMTPHost,
		SMTPPort:    config.GlobalConfig.SMTPPort,
		Username:    config.GlobalConfig.SMTPUser,
		Password:    config.GlobalConfig.SMTPPass,
		FromAddress: config.GlobalConfig.FromAddress,
		FromName:    config.GlobalConfig.FromName,
	}
	emailService := goeloquent.NewEmailService(emailConfig)
	if err := emailService.SendEmail("recipient@example.com", "تجربة إرسال بريد من GoEloquent", "<h1>مرحباً!</h1><p>هذه رسالة اختبارية من مكتبة GoEloquent.</p>"); err != nil {
		log.Printf("❌ Email sending error: %v", err)
	}

	// تجربة التخزين: حفظ واسترجاع ملف
	storageBase := "storage"
	relativePath := "uploads/images"
	// يمكنك استبدال هذا المثال بفتح ملف حقيقي باستخدام os.Open
	// هنا مثال باستخدام بافر تجريبي لنص
	fromContent := "هذا محتوى الملف التجريبي."
	fileReader := bytes.NewBufferString(fromContent)
	paths, err := goeloquent.SaveMediaFile(storageBase, relativePath, fileReader, "sample.jpg")
	if err != nil {
		log.Printf("❌ Error saving media file: %v", err)
	} else {
		fmt.Printf("Media file paths: %+v\n", paths)
	}

	// تجربة استرجاع نسخة من الملف (مثلاً النسخة المصغرة)
	retrievedData, err := goeloquent.GetMediaFileVersion(storageBase, relativePath, filepath.Base(paths["original"]), "thumbnail")
	if err != nil {
		log.Printf("❌ Error retrieving media file version: %v", err)
	} else {
		fmt.Printf("Retrieved media file version size: %d bytes\n", len(retrievedData))
	}

	// تشغيل السيرفر باستخدام Gin
	router := gin.Default()
	routes.SetupRoutes(router)
	log.Printf("🚀 %s Server running on port %s", config.GlobalConfig.AppName, config.GlobalConfig.Port)
	router.Run(":" + config.GlobalConfig.Port)
}
