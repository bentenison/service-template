package conf

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	BookAPIPort     string
	AuthAPIPort     string
	DebugPort       string
	DBDSN           string
	DBName          string
	User            string
	Password        string
	Host            string
	Environment     string
	MaxIdleConns    int
	MaxOpenConns    int
	ShutdownTimeout int
	ResponseTimeOut int
	RequestTimeOut  int
}

// LoadConfig loads configuration from environment variables and optional .env file
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables only")
	}

	config := &Config{
		BookAPIPort: getEnv("BOOKAPI_PORT", "8000"),
		AuthAPIPort: getEnv("AUTHAPI_PORT", "8001"),
		DebugPort:   getEnv("DEBUG_PORT", ":8002"),
		DBDSN:       getEnv("DBDSN", "postgres://user:password@localhost:5432"),
		User:        getEnv("DBUSER", "postgres"),
		Password:    getEnv("DBPASSWORD", "root"),
		Host:        getEnv("HOST", "localhost"),
		DBName:      getEnv("DBNAME", "library"),
		Environment: getEnv("ENV", "devlopment"),
	}
	idleConns, _ := strconv.Atoi(getEnv("MAXIDLECONNS", "10"))
	openConns, _ := strconv.Atoi(getEnv("MAXOPENCONNS", "10"))
	shutdownTimeout, _ := strconv.Atoi(getEnv("SHUTDOWN_TIMEOUT", "5"))
	requestTimeout, _ := strconv.Atoi(getEnv("REQUESTTIMEOUT_SEC", "1"))
	responseTimeout, _ := strconv.Atoi(getEnv("RESPONSETIMEOUT_SEC", "1"))

	config.MaxIdleConns = idleConns
	config.MaxOpenConns = openConns
	config.ResponseTimeOut = responseTimeout
	config.RequestTimeOut = requestTimeout
	config.ShutdownTimeout = shutdownTimeout
	return config, nil
}

// Helper function to get environment variables with a fallback default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
