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
	FindProductByID(ID int64)(*model.Product,error)
	FindProductByName(name string)(product []model.Product,err error)
	FindProductByArea(area string)(product []model.Product,err error)
	FindProductByCustom(custom int64)(product []model.Product,err error)
	FindProductByStay(is bool)(product []model.Product,err error)
	FindProductByImportant(is bool)(product []model.Product,err error)
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
	product := &model.Product{}
	data.ID = id
	temp := p.mysql.Model(&product).Where("id = ?",id).Updates(&data)
	if data.Is == false{
		temp2 := p.mysql.Model(&product).Where("id = ?",id).Update("is",false)
		return temp2.Error
	}
	if data.Important == false{
		temp3 := p.mysql.Model(&product).Where("id = ?",id).Update("important",false)
		return temp3.Error
	}
	return temp.Error
}
func(p *ProductRepository) FindProductByID(ID int64)(*model.Product,error){
	product := &model.Product{}
	return product,p.mysql.First(product,ID).Error
}
func(p *ProductRepository) FindProductByName(name string)(product []model.Product,err error){
	return product,p.mysql.Where("ProductName = ?",name).Find(product).Error
}
func(p *ProductRepository) FindProductByArea(area string)(product []model.Product,err error){
	return product,p.mysql.Where("BelongArea = ?",area).Find(product).Error
}
func(p *ProductRepository) FindProductByCustom(custom int64)(product []model.Product,err error){
	return product,p.mysql.Where("BelongCustom = ?",custom).Find(product).Error
}
func(p *ProductRepository) FindProductByStay(is bool)(product []model.Product,err error){
	if is{
		return product,p.mysql.Where("Is = ?",true).Find(product).Error
	}
	return product,p.mysql.Where("Is = ?",false).Find(product).Error
}
func(p *ProductRepository) FindProductByImportant(is bool)(product []model.Product,err error){
	if is{
		return product,p.mysql.Where("Important = ?",true).Find(product).Error
	}
	return product,p.mysql.Where("Important = ?",false).Find(product).Error
}