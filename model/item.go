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
	Role      int    `json:"role" form:"role"`
	Status    int    `json:"status" form:"status"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
}

type ItemInterface interface {
	CreateItem(data Item) *Item
	GetItemsWithPaginationAndSearch(judul string, page int, perPage int) []Item
	// Login(data User) (*User, error)
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

func (im *ItemModel) CreateItem(data Item) *Item {
	if err := im.db.Create(&data).Error; err != nil {
		logrus.Error("model : error register user")
	}

	return &data
}

func (im *ItemModel) GetItemsWithPaginationAndSearch(judul string, page int, perPage int) []Item {
	var listItem = []Item{}

	if judul != "" {
		if err := im.db.Where("judul LIKE ?", "%"+judul+"%").Offset((page - 1) * perPage).Limit(perPage).Find(&listItem).Error; err != nil {
			logrus.Error("model : error get item", err.Error())
			return nil
		}
	}

	if judul == "" {
		if err := im.db.Offset((page - 1) * perPage).Limit(perPage).Find(&listItem).Error; err != nil {
			logrus.Error("model : error get item", err.Error())
			return nil
		}
	}

	return listItem

}
