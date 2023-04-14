package model

type Contact struct {
	ContactName       string `json:"contactName" validate:"required"`        //联系人姓名
	ContactEmail      string `json:"contactEmail" validate:"required,email"` // 联系人电子邮箱
	ContactPhone      string `json:"contactPhone" validate:"required"`       // 联系人手机号
	ContactQQ         string `json:"contactQQ" validate:"required"`          // 联系人qq
	ContactDepartment string `json:"contactDepartment" validate:"required"`  //部门职务
	ContactWeChat     string `json:"contactWeChat" validate:"required"`      //微信账号
	CustomerId        string `json:"customerId" validate:"required"`         // 关联客户
}
