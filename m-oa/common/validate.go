package common

import "regexp"

func VerifyMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	regular := "/^1[3-9]\\d{9}$/"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}
