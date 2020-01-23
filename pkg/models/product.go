package models

import (
	"simple-REST/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	//Id          string `json:"id"`
	Name       string `gorm:""json:"name"`
	Price      string `json:"price"`
	Category   string `json:"category"`
	Stock      string  `json:"stock"`

}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func (b *Product) CreateProduct() *Product {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func  GetAllProducts() []Product {
	var Products []Product
	db.Find(&Products)
	return Products
}

func GetProductById(Id int64) (*Product , *gorm.DB){
	var getProduct Product
	db:=db.Where("ID = ?", Id).Find(&getProduct)
	return &getProduct, db
}

func DeleteProduct(ID int64) Product {
	var product Product
	db.Where("ID = ?", ID).Delete(product)
	return product
}