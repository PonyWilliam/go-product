package handler

import (
	"context"
	"github.com/PonyWilliam/go-product/domain/model"
	service "github.com/PonyWilliam/go-product/domain/service"
	products "github.com/PonyWilliam/go-product/proto"
	"github.com/micro/go-micro/v2/util/log"
	"strconv"
)
type Product struct{
	ProductServices service.IProductServices
}
func(p *Product)AddProduct(ctx context.Context,info *products.Request_ProductInfo,product *products.Response_Product) error {
	IProduct := &model.Product{
		ID: info.Id,
		ProductName: info.ProductName,
		Level: info.ProductLevel,
		ProductDescription: info.ProductDescription,
		Category: info.ProductBelongCategory,
		Important: info.IsImportant,
		Is:info.ProductIs,
		BelongCustom: info.ProductBelongCustom,
		BelongArea: info.ProductBelongArea,
		Location: info.ProductLocation,
		Rfid: info.ProductRfid,
		ImageID:info.ImageId,
		ThreeD: info.ThreeUrl,
	}
	id,err := p.ProductServices.AddProduct(IProduct)
	if err!=nil{
		product.Message = err.Error()
		return err
	}
	product.Message = strconv.FormatInt(id, 10)
	return nil
}
func(p *Product)DelProduct(ctx context.Context,id *products.Request_ProductID,product *products.Response_DelProduct)error{
	err := p.ProductServices.DelProductByID(id.Id)
	if err!=nil{
		product.Message = err.Error()
		return err
	}
	return nil
}
func(p *Product)ChangeProduct(ctx context.Context,info *products.Request_ProductInfo,product *products.Response_Product)error{
	IProduct := &model.Product{
		ProductName: info.ProductName,
		Level: info.ProductLevel,
		ProductDescription: info.ProductDescription,
		Category: info.ProductBelongCategory,
		Important: info.IsImportant,
		Is:info.ProductIs,
		BelongCustom: info.ProductBelongCustom,
		BelongArea: info.ProductBelongArea,
		Location: info.ProductLocation,
		Rfid: info.ProductRfid,
		ImageID:info.ImageId,
		ThreeD: info.ThreeUrl,
	}
	err := p.ProductServices.UpdateProductByID(info.Id,IProduct)
	if err!=nil{
		product.Message = err.Error()
		log.Error(err)
		return err
	}
	product.Message = "success,id:" + strconv.FormatInt(info.Id, 10)
	return nil
}
func(p *Product)FindProductByID(ctx context.Context,req *products.Request_ProductID,rsp *products.Response_ProductInfo)error{
	productMessage,err := p.ProductServices.FindProductByID(req.Id)
	if err!=nil{
		rsp = nil
		return err
	}
	SwapToProduct(productMessage,rsp)
	return nil
}
func(p *Product)FindProductByRFID(ctx context.Context,req *products.Request_ProductRFID,rsp *products.Response_ProductInfo)error{
	productMessage,err := p.ProductServices.FindProductByRFID(req.Rfid)
	if err!=nil{
		rsp = nil
		return err
	}
	SwapToProduct(productMessage,rsp)
	return nil
}
//rpc AddProduct(Request_ProductInfo)returns(Response_Product);
//rpc DelProduct(Request_ProductID)returns(Response_DelProduct);
//rpc ChangeProduct(Request_ProductInfo)returns(Response_Product);//考虑状态变更
//rpc FindProductByID(Request_ProductID)returns(Response_ProductInfo);
//rpc FindProductByRFID(Request_ProductRFID)returns(Response_ProductInfo);
//rpc FindProductByName(Request_ProductName)returns(Response_ProductInfos);
//rpc FindProductByArea(Request_ProductArea)returns(Response_ProductInfos);
//rpc FindProductByCustom(Request_ProductCustom)returns(Response_ProductInfos);
func (p *Product)FindProductByName(ctx context.Context,req *products.Request_ProductName,rsp *products.Response_ProductInfos)error{
	productMessages,err := p.ProductServices.FindProductByName(req.Name)
	if err!=nil{
		rsp.Infos = nil
		return err
	}
	for _,v := range productMessages{
		product := &products.Response_ProductInfo{}
		SwapToProduct(v,product)
		rsp.Infos = append(rsp.Infos,product)
	}
	return nil
}
func (p *Product)FindProductByArea(ctx context.Context,req *products.Request_ProductArea,rsp *products.Response_ProductInfos)error{
	productMessages,err := p.ProductServices.FindProductByArea(req.Aid)
	if err!=nil{
		rsp.Infos = nil
		return err
	}
	for _,v := range productMessages{
		product := &products.Response_ProductInfo{}
		SwapToProduct(v,product)
		rsp.Infos = append(rsp.Infos,product)
	}
	return nil
}
func (p *Product)FindProductByCustom(ctx context.Context,req *products.Request_ProductCustom,rsp *products.Response_ProductInfos)error{
	productMessages,err := p.ProductServices.FindProductByCustom(req.Wid)
	if err!=nil{
		rsp.Infos = nil
		return err
	}
	for _,v := range productMessages{
		product := &products.Response_ProductInfo{}
		SwapToProduct(v,product)
		rsp.Infos = append(rsp.Infos,product)
	}
	return nil
}
func (p *Product)FindAll(ctx context.Context,req *products.Request_Null,rsp *products.Response_ProductInfos)error{
	productMessages,err := p.ProductServices.FindProductAll()
	if err!=nil{
		rsp.Infos = nil
		return err
	}
	for _,v := range productMessages{
		product := &products.Response_ProductInfo{}
		SwapToProduct(v,product)
		rsp.Infos = append(rsp.Infos,product)
	}
	return nil
}
func (p *Product)FindProductByCategory(ctx context.Context,req *products.Request_ProductCategory,rsp *products.Response_ProductInfos)error{
	productMessages,err := p.ProductServices.FindProductByCategory(req.Cid)
	if err!=nil{
		rsp.Infos = nil
		return err
	}
	for _,v := range productMessages{
		product := &products.Response_ProductInfo{}
		SwapToProduct(v,product)
		rsp.Infos = append(rsp.Infos,product)
	}
	return nil
}
func SwapToProduct(req model.Product,rsp *products.Response_ProductInfo){
	rsp.Id = req.ID
	rsp.ProductLevel = req.Level
	rsp.ProductIs = req.Is
	rsp.IsImportant = req.Important
	rsp.ImageId = req.ImageID
	rsp.ProductBelongCustom = req.BelongCustom
	rsp.ProductBelongArea = req.BelongArea
	rsp.ProductLocation = req.Location
	rsp.ProductBelongCategory = req.Category
	rsp.ProductDescription = req.ProductDescription
	rsp.ProductRfid = req.Rfid
	rsp.ProductName = req.ProductName
	rsp.ThreeUrl = req.ThreeD
}