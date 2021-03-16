package handler

import (
	"context"
	"fmt"
	"github.com/PonyWilliam/go-product/domain/model"
	service "github.com/PonyWilliam/go-product/domain/service"
	products "github.com/PonyWilliam/go-product/proto"
	"github.com/micro/go-micro/v2/util/log"
	"strconv"
)
/*
package model
type Product struct{
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	ProductName string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	Level int64 `json:"level"`
	Category int64 `json:"category"`//指向categoryid
	Important bool `json:"important"`//说明是否重要
	Is bool `json:"is"`//是否在库
	BelongCustom int64 `json:"belong_custom"`//当前所属用户ID
	BelongArea int64 `json:"belong_area"`//所属库房
	Location string `json:"location"`//最新的定位信息
	Rfid string `json:"rfid"`//rfid标记
	ImageID int64 `json:"image_id"`//图片地址对应的id（可上传）
}

*/

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
	fmt.Println(info.ProductIs)
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
	}
	fmt.Println(IProduct)
	fmt.Println(IProduct.Is)
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
		log.Error(err)
		rsp.Info = nil
		return err
	}
	/*
	int64 id = 1;//唯一ID
		string product_name = 2;//名称
		string product_rfid = 3;//rfid特征码
		int64 product_level = 4;//物品等级
		string product_description = 5;//物品描述信息
		bool product_is = 6;//是否还在仓库
		int64 product_belong_category = 7;//物品所属分类id
		string product_location = 8;//物品最新位置,如果在库则不做记录
		int64 product_belong_area = 9;//物品所在库房id
		int64 product_belong_custom = 10;//借走的用户ID,状态变更记录在变更表。通过id快速追踪。
		int64 image_id = 11;
		bool is_important = 12;
	*/
	temp := &products.Request_ProductInfo{
			Id: productMessage.ID,
			ProductName: productMessage.ProductName,
			ProductLevel: productMessage.Level,
			ProductDescription: productMessage.ProductDescription,
			ProductIs:productMessage.Is,
			ProductBelongCategory: productMessage.Category,
			ProductLocation: productMessage.Location,
			ProductBelongArea: productMessage.BelongArea,
			ProductBelongCustom: productMessage.BelongCustom,
			ProductRfid: productMessage.Rfid,
			ImageId: productMessage.ImageID,
			IsImportant: productMessage.Important,
		}
	rsp.Info = temp
	return nil
}
// Call is a single request handler called via client.Call or the generated client code
