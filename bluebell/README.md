### 请求参数的转换
```go
import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var Trans ut.Translator

func InitTrans(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()

		uni := ut.New(enT, zhT, enT)

		var ok bool

		Trans, ok = uni.GetTranslator(local)
		if !ok {
			return fmt.Errorf("转换失败%s", local)
		}
		switch local {
		case "en":
			err = enTranslation.RegisterDefaultTranslations(v, Trans)
			break
		case "zh":
			err = zhTranslation.RegisterDefaultTranslations(v, Trans)
			break
		default:
			err = enTranslation.RegisterDefaultTranslations(v, Trans)
			break

		}
		return
	}
	return err
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
```
### 请求参数的校验
```go
func RegisterHandler(c *gin.Context) {

	//var reg model.UserRegister
	user := new(model.UserRegister)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "请求成功"})
}
```