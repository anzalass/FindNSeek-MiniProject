package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitModel() *gorm.DB {
	dsn := "anzalas:napoli098@tcp(database-1.crr9xsi3mimc.ap-southeast-2.rds.amazonaws.com:3306)/findnseek?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Couldn't conect database")
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Item{}, &Pengajuan{}, &Persetujuan{})
}
