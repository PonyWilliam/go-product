package repository

import (
	model "github.com/PonyWilliam/go-product/domain/model"
	"github.com/jinzhu/gorm"
)

type IProduct interface {
	InitTable() error
	AddProduct(product *model.Product) (int64,error)
	DelProductByID(ID int64) error
	UpdateProductByID(id int64,data *model.Product)error
	FindProductByID(ID int64)(model.Product,error)
	FindProductByName(name string)(product []model.Product,err error)
	FindProductByArea(area int64)(product []model.Product,err error)
	FindProductByCustom(custom int64)(product []model.Product,err error)
	FindProductByStay(is bool)(product []model.Product,err error)
	FindProductByImportant(is bool)(product []model.Product,err error)
	FindProductByRFID(rfid string)(product model.Product,err error)
	FindProductAll()([]model.Product,error)
}
func NewProductRepository(db *gorm.DB) IProduct{
	return &ProductRepository{mysql: db}
}
type ProductRepository struct {
	mysql *gorm.DB
}
func(p *ProductRepository) InitTable() error{
	if(p.mysql.HasTable(&model.Product{})){
		return nil
	}
	return p.mysql.CreateTable(&model.Product{}).Error
}
func(p *ProductRepository) AddProduct(product *model.Product) (int64,error){
	return product.ID,p.mysql.Model(product).Create(&product).Error
}
func(p *ProductRepository) DelProductByID(ID int64) error{
	return p.mysql.Where("id = ?",ID).Delete(&model.Product{}).Error
}
func(p *ProductRepository) UpdateProductByID(id int64,data *model.Product) error{

	temp := map[string]interface{}{
		"ProductName":data.ProductName,
		"ProductDescription":data.ProductDescription,
		"Level":data.Level,
		"Category":data.Category,
		"Important":data.Important,
		"Is":data.Is,
		"BelongCustom":data.BelongCustom,
		"BelongArea":data.BelongArea,
		"Location":data.Location,
		"Rfid":data.Rfid,
		"ImageID":data.ImageID,
	}
	return p.mysql.Model(&model.Product{}).Where("id = ?",id).Updates(temp).Error
}
func(p *ProductRepository) FindProductByID(ID int64)(product model.Product,err error){
	return product,p.mysql.Where("id = ?",ID).First(&product).Error
}
func(p *ProductRepository) FindProductByName(name string)(product []model.Product,err error){
	return product,p.mysql.Where("product_name = ?",name).Find(&product).Error
}
func(p *ProductRepository) FindProductByArea(area int64)(product []model.Product,err error){
	return product,p.mysql.Where("belong_area = ?",area).Find(&product).Error
}
func(p *ProductRepository) FindProductByCustom(custom int64)(product []model.Product,err error){
	return product,p.mysql.Where("belong_custom = ?",custom).Find(&product).Error
}
func(p *ProductRepository) FindProductByStay(is bool)(product []model.Product,err error){
	if is{
		return product,p.mysql.Where("Is = ?",true).Find(product).Error
	}
	return product,p.mysql.Where("Is = ?",false).Find(&product).Error
}
func(p *ProductRepository) FindProductByImportant(is bool)(product []model.Product,err error){
	if is{
		return product,p.mysql.Where("Important = ?",true).Find(&product).Error
	}
	return product,p.mysql.Where("Important = ?",false).Find(&product).Error
}
func(p *ProductRepository) FindProductByRFID(rfid string)(product model.Product,err error){
	return product,p.mysql.Where("rfid = ?",rfid).First(&product).Error
}
func(p *ProductRepository) FindProductAll()(products []model.Product,err error){
	return products,p.mysql.Model(&products).Find(&products).Error
}