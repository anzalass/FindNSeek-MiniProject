package model

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Persetujuan struct {
	gorm.Model
	ID           string `json:"id" form:"id"`
	Id_User      string `json:"id_user" form:"id_user"`
	Id_Pengajuan string `json:"id_pengajuan" form:"id_pengajuan"`
	Id_Item      string `json:"id_item" form:"id_item"`
}

type PersetujuanInterface interface {
	CreatePersetujuan(data Persetujuan) (*Persetujuan, error)
	GetPersetujuanByID(id_item string) ([]Res, error)
}

type PersetujuanModel struct {
	db *gorm.DB
}

func (pst *PersetujuanModel) Init(db *gorm.DB) {
	pst.db = db
}

func NewPersetujuanModel(db *gorm.DB) PersetujuanInterface {
	return &PersetujuanModel{
		db: db,
	}
}

func (pst *PersetujuanModel) CreatePersetujuan(data Persetujuan) (*Persetujuan, error) {

	if err := pst.db.Create(&data).Error; err != nil {
		logrus.Error("model : error create perrsetujuan")
		return nil, err

	}

	return &data, nil
}

type Res struct {
	IdItem        string
	IdPersetujuan string
	IdPengajuan   string
	IdUser        string
}

func (pst *PersetujuanModel) GetPersetujuanByID(id_item string) ([]Res, error) {
	var ress []Res

	var query = pst.db.Raw("SELECT items.id AS IdItem, persetujuans.id AS IdPersetujuan, persetujuans.id_pengajuan AS id_pengajuan, persetujuans.id_user AS IdUser FROM `items` JOIN persetujuans ON persetujuans.id_item = items.id WHERE items.id = ?", id_item).Scan(&ress)
	if query.Error != nil {
		return nil, query.Error
	}
	fmt.Println(ress)
	return ress, nil

}
