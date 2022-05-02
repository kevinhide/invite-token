package db

import (
	"database/sql"
	"log"
	"time"

	"invite-token/src/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	URL               string
	MaxDBConn         int
	MigrationFilePath string
}

//Init the connection to DB
func Init(config *Config) error {
	if DB == nil {
		sqlDB, err := sql.Open("postgres", config.URL)
		if err != nil {
			log.Println("Unable to open postges connection. Err:", err)
			return err
		}

		sqlDB.SetMaxIdleConns(config.MaxDBConn)
		sqlDB.SetMaxOpenConns(config.MaxDBConn)
		sqlDB.SetConnMaxLifetime(time.Hour)

		DB, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			log.Println("Unable to open postges gorm connection. Err:", err)
			return err
		}

		log.Println("Successfully established database connection to ", config.URL)
	}

	Migrate()

	return nil
}

func Migrate() {
	DB.AutoMigrate(&models.InviteToken{})
}
