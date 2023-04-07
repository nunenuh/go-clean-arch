package database

import (
	"fmt"

	"github.com/zeimedee/go-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"internal/app/adapter/config"
)

type DBConn struct {
	Db *gorm.DB
}

type DBConnInfo struct {
	Info       config.DBConnInfo
	SilentMode bool
}

func (conn *DBConnInfo) Connection() (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN: fmt.Sprintf(
				"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
				conn.Info.User,
				conn.Info.Pass,
				conn.Info.DBName,
				conn.Info.Host,
				conn.Info.Port,
			),
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{},
	)

	if err != nil {
		return db, err
	}

	db.Logger = db.Logger.LogMode(logger.Silent)
	if !conn.SilentMode {
		db.Logger = db.Logger.LogMode(logger.Info)
	}

	dbGorm, err := db.DB()
	if err != nil {
		return db, err
	}

	dbGorm.SetMaxOpenConns(conn.Info.MaxOpenConn)
	dbGorm.SetMaxIdleConns(conn.Info.MaxIdleConn)
	dbGorm.SetConnMaxLifetime(conn.Info.MaxLifetime)

	// migrateDDL(db)

	return db, err
}

func migrateDDL(db *gorm.DB) {
	db.AutoMigrate(&models.Book{})
}
