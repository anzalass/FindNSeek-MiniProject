package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Pengajuan struct {
	gorm.Model
	ID        string `json:"id" form:"id`
	Id_User   string `json:"id_user" form:"id_user"`
	Judul     string `json:"judul" form:"judul"`
	Kategori  string `json:"kategori" form:"kategori"`
	Id_Item   string `json:"id_item" form:"id_item"`
	Tanggal   string `json:"tanggal" form:"tanggal"`
	Lokasi    string `json:"lokasi" form:"lokasi"`
	Foto      string `json:"foto" form:"foto"`
	Alamat    string `json:"alamat" form:"alamat"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
}

type PengajuanInterface interface {
	CreatePengajuan(data Pengajuan) *Pengajuan
}

type PengajuanModel struct {
	db *gorm.DB
}

func (pm *PengajuanModel) Init(db *gorm.DB) {
	pm.db = db
}

func NewPengajuanModel(db *gorm.DB) PengajuanInterface {
	return &PengajuanModel{
		db: db,
	}
}

func (pm *PengajuanModel) CreatePengajuan(data Pengajuan) *Pengajuan {
	if err := pm.db.Create(&data).Error; err != nil {
		logrus.Error("model : error create pengajuan")
	}

	return &data
}
