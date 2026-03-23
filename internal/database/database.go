package database

import (
    "log"
    "os"
    "time"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := os.Getenv("DB_USER") + ":" +
        os.Getenv("DB_PASSWORD") +
        "@tcp(" + os.Getenv("DB_HOST") +
        ":" + os.Getenv("DB_PORT") +
        ")/" + os.Getenv("DB_NAME") +
        "?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta" 

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Warn),
    })
    if err != nil {
        log.Fatal("Failed to connect database:", err)
    }

    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)

    if err = sqlDB.Ping(); err != nil {
        log.Fatal("Database not reachable:", err)
    }

    DB = db
    log.Println("Database connected successfully")
}

func RunMigrations() {
    dsn := "mysql://" +
        os.Getenv("DB_USER") + ":" +
        os.Getenv("DB_PASSWORD") +
        "@tcp(" + os.Getenv("DB_HOST") +
        ":" + os.Getenv("DB_PORT") +
        ")/" + os.Getenv("DB_NAME") +
        "?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta"

    m, err := migrate.New("file://migrations", dsn)
    if err != nil {
        log.Fatal("Migration init failed:", err)
    }

    if err = m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal("Migration failed:", err)
    }

    log.Println("Migrations applied successfully")
}