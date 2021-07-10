package model
type Product struct{
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	ProductName string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	Level int64 `json:"level"`
	Category int64 `json:"category"`//指向categoryid
	Important bool `json:"important,bool"`//说明是否重要
	Is bool `json:"is"`//是否在库
	BelongCustom int64 `json:"belong_custom"`//当前所属用户ID
	BelongArea int64 `json:"belong_area"`//所属库房
	Location string `json:"location"`//最新的定位信息
	Rfid string `json:"rfid" gorm:"unique_index;not_null"`//rfid标记
	ImageID int64 `json:"image_id"`//图片地址对应的id（可上传）
	ThreeD string `json:"three_d"`//模型URL地址
}
