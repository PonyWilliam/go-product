package service

import (
	"github.com/PonyWilliam/go-product/domain/model"
	"github.com/PonyWilliam/go-product/domain/repository"
)

type IProductServices interface {
	AddProduct(product *model.Product) (int64,error)
	DelProductByID(ID int64) error
	UpdateProductByID(id int64,data *model.Product)error
	FindProductByID(ID int64)(*model.Product,error)
	FindProductByName(name string)(product []model.Product,err error)
	FindProductByArea(area string)(product []model.Product,err error)
	FindProductByCustom(custom int64)(product []model.Product,err error)
	FindProductByStay(is bool)(product []model.Product,err error)
	FindProductByImportant(is bool)(product []model.Product,err error)
	FindProductByRFID(rfid string)(product model.Product,err error)
}
func NewProductServices(repository repository.IProduct)IProductServices{
	return &ProductServices{repository}
}
type ProductServices struct{
	productRepository repository.IProduct
}
func(p *ProductServices) AddProduct(product *model.Product)(int64,error){
	return p.productRepository.AddProduct(product)
}
func(p *ProductServices) DelProductByID(ID int64) error{
	return p.productRepository.DelProductByID(ID)
}
func(p *ProductServices) UpdateProductByID(id int64,data *model.Product)error{
	return p.productRepository.UpdateProductByID(id,data)
}
func(p *ProductServices) FindProductByID(ID int64)(*model.Product,error){
	return p.productRepository.FindProductByID(ID)
}
func(p *ProductServices) FindProductByName(name string)(product []model.Product,err error){
	return p.productRepository.FindProductByName(name)
}
func(p *ProductServices) FindProductByArea(area string)(product []model.Product,err error){
	return p.productRepository.FindProductByArea(area)
}
func(p *ProductServices) FindProductByCustom(custom int64)(product []model.Product,err error){
	return p.productRepository.FindProductByCustom(custom)
}
func(p *ProductServices) FindProductByStay(is bool)(product []model.Product,err error){
	return p.productRepository.FindProductByStay(is)
}
func(p *ProductServices) FindProductByImportant(is bool)(product []model.Product,err error){
	return p.productRepository.FindProductByImportant(is)
}
func(p *ProductServices) FindProductByRFID(rfid string)(product model.Product,err error){
	return p.productRepository.FindProductByRFID(rfid)
}
