package config

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// AppConfig menyimpan konfigurasi global aplikasi.
type AppConfig struct {
	AppName   string
	AppPort   string
	AppEnv    string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
	JWTTTL    time.Duration
	SMTPHost  string
	SMTPPort  int
	SMTPUser  string
	SMTPPass  string
	SMTPFrom  string
}

var (
	cfgOnce sync.Once
	appCfg  *AppConfig
	cfgErr  error

	// DB adalah koneksi global yang dipakai oleh seluruh modul.
	DB *gorm.DB
)

// Load membaca variabel environment dan menyimpannya di AppConfig.
func Load() (*AppConfig, error) {
	cfgOnce.Do(func() {
		_ = godotenv.Load()

		ttl := parseDuration(getEnv("JWT_TTL", "12h"))
		smtpPort := parseInt(getEnv("SMTP_PORT", "587"))

		appCfg = &AppConfig{
			AppName:   getEnv("APP_NAME", "AMK BE"),
			AppPort:   getEnv("APP_PORT", "3000"),
			AppEnv:    getEnv("APP_ENV", "development"),
			DBHost:    normalizeDBHost(getEnv("DB_HOST", "127.0.0.1")),
			DBPort:    getEnv("DB_PORT", "3306"),
			DBUser:    getEnv("DB_USER", "root"),
			DBPass:    getEnv("DB_PASS", ""),
			DBName:    getEnv("DB_NAME", "amk_db"),
			JWTSecret: getEnv("JWT_SECRET", "secret"),
			JWTTTL:    ttl,
			SMTPHost:  getEnv("SMTP_HOST", ""),
			SMTPPort:  smtpPort,
			SMTPUser:  getEnv("SMTP_USER", ""),
			SMTPPass:  getEnv("SMTP_PASS", ""),
			SMTPFrom:  getEnv("SMTP_FROM", "no-reply@example.com"),
		}
	})

	return appCfg, cfgErr
}

// MustLoad sama seperti Load namun panic bila gagal.
func MustLoad() *AppConfig {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}

	return cfg
}

// InitDB membuka koneksi MySQL menggunakan konfigurasi dari .env.
func InitDB() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	cfg, err := Load()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	DB = db
	return DB, nil
}

// Get mengembalikan konfigurasi yang sudah dimuat.
func Get() *AppConfig {
	cfg, _ := Load()
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseDuration(value string) time.Duration {
	d, err := time.ParseDuration(value)
	if err != nil {
		return 12 * time.Hour
	}
	return d
}

func parseInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return i
}

func normalizeDBHost(value string) string {
	host := strings.TrimSpace(value)
	if host == "" {
		return "127.0.0.1"
	}
	if net.ParseIP(host) != nil || host == "localhost" {
		return host
	}
	if _, err := net.LookupHost(host); err != nil {
		log.Printf("config: DB_HOST=%q tidak bisa di-resolve, fallback ke 127.0.0.1", host)
		return "127.0.0.1"
	}
	return host
}
