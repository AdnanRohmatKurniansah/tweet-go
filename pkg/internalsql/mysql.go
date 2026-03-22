package internalsql

import (
    "fmt"
    "log"

    "github.com/AdnanRohmatKurniansah/tweet-go/internal/config"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func ConnectMySQL(cfg *config.Config) (*gorm.DB, error) {
    datasourceName := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
        cfg.DB_USER,
        cfg.DB_PASSWORD,
        cfg.DB_HOST,
        cfg.DB_PORT,
        cfg.DB_NAME,
        "Asia%2FJakarta",
    )

    db, err := gorm.Open(mysql.Open(datasourceName), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("error connecting to database: %w", err)
    }

    log.Println("Database connected")

    return db, nil
}