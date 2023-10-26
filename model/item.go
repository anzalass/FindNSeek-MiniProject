package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID        string `json:"id" form:"id`
	Id_User   string `json:"id_user" form:"id_user"`
	Judul     string `json:"judul" form:"judul"`
	Kategori  string `json:"kategori" form:"kategori"`
	Tanggal   string `json:"tanggal" form:"tanggal"`
	Lokasi    string `json:"lokasi" form:"lokasi"`
	Foto      string `json:"foto" form:"foto"`
	Alamat    string `json:"alamat" form:"alamat"`
	NoHp      string `json:"no_hp" form:"no_hp"`
	Email     string `json:"email" form:"email"`
	Status    int    `json:"status" form:"status"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
}

type UpdateStatusItem struct {
	Id_Item string `json:"id_item" form:"id_item"`
	Status  int    `json:"status" form:"status"`
}

type ItemInterface interface {
	CreateItem(data Item) (*Item, error)
	UpdateStatusItem(id string) error
	GetItemsWithPaginationAndSearch(judul string, kategori string, page int, perPage int) []Item
	GetItemsByID(id string) (*Item, error)
	GetPersetujuanByID(id string) (*Persetujuan, error)
	GetPengajuanByItemId(id string) ([]Pengajuan, error)
	UpdateItemsById(data Item) (*Item, error)
	DeleteItemsById(id string) error
}

type ItemModel struct {
	db *gorm.DB
}

func (im *ItemModel) Init(db *gorm.DB) {
	im.db = db
}

func NewItemModel(db *gorm.DB) ItemInterface {
	return &ItemModel{
		db: db,
	}
}

func (im *ItemModel) CreateItem(data Item) (*Item, error) {
	if err := im.db.Create(&data).Error; err != nil {
		logrus.Error("model : error register user")
		return nil, err
	}

	return &data, nil
}

func (im *ItemModel) UpdateStatusItem(id string) error {
	if err := im.db.Model(&Item{}).Where("id = ?", id).Update("status", 1).Error; err != nil {
		logrus.Error("model : error update")
		return nil
	}

	return nil
}

func (im *ItemModel) GetItemsWithPaginationAndSearch(judul string, kategori string, page int, perPage int) []Item {
	var listItem = []Item{}

	if judul != "" || kategori != "" {
		if err := im.db.Where("judul LIKE ? OR  kategori LIKE ?", "%"+judul+"%", "%"+kategori+"%").Offset((page - 1) * perPage).Limit(perPage).Find(&listItem).Error; err != nil {
			logrus.Error("model : error get item", err.Error())
			return nil
		}
	}

	if judul != "" || kategori == "" {
		if err := im.db.Where("judul LIKE ?", "%"+judul+"%").Offset((page - 1) * perPage).Limit(perPage).Find(&listItem).Error; err != nil {
			logrus.Error("model : error get item", err.Error())
			return nil
		}
	}

	if kategori != "" || judul == "" {
		if err := im.db.Where("kategori LIKE ?", "%"+kategori+"%").Offset((page - 1) * perPage).Limit(perPage).Find(&listItem).Error; err != nil {
			logrus.Error("model : error get item", err.Error())
			return nil
		}
	}

	if judul == "" && kategori == "" {
		if err := im.db.Offset((page - 1) * perPage).Limit(perPage).Find(&listItem).Error; err != nil {
			logrus.Error("model : error get item", err.Error())
			return nil
		}
	}

	return listItem

}

func (im *ItemModel) GetItemsByID(id string) (*Item, error) {
	var item = Item{}
	if err := im.db.Where("id = ?", id).First(&item).Error; err != nil {
		logrus.Error("model : error getting items by id: ", err)
		return nil, err
	}

	return &item, nil

}

func (im *ItemModel) GetPengajuanByItemId(id string) ([]Pengajuan, error) {
	var listPengajuan = []Pengajuan{}

	if err := im.db.Where("id_item = ?", id).Find(&listPengajuan).Error; err != nil {
		logrus.Error("model : error get pengajuan by items id", err)
		return nil, err
	}

	return listPengajuan, nil

}

func (im *ItemModel) GetPersetujuanByID(id string) (*Persetujuan, error) {
	var persetujuan = Persetujuan{}

	if err := im.db.Where("id_item = ?", id).Find(&persetujuan).Error; err != nil {
		logrus.Error("model : error get persetujuan by items id", err)
		return nil, err
	}

	return &persetujuan, nil

}

func (im *ItemModel) UpdateItemsById(data Item) (*Item, error) {
	var item = Item{}

	if err := im.db.Model(&item).Where("id = ?", data.ID).Updates(map[string]interface{}{"judul": data.Judul, "kategori": data.Kategori, "tanggal": data.Tanggal, "lokasi": data.Lokasi, "foto": data.Foto, "alamat": data.Alamat, "deskripsi": data.Deskripsi, "email": data.Email, "no_hp": data.NoHp}).Error; err != nil {
		logrus.Error("model : error updating item : ", err)
		return &item, nil
	}

	return &item, nil
}

func (im *ItemModel) DeleteItemsById(id string) error {
	item := Item{}
	if err := im.db.Where("id=?", id).Delete(&item).Error; err != nil {
		logrus.Error("model : error deleting item : ", err.Error())
		return nil
	}

	return nil
}
