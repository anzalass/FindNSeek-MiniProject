package model

import (
	"findnseek/config"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitModel(c config.ProgramConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	// dsn := "anzalas:napoli098@tcp(database-1.crr9xsi3mimc.ap-southeast-2.rds.amazonaws.com:3306)/findnseek?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Couldn't conect database")
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Item{}, &Pengajuan{}, &Persetujuan{})
}
