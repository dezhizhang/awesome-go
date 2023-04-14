package model

import "time"

//原型地址https://www.axureshop.com/ys/1967831

type Customer struct {
	CustomerName        string    `json:"customerName" binding:"required"`      //客户名称
	CustomerCategory    int       `json:"customerCategory" binding:"required"`  //客户类别
	CustomerIndustry    string    `json:"customerIndustry"`                     //客户行业
	EnterpriseScale     string    `json:"enterpriseScale"`                      //企业规模
	CustomerSource      int       `json:"customerSource"`                       //客户来源
	CustomerAttribution string    `json:"customerAttribution"`                  //客户归属
	ContactName         string    `json:"contactName" binding:"required"`       //系人姓名
	ContactEmail        string    `json:"contactEmail"`                         //电子邮箱
	ContactNumber       string    `json:"contactNumber" binding:"required"`     //联系号码
	ContactQQ           string    `json:"contactQQ" binding:"required"`         //QQ号码
	ContactDepartment   string    `json:"contactDepartment" binding:"required"` //部门职务
	ContactWeChat       string    `json:"contactWeChat" binding:"required"`     //微信账号
	CreateTime          time.Time `json:"createTime"`                           // 创建时间
	UpdateTime          time.Time `json:"updateTime"`                           //更新时间
}
