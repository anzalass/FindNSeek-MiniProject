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
	NoHp      string `json:"no_hp" form:"no_hp"`
	Email     string `json:"email" form:"email"`
	Alamat    string `json:"alamat" form:"alamat"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
}

type PengajuanInterface interface {
	CreatePengajuan(data Pengajuan) (*Pengajuan, error)
	GetPengajuanByItemId(itemId string) ([]Result, error)
	GetPengajuanById(id string) (*Pengajuan, error)
	CekStatusItemFromPengajuan(id_item string) (*Item, error)
	GetUserNameForSend(id_user string) (*User, error)
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

func (pm *PengajuanModel) CreatePengajuan(data Pengajuan) (*Pengajuan, error) {

	if err := pm.db.Create(&data).Error; err != nil {
		logrus.Error("model : error create pengajuan")
		return nil, err
	}

	return &data, nil
}

func (pm *PengajuanModel) CekStatusItemFromPengajuan(id string) (*Item, error) {
	var item = Item{}
	// var db = InitModel()
	if err := pm.db.First(&item, "id = ?", id).Error; err != nil {
		logrus.Error("model : cek status gagal")
		return nil, err
	}

	return &item, nil
}

func (pm *PengajuanModel) GetUserNameForSend(id string) (*User, error) {
	var user = User{}
	// var db = InitModel()
	if err := pm.db.First(&user, "id = ?", id).Error; err != nil {
		logrus.Error("model :Error getting user")
		return nil, err
	}
	return &user, nil
}

type Result struct {
	ItemsID         string `json:"items_id"`
	ItemsJudul      string `json:"items_judul"`
	PengajuanIDUser string `json:"pengajuan_id_user"`
	PengajuanJudul  string `json:"pengajuan_judul"`
}

func (pm *PengajuanModel) GetPengajuanByItemId(itemId string) ([]Result, error) {
	var result []Result

	query := pm.db.Model(&Item{}).Select("items.id AS items_id , items.judul AS items_judul, pengajuans.id_user AS pengajuan_id_user, pengajuans.judul AS pengajuan_judul").Joins("left join pengajuans on pengajuans.id_item = items.id").Where("items.id = ? AND pengajuans.id_item = ?", itemId, itemId).Scan(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil

}

func (pm *PengajuanModel) GetPengajuanById(id string) (*Pengajuan, error) {
	var pengajuan = Pengajuan{}

	if err := pm.db.First(&pengajuan, "id = ?", id).Error; err != nil {
		logrus.Error("model : gagal mendapatkan data", err.Error())
		return nil, err
	}

	return &pengajuan, nil

}
