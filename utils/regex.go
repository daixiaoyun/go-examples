package utils

import (
	"regexp"
)

func IsDigit(s string) bool {
	if m, _ := regexp.MatchString("^[0-9]+$", s); m {
		return true
	}

	return false
}


func IsPhone(s string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", s); m {
		return true
	}

	return false
}

func IsEmail(s string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, s); m {
		return true
	}

	return false
}

func IsChinese(s string) bool {
	if m, _ := regexp.MatchString("^\\p{Han}+$", s); m {
		return true
	}
	return false
}


func IsIdCardNumber(s string) bool {
	//验证15位身份证，15位的是全部数字
	if m, _ := regexp.MatchString(`^(\d{15})$`, s); m {
		return true
	}

	//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, s); m {
		return true
	}

	return false
}
