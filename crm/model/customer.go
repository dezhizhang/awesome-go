package model

//原型地址https://www.axureshop.com/ys/1967831

type Customer struct {
	CustomerName        string `json:"customerName" binding:"required"`     //客户名称
	CustomerCategory    int    `json:"customerCategory" binding:"required"` //客户类别
	CustomerIndustry    string `json:"customerIndustry"`                    //客户行业
	EnterpriseScale     string `json:"enterpriseScale"`                     //企业规模
	CustomerSource      int    `json:"customerSource"`                      //客户来源
	CustomerAttribution string `json:"customerAttribution"`                 //客户归属
}
