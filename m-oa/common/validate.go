package common

import "regexp"

func VerifyMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	regular := "/^1(3[0-9]|4[01456879]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\\d{8}$/"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}
